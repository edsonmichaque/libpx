package libpay

type Gateway interface {
	Authorize(amount int64, card Card, opts ...interface{}) (*Authorization, error)
	Capture(amount int64, auth Authorization, opts ...interface{}) (*Capture, error)
}
