package currency

import "errors"

type Currency struct {
	Code      string
	Precision int
}

var (
	USD = Currency{Code: "USD", Precision: 2}
	MZN = Currency{Code: "MZN", Precision: 2}
)

func From(code string) (*Currency, error) {
	currencies := map[string]Currency{
		"USD": USD,
		"MZN": MZN,
	}

	if currency, ok := currencies[code]; ok {
		return &currency, nil
	}

	return nil, errors.New("not found")
}
