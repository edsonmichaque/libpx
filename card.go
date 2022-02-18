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
