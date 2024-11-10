package adapt

import (
	"encoding/json"

	"github.com/hyperjiang/futu/pb/qotstockfilter"
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

// WithBaseFilters sets the base filter list.
func WithBaseFilters(filters ...*qotstockfilter.BaseFilter) Option {
	return func(o Options) {
		o["baseFilterList"] = filters
	}
}

// WithAccumulateFilters sets the accumulate filter list.
func WithAccumulateFilters(filters ...*qotstockfilter.AccumulateFilter) Option {
	return func(o Options) {
		o["accumulateFilterList"] = filters
	}
}

// WithFinancialFilters sets the financial filter list.
func WithFinancialFilters(filters ...*qotstockfilter.FinancialFilter) Option {
	return func(o Options) {
		o["financialFilterList"] = filters
	}
}

// WithPatternFilters sets the pattern filter list.
func WithPatternFilters(filters ...*qotstockfilter.PatternFilter) Option {
	return func(o Options) {
		o["patternFilterList"] = filters
	}
}

// WithCustomIndicatorFilters sets the custom indicator filter list.
func WithCustomIndicatorFilters(filters ...*qotstockfilter.CustomIndicatorFilter) Option {
	return func(o Options) {
		o["customIndicatorFilterList"] = filters
	}
}

// ToProto converts options to proto message.
func (o Options) ToProto(msg proto.Message) error {
	b, err := json.Marshal(o)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, msg)
}
