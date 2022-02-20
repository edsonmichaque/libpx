package libpx

import (
	"gitlab.com/edsonmichaque/libpx/currency"
)

type Provider interface {
	SupportedSources() []string
	Schema() map[string]Schema
	Currencies() []currency.Currency
	Configure(args ...ProviderConfigOption) error
	Authorize(Card, Amount, ...Option) (*Authorization, error)
	Capture(Authorization, Amount, ...Option) (*Capture, error)
	Purchase(Authorization, Amount, ...Option) (*Purchase, error)
	Refund(Authorization, Amount) (*Refund, error)
	Void(Authorization) (*Void, error)
	Verify(Card, ...Option) (*Verification, error)
}

type Amount struct {
	Currency currency.Currency
	Value    int64
}

type Purchase struct{}

type Verification struct{}

type Refund struct{}

type Capture struct{}

type Authorization struct{}

type Address struct {
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	Country  string
	Zip      string
	Phone    string
}

type Void struct{}
