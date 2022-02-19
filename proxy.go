package libpx

import "github.com/edsonmichaque/libpx/card"

func Create(p Provider) Proxy {
	return Proxy{
		provider: p,
	}
}

type Proxy struct {
	provider Provider
}

func (p Proxy) Configure(args ...ConfigOption) error {
	return p.provider.Configure(args...)
}

func (p Proxy) Authorize(c card.Card, amount Amount, opts ...Option) (*Authorization, error) {
	return p.provider.Authorize(c, amount, opts...)
}
