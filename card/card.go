package card

import (
	"errors"
	"strings"
)

type Type uint

const (
	Credit Type = iota
	Debit
	PrePaid
)

type Network uint

const (
	Visa Network = iota
	MasterCard
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
	return c.validate()
}

func (c Card) validate() error {
	if strings.HasPrefix(c.Number, "4") {
		return nil
	}

	return errors.New("")
}
