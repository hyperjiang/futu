package adapt

import (
	"encoding/json"
	"strings"

	"github.com/hyperjiang/futu/pb/qotstockfilter"
	"github.com/hyperjiang/futu/pb/trdcommon"
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

// WithFilterConditions sets the filter conditions for trade.
func WithFilterConditions(conditions *trdcommon.TrdFilterConditions) Option {
	return func(o Options) {
		o["filterConditions"] = conditions
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

// SetCodeForTrade sets the code for trade.
func (o Options) SetCodeForTrade(code string) Options {
	arr := strings.Split(code, ".")
	if len(arr) == 2 {
		if id := GetTrdMarketID(arr[0]); id > 0 {
			o["secMarket"] = id
			o["code"] = arr[1]
		}
	}

	if _, ok := o["code"]; !ok {
		o["code"] = code
	}

	return o
}
