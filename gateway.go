package libpx

import (
	"fmt"
	"strconv"

	"github.com/edsonmichaque/libpx/currency"
)

type Gateway interface {
	Authorize(CreditCard, Amount, ...Option) (*Authorization, error)
	Capture(Authorization, Amount, ...Option) (*Capture, error)
	Purchase(Authorization, Amount, ...Option) (*Purchase, error)
	Refund(Authorization, Amount) (*Refund, error)
	Void(Authorization) (*Void, error)
	Verify(CreditCard, ...Option) (*Verification, error)
}

type ImprovedGateway interface {
	Gateway
	Credit(CreditCard, Amount, ...Option)
	Recurring(CreditCard, Amount, ...Option)
	Store(CreditCard, ...Option)
	Unstore(CreditCard, ...Option)
}

type Amount struct {
	Currency currency.Currency
	Value    int64
}

func (a Amount) Format() string {
	amount := float64(a.Value) / 100

	tpl := a.Currency.Code + " %." + strconv.Itoa(a.Currency.Precision) + "f"

	return fmt.Sprintf(tpl, amount)
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
