package libpx

type Error struct{}

func (e Error) Error() string {
	return ""
}

type InvalidCreditCardError struct{}

func (e InvalidCreditCardError) Error() string {
	return ""
}
