package libpx

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

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
