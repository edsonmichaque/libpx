package card

import "errors"

type Type uint

const (
	Credit Type = iota
	Debit
	PrePaif
)

type Network uint

const (
	Visa Network = iota
	MasterCard
	Amex
)

type Card struct {
	Type            Type
	Network         Network
	Number          string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
	FirstName       string
	LastName        string
}

func (c Card) Validate() error {
	return errors.New("")
}
