package libpay

type Amount struct {
	Currency string
	Value    int64
}

type Gateway interface {
	Authorize(card Card, amount Amount, opts ...interface{}) (*Authorization, error)
	Capture(auth Authorization, amount Amount, opts ...interface{}) (*Capture, error)
}
