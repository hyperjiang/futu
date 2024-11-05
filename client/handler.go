package client

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// Handler is the definition of a handler function.
type Handler func(s2c proto.Message) error

func defaultHandler(s2c proto.Message) error {
	log.Info().Interface("s2c", s2c).Msg("notification")
	return nil
}
