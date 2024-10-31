package client

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hyperjiang/futu/infra"
	"github.com/hyperjiang/futu/pb/common"
	"github.com/hyperjiang/futu/pb/initconnect"
	"github.com/hyperjiang/futu/pb/keepalive"
	"github.com/hyperjiang/futu/pb/notify"
	"github.com/hyperjiang/futu/pb/qotupdatebasicqot"
	"github.com/hyperjiang/futu/pb/qotupdatebroker"
	"github.com/hyperjiang/futu/pb/qotupdatekl"
	"github.com/hyperjiang/futu/pb/qotupdateorderbook"
	"github.com/hyperjiang/futu/pb/qotupdatepricereminder"
	"github.com/hyperjiang/futu/pb/qotupdatert"
	"github.com/hyperjiang/futu/pb/qotupdateticker"
	"github.com/hyperjiang/futu/pb/trdupdateorder"
	"github.com/hyperjiang/futu/pb/trdupdateorderfill"
	"github.com/hyperjiang/futu/protoid"
	"github.com/hyperjiang/rsa"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// Client is the client to connect to Futu OpenD.
type Client struct {
	Options

	conn     net.Conn
	sn       atomic.Uint32 // serial number
	resChan  chan response // response channel
	closed   chan struct{} // indicate the client is closed
	hub      *infra.DispatcherHub
	connID   uint64
	userID   uint64
	crypto   *infra.Crypto
	handlers sync.Map // push notification handlers
}

// New creates a new client.
func New(opts ...Option) (*Client, error) {
	client := &Client{
		Options:  NewOptions(opts...),
		closed:   make(chan struct{}),
		hub:      infra.NewDispatcherHub(),
		handlers: sync.Map{},
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

	if client.PrivateKey != nil || client.PublicKey != nil {
		client.crypto, err = infra.NewCrypto([]byte(s2c.GetConnAESKey()), []byte(s2c.GetAesCBCiv()))
		if err != nil {
			client.Close()
			return nil, err
		}
	}

	if interval := s2c.GetKeepAliveInterval(); interval > 0 {
		go client.heartbeat(time.Second * time.Duration(interval))
	}

	go client.watchNotification()

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

// Request sends a request to the server.
func (client *Client) Request(protoID uint32, req proto.Message, resCh *infra.ProtobufChan) error {
	var buf bytes.Buffer

	b, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	sha1Value := sha1.Sum(b)

	if client.PublicKey != nil {
		if protoID == protoid.InitConnect {
			b, err = rsa.EncryptByPublicKey(b, client.PublicKey)
			if err != nil {
				return err
			}
		} else {
			b = client.crypto.Encrypt(b)
		}
	}

	sn := client.nextSN()

	h := futuHeader{
		HeaderFlag:   [2]byte{'F', 'T'},
		ProtoID:      protoID,
		ProtoFmtType: 0,
		ProtoVer:     0,
		SerialNo:     sn,
		BodyLen:      uint32(len(b)),
		BodySHA1:     sha1Value,
	}

	client.registerDispatcher(protoID, sn, resCh)

	if err := binary.Write(&buf, binary.LittleEndian, &h); err != nil {
		return err
	}

	if _, err := buf.Write(b); err != nil {
		return err
	}

	_, err = buf.WriteTo(client.conn)

	return err
}

// RegisterHandler registers a handler for push notifications of the specified protoID.
func (client *Client) RegisterHandler(protoID uint32, h Handler) *Client {
	client.handlers.Store(protoID, h)
	return client
}

func (client *Client) getHandler(protoID uint32) Handler {
	if h, ok := client.handlers.Load(protoID); ok {
		return h.(Handler)
	}

	return defaultHandler
}

// watchNotification watches the push notification.
// no need to close the channels in this function,
// because they will be closed by the dispatcher hub when the client is closed.
func (client *Client) watchNotification() {
	notifyCh := make(chan *notify.Response, 1)
	client.registerDispatcher(protoid.Notify, 0, infra.NewProtobufChan(notifyCh))

	updateOrderCh := make(chan *trdupdateorder.Response, 1)
	client.registerDispatcher(protoid.TrdUpdateOrder, 0, infra.NewProtobufChan(updateOrderCh))

	updateOrderFillCh := make(chan *trdupdateorderfill.Response, 1)
	client.registerDispatcher(protoid.TrdUpdateOrderFill, 0, infra.NewProtobufChan(updateOrderFillCh))

	updateBasicQotCh := make(chan *qotupdatebasicqot.Response, 1)
	client.registerDispatcher(protoid.QotUpdateBasicQot, 0, infra.NewProtobufChan(updateBasicQotCh))

	updateKLCh := make(chan *qotupdatekl.Response, 1)
	client.registerDispatcher(protoid.QotUpdateKL, 0, infra.NewProtobufChan(updateKLCh))

	updateRT := make(chan *qotupdatert.Response, 1)
	client.registerDispatcher(protoid.QotUpdateRT, 0, infra.NewProtobufChan(updateRT))

	updateTicker := make(chan *qotupdateticker.Response, 1)
	client.registerDispatcher(protoid.QotUpdateTicker, 0, infra.NewProtobufChan(updateTicker))

	updateOrderBook := make(chan *qotupdateorderbook.Response, 1)
	client.registerDispatcher(protoid.QotUpdateOrderBook, 0, infra.NewProtobufChan(updateOrderBook))

	updateBroker := make(chan *qotupdatebroker.Response, 1)
	client.registerDispatcher(protoid.QotUpdateBroker, 0, infra.NewProtobufChan(updateBroker))

	updatePriceReminder := make(chan *qotupdatepricereminder.Response, 1)
	client.registerDispatcher(protoid.QotUpdatePriceReminder, 0, infra.NewProtobufChan(updatePriceReminder))

	for {
		select {
		case <-client.closed:
			log.Info().Msg("stop watching notification")
			return
		case resp, ok := <-notifyCh:
			if !ok {
				log.Info().Msg("notification channel closed")
				break
			}
			if err := client.getHandler(protoid.Notify)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("notification handle error")
			}

		case resp, ok := <-updateOrderCh:
			if !ok {
				log.Info().Msg("update order channel closed")
				break
			}
			if err := client.getHandler(protoid.TrdUpdateOrder)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update order handle error")
			}
		case resp, ok := <-updateOrderFillCh:
			if !ok {
				log.Info().Msg("update order fill channel closed")
				break
			}
			if err := client.getHandler(protoid.TrdUpdateOrderFill)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update order fill handle error")
			}
		case resp, ok := <-updateBasicQotCh:
			if !ok {
				log.Info().Msg("update basic qot channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateBasicQot)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update basic quote handle error")
			}
		case resp, ok := <-updateKLCh:
			if !ok {
				log.Info().Msg("update KL channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateKL)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update KL handle error")
			}
		case resp, ok := <-updateRT:
			if !ok {
				log.Info().Msg("update RT channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateRT)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update RT handle error")
			}
		case resp, ok := <-updateTicker:
			if !ok {
				log.Info().Msg("update ticker channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateTicker)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update ticker handle error")
			}
		case resp, ok := <-updateOrderBook:
			if !ok {
				log.Info().Msg("update order book channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateOrderBook)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update order book handle error")
			}
		case resp, ok := <-updateBroker:
			if !ok {
				log.Info().Msg("update broker channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdateBroker)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update broker handle error")
			}
		case resp, ok := <-updatePriceReminder:
			if !ok {
				log.Info().Msg("update price reminder channel closed")
				break
			}
			if err := client.getHandler(protoid.QotUpdatePriceReminder)(resp.GetS2C()); err != nil {
				log.Error().Err(err).Msg("update price reminder handle error")
			}
		}
	}
}

func (client *Client) registerDispatcher(protoID uint32, sn uint32, ch *infra.ProtobufChan) {
	client.hub.Register(protoID, sn, ch)
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
			log.Info().Uint32("proto_id", res.ProtoID).Uint32("sn", res.SerialNo).Msg("listen")
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

	// read header, it will block until the header is read
	var h futuHeader
	if err := binary.Read(client.conn, binary.LittleEndian, &h); err != nil {
		return err
	}
	if h.HeaderFlag != [2]byte{'F', 'T'} {
		return errors.New("header flag error")
	}
	// read body, it will block until the body is read
	b := make([]byte, h.BodyLen)
	if _, err := io.ReadFull(client.conn, b); err != nil {
		return err
	}

	if client.PrivateKey != nil {
		if h.ProtoID == protoid.InitConnect {
			var err error
			b, err = rsa.DecryptByPrivateKey(b, client.PrivateKey)
			if err != nil {
				return err
			}
		} else {
			b = client.crypto.Decrypt(b)
		}
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
				log.Error().Err(err).Msg("connection closed")
				return
			} else {
				log.Error().Err(err).Msg("decode error")
				return
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
			PacketEncAlgo:       proto.Int32(int32(*common.PacketEncAlgo_PacketEncAlgo_AES_CBC.Enum())),
			ProgrammingLanguage: proto.String("Go"),
		},
	}

	ch := make(chan *initconnect.Response, 1)
	defer close(ch)

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
		return resp.GetS2C(), infra.Error(resp)
	}
}

// keepAlive sends a keep alive request.
// The server will return the server timestamp in seconds.
func (client *Client) keepAlive(ch chan *keepalive.Response, ticker *time.Ticker) (int64, error) {
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
	case <-ticker.C:
		return 0, ErrTimeout
	case resp, ok := <-ch:
		if !ok {
			return 0, ErrChannelClosed
		}
		return resp.GetS2C().GetTime(), infra.Error(resp)
	}
}

func (client *Client) heartbeat(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	ch := make(chan *keepalive.Response, 1)
	defer close(ch)

	for {
		select {
		case <-client.closed:
			log.Info().Msg("heartbeat stopped")
			return
		case <-ticker.C:
			if _, err := client.keepAlive(ch, ticker); err != nil {
				return
			}
		}
	}
}
