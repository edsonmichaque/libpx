package libpay

type Authorization struct {
	success bool
}

func (a Authorization) Success() bool {
	return a.success
}

func (a Authorization) Failure() bool {
	return a.Success()
}
