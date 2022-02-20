package card

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
