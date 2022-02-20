package libpx

import (
	"gitlab.com/edsonmichaque/libpx/card"
	"gitlab.com/edsonmichaque/libpx/currency"
	"gitlab.com/edsonmichaque/libpx/schema"
)

type Config struct {
	Name  string
	Value interface{}
}

type Configs []Config

type ConfigOption func(*Configs)

func WithConfig(name string, value interface{}) ConfigOption {
	return func(c *Configs) {

		var cs []Config
		if c != nil {
			cs = []Config(*c)
		} else {
			cs = make([]Config, 0)
		}

		cs = append(cs, Config{Name: name, Value: value})
		aux := Configs(cs)

		c = &aux
	}
}

type Provider interface {
	SupportedSources() []string
	Schema() map[string]schema.Schema
	Currencies() []currency.Currency
	Configure(args ...ConfigOption) error
	Authorize(card.Card, Amount, ...Option) (*Authorization, error)
	Capture(Authorization, Amount, ...Option) (*Capture, error)
	Purchase(Authorization, Amount, ...Option) (*Purchase, error)
	Refund(Authorization, Amount) (*Refund, error)
	Void(Authorization) (*Void, error)
	Verify(card.Card, ...Option) (*Verification, error)
}

type Amount struct {
	Currency currency.Currency
	Value    int64
}

type Option func(*Options)

type Options struct {
	ShippingAddress   Address
	BillingAddress    Address
	IP                string
	Email             string
	AdditionalOptions []AdditionalOption
}

func WithShippingAddress(a Address) Option {
	return func(o *Options) {
		o.ShippingAddress = a
	}
}

func WithBillingAddress(a Address) Option {
	return func(o *Options) {
		o.BillingAddress = a
	}
}

func WithAdditionalOption(name, value string) Option {
	return func(o *Options) {
		if o.AdditionalOptions == nil {
			o.AdditionalOptions = make([]AdditionalOption, 0)
		}

		o.AdditionalOptions = append(o.AdditionalOptions, AdditionalOption{Name: name, Value: value})
	}
}

type AdditionalOption struct {
	Name  string
	Value interface{}
}
