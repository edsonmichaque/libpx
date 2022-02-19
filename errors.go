package libpx

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return ""
}

type InvalidCreditCardError struct{}

func (e InvalidCreditCardError) Error() string {
	return ""
}

type NotSupportedError struct{}

func (n NotSupportedError) Error() string {
	return ""
}
