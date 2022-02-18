package libpay

type Card struct {
	Number           string
	Month            int
	Year             int
	VerificationCode string
}

func (c Card) Validate() error {
	return nil
}

type Gateway interface {
	Authorize(amount int64, card Card, opts ...interface{}) (*Authorization, error)
}

type Authorization struct{}
