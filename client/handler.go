package client

import (
	"github.com/rs/zerolog/log"
)

// Handler is the definition of a handler function.
type Handler func(s2c interface{}) error

func defaultNotificationHandler(s2c interface{}) error {
	log.Info().Interface("s2c", s2c).Msg("notification")
	return nil
}
