package libpx

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

type Void struct{}

type Capabilities struct {
	Currencies []string
	Countries  []string
	Features   []string
}

type Currency struct {
	Code      string
	Name      string
	Number    string
	Precision int
}

var (
	MZN = Currency{Code: "MZN", Precision: 2}
)
