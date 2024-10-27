package infra

import (
	"reflect"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// ProtobufChan is a wrapper for chan *T, T is a struct and *T implements proto.Message.
type ProtobufChan struct {
	v reflect.Value
	t reflect.Type
}

// NewProtobufChan creates a new ProtobufChan.
func NewProtobufChan(i any) *ProtobufChan {
	ct := reflect.TypeOf(i)
	if ct == nil {
		return nil
	}

	// must be a channel type
	if ct.Kind() != reflect.Chan {
		return nil
	}
	// it must be a channel of pointer to the response type which implements proto.Message interface
	pt := ct.Elem()
	if pt.Kind() != reflect.Ptr || !pt.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return nil
	}
	st := pt.Elem()
	if st.Kind() != reflect.Struct {
		return nil
	}

	return &ProtobufChan{v: reflect.ValueOf(i), t: st}
}

// Send unmarshals b into a proto message and send it to the channel.
func (ch *ProtobufChan) Send(b []byte) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("recover", r).Msg("protobuf chan send panic")
		}
	}()

	buf := reflect.New(ch.t)
	if err := proto.Unmarshal(b, buf.Interface().(proto.Message)); err != nil {
		return err
	}

	ch.v.Send(buf)

	return nil
}

// Close closes the channel.
func (ch *ProtobufChan) Close() {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Interface("recover", r).Msg("protobuf chan close panic")
		}
	}()

	if ch != nil {
		ch.v.Close()
	}
}
