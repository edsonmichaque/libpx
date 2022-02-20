package libpx

import (
	"errors"
	"strings"
)

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

type CardType uint

const (
	CardCredit CardType = iota
	CardDebit
	CardPrePaid
)

type CardNetwork uint

const (
	Visa CardNetwork = iota
	MasterCard
)

type Card struct {
	Type            CardType
	Network         CardNetwork
	Number          string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
	FirstName       string
	LastName        string
}

func (c Card) Validate() error {
	return c.validate()
}

func (c Card) validate() error {
	if strings.HasPrefix(c.Number, "4") {
		return nil
	}

	return errors.New("")
}
