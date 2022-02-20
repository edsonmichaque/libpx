package libpx

import (
	"gitlab.com/edsonmichaque/libpx/schema"
)

func Build(p Provider) Wrapper {
	return Wrapper{
		provider: p,
	}
}

type Wrapper struct {
	provider Provider
}

func (g Wrapper) Configure(args ...ConfigOption) error {
	if err := g.enforceSchema(g.provider.Schema(), args...); err != nil {
		return err
	}

	return g.provider.Configure(args...)
}

func (g Wrapper) Authorize(card Card, amount Amount, opts ...Option) (*Authorization, error) {
	return g.provider.Authorize(card, amount, opts...)
}

func (g Wrapper) Capture(auth Authorization, amount Amount, opts ...Option) (*Capture, error) {
	return g.provider.Capture(auth, amount, opts...)
}

func (g Wrapper) Purchase(auth Authorization, amount Amount, opts ...Option) (*Purchase, error) {
	return g.provider.Purchase(auth, amount, opts...)
}

func (g Wrapper) Refund(auth Authorization, amount Amount) (*Refund, error) {
	return g.provider.Refund(auth, amount)
}

func (g Wrapper) Void(auth Authorization) (*Void, error) {
	return g.provider.Void(auth)
}

func (g Wrapper) Verify(card Card, opts ...Option) (*Verification, error) {
	return g.provider.Verify(card, opts...)
}

func (g Wrapper) enforceSchema(schema map[string]schema.Schema, args ...ConfigOption) error {
	return nil
}
