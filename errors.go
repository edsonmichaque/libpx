package libpay

type Error struct{}

func (e Error) Error() string {
	return ""
}
