package libpay

type Purchase struct{}

type Verification struct{}

type Refund struct{}

type Capture struct {
	Redirect             string
	TransactionReference string
}

type Authorization struct {
	Redirect             string
	TransactionReference string
	success              bool
}

func (a Authorization) Success() bool {
	return a.success
}

func (a Authorization) Failure() bool {
	return a.Success()
}

type Address struct {
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	Country  string
	Zip      string
	Phone    string
}

type Network int

const (
	NetworkVisa = iota
	NetworkMastercard
)

type CreditCard struct {
	Number          string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
	FirstName       string
	LastName        string
	Type            Network
}

func (c CreditCard) Validate() error {
	return nil
}
