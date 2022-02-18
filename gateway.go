package libpay

type Amount struct {
	Currency string
	Value    int64
}

type Option func(*Options)

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

func WithExtra(key, value string) Option {
	return func(o *Options) {
		if o.Extra == nil {
			o.Extra = make([]Extra, 0)
		}

		o.Extra = append(o.Extra, Extra{Key: key, Value: value})
	}
}

type Options struct {
	ShippingAddress string
	BillingAddress  string
	IP              string
	Extra           []Extra
}

type Extra struct {
	Key   string
	Value string
}

type Gateway interface {
	Authorize(card Card, amount Amount, opts ...Option) (*Authorization, error)
	Capture(auth Authorization, amount Amount, opts ...Option) (*Capture, error)
	Purchase(auth Authorization, amount Amount, opts ...Option) (*Capture, error)
	Refund(auth Authorization, amount Amount) (*Capture, error)
	Void(auth Authorization) (*Capture, error)
	Verify(card Card, opts ...Option) (*Capture, error)
}
