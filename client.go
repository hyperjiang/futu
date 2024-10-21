package futu

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/initconnect"
	"github.com/hyperjiang/futu/pb/keepalive"
	"github.com/hyperjiang/futu/pb/notify"
	"github.com/hyperjiang/futu/protoid"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Options

	conn     net.Conn
	sn       atomic.Uint32         // serial number
	resChan  chan response         // response channel
	closed   chan struct{}         // indicate the client is closed
	notiChan chan *notify.Response // notification channel (1003)
	hub      *infra.DispatcherHub
	ticker   *time.Ticker
	connID   uint64
	userID   uint64
}

// NewClient creates a new client.
func NewClient(opts ...Option) (*Client, error) {
	client := &Client{
		Options:  NewOptions(opts...),
		closed:   make(chan struct{}),
		notiChan: make(chan *notify.Response),
		hub:      infra.NewDispatcherHub(),
	}
	client.resChan = make(chan response, client.ResChanSize)

	if err := client.dial(); err != nil {
		return nil, err
	}

	go client.listen()
	go client.infiniteRead()

	s2c, err := client.initConnect()
	if err != nil {
		client.Close()
		return nil, err
	}

	log.Info().
		Int32("server_ver", s2c.GetServerVer()).
		Uint64("conn_id", s2c.GetConnID()).
		Uint64("user_id", s2c.GetLoginUserID()).
		Int32("keep_alive_interval", s2c.GetKeepAliveInterval()).
		Int32("user_attr", s2c.GetUserAttribution()).
		Str("conn_aes_key", s2c.GetConnAESKey()).
		Str("aes_cbc_iv", s2c.GetAesCBCiv()).
		Msg("init connect success")

	client.connID = s2c.GetConnID()
	client.userID = s2c.GetLoginUserID()
	if d := s2c.GetKeepAliveInterval(); d > 0 {
		client.ticker = time.NewTicker(time.Second * time.Duration(d))
		go client.heartbeat()
	}

	// client.RegisterDispatcher(protoid.Notify, 0, infra.NewProtobufChan(client.notiChan))

	return client, nil
}

// GetConnID returns the connection ID.
func (client *Client) GetConnID() uint64 {
	return client.connID
}

// GetUserID returns the user ID.
func (client *Client) GetUserID() uint64 {
	return client.userID
}

// Close closes the client.
func (client *Client) Close() error {
	close(client.closed)

	client.hub.Close()

	if client.conn == nil {
		return nil
	}

	return client.conn.Close()
}

// RegisterDispatcher registers a dispatcher.
func (client *Client) RegisterDispatcher(protoID uint32, sn uint32, ch *infra.ProtobufChan) {
	client.hub.Register(protoID, sn, ch)
}

// Request sends a request to the server.
func (client *Client) Request(protoID uint32, req proto.Message, resCh *infra.ProtobufChan) error {
	var buf bytes.Buffer

	b, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	sn := client.nextSN()

	h := futuHeader{
		HeaderFlag:   [2]byte{'F', 'T'},
		ProtoID:      protoID,
		ProtoFmtType: 0,
		ProtoVer:     0,
		SerialNo:     sn,
		BodyLen:      uint32(len(b)),
		BodySHA1:     sha1.Sum(b),
	}

	client.RegisterDispatcher(protoID, sn, resCh)

	if err := binary.Write(&buf, binary.LittleEndian, &h); err != nil {
		return err
	}

	if _, err := buf.Write(b); err != nil {
		return err
	}

	_, err = buf.WriteTo(client.conn)

	return err
}

// GetNotiChan returns the notification channel.
func (client *Client) GetNotiChan() <-chan *notify.Response {
	return client.notiChan
}

// nextSN returns the next serial number.
func (client *Client) nextSN() uint32 {
	return client.sn.Add(1)
}

func (client *Client) dial() error {
	conn, err := net.Dial("tcp", client.Addr)
	if err != nil {
		log.Error().Err(err).Msg("dial failed")
		return err
	}

	client.conn = conn

	return nil
}

func (client *Client) listen() {
	for {
		select {
		case <-client.closed:
			return
		case res := <-client.resChan:
			log.Info().Uint32("proto_id", res.ProtoID).Uint32("sn", res.SerialNo).Msg("")
			if err := client.hub.Dispatch(res.ProtoID, res.SerialNo, res.Body); err != nil {
				log.Error().Err(err).Msg("dispatch error")
			}
		}
	}
}

func (client *Client) read() error {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("recover", r).Msg("")
		}
	}()

	// read header
	var h futuHeader
	if err := binary.Read(client.conn, binary.LittleEndian, &h); err != nil {
		return err
	}
	if h.HeaderFlag != [2]byte{'F', 'T'} {
		return errors.New("header flag error")
	}
	// read body
	b := make([]byte, h.BodyLen)
	if _, err := io.ReadFull(client.conn, b); err != nil {
		return err
	}
	// verify body
	if h.BodySHA1 != sha1.Sum(b) {
		return errors.New("sha1 sum error")
	}

	res := response{
		ProtoID:  h.ProtoID,
		SerialNo: h.SerialNo,
		Body:     b,
	}

	client.resChan <- res

	return nil
}

func (client *Client) infiniteRead() {
	for {
		if err := client.read(); err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				// If the connection is closed, stop receiving data.
				// io.EOF: The connection is closed by the remote end.
				// net.ErrClosed: The connection is closed by the local end.
				return
			} else {
				log.Error().Err(err).Msg("decode error")
			}
		}
	}
}

func (client *Client) initConnect() (*initconnect.S2C, error) {
	req := &initconnect.Request{
		C2S: &initconnect.C2S{
			ClientVer:           proto.Int32(ClientVersion),
			ClientID:            proto.String(client.ID),
			RecvNotify:          proto.Bool(client.RecvNotify),
			ProgrammingLanguage: proto.String("Go"),
		},
	}

	ch := make(chan *initconnect.Response)

	if err := client.Request(protoid.InitConnect, req, infra.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-client.closed:
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		close(ch)
		return resp.GetS2C(), infra.Error(resp)
	}
}

// keepAlive sends a keep alive request.
// The server will return the server timestamp in seconds.
func (client *Client) keepAlive(ch chan *keepalive.Response) (int64, error) {
	req := &keepalive.Request{
		C2S: &keepalive.C2S{
			Time: proto.Int64(time.Now().Unix()),
		},
	}

	if err := client.Request(protoid.KeepAlive, req, infra.NewProtobufChan(ch)); err != nil {
		return 0, err
	}

	select {
	case <-client.closed:
		return 0, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return 0, ErrChannelClosed
		}
		return resp.GetS2C().GetTime(), infra.Error(resp)
	}
}

func (client *Client) heartbeat() {
	ch := make(chan *keepalive.Response)

	for {
		select {
		case <-client.closed:
			client.ticker.Stop()
			close(ch)
			log.Info().Msg("heartbeat stopped")
			return
		case <-client.ticker.C:
			if _, err := client.keepAlive(ch); err != nil {
				return
			}
		}
	}
}
