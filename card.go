package libpx

import (
	"errors"
	"strings"
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
