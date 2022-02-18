package libpay

type Gateway interface {
	Authorize(card Card, amount Amount, opts ...Option) (*Authorization, error)
	Capture(auth Authorization, amount Amount, opts ...Option) (*Capture, error)
	Purchase(auth Authorization, amount Amount, opts ...Option) (*Purchase, error)
	Refund(auth Authorization, amount Amount) (*Refund, error)
	Void(auth Authorization) (*Capture, error)
	Verify(card Card, opts ...Option) (*Verification, error)
}

type Amount struct {
	Currency string
	Value    int64
}

type Option func(*Options)

type Options struct {
	ShippingAddress   string
	BillingAddress    string
	IP                string
	Email             string
	AdditionalOptions []AdditionalOption
}

func WithShippingAddress(a string) Option {
	return func(o *Options) {
		o.ShippingAddress = a
	}
}

func WithBillingAddress(a string) Option {
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
