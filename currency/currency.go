package currency

import (
	"errors"
	"strings"
)

type Currency struct {
	Code      string
	Symbol    string
	Precision int
}

var (
	USD = Currency{Code: "USD", Precision: 2}
	MZN = Currency{Code: "MZN", Precision: 2}
)

var currencies = map[string]Currency{
	"USD": USD,
	"MZN": MZN,
}

func From(code string) (*Currency, error) {

	if currency, ok := currencies[strings.ToUpper(code)]; ok {
		return &currency, nil
	}

	return nil, errors.New("not found")
}
