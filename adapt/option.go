package adapt

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

// Options is a map of options.
type Options map[string]any

// NewOptions creates options with defaults.
func NewOptions(opts ...Option) Options {
	options := make(Options)
	for _, opt := range opts {
		opt(options)
	}

	return options
}

// Option is for setting options.
type Option func(Options)

// With sets the key-value pair.
func With(k string, v any) Option {
	return func(o Options) {
		o[k] = v
	}
}

// WithSecurity sets the security code.
func WithSecurity(code string) Option {
	return With("security", NewSecurity(code))
}

// WithSecurities sets the security list.
func WithSecurities(codes []string) Option {
	return With("securityList", NewSecurities(codes))
}

// ToProto converts options to proto message.
func (o Options) ToProto(msg proto.Message) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, msg)
}
