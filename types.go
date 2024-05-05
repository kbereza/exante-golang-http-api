package exante

type OperationType string

func (t OperationType) String() string {
	return string(t)
}

const (
	OtTypeTrade      OperationType = "TRADE"
	OtTypeFundWith   OperationType = "FUNDING/WITHDRAWAL"
	OtTypeCommission OperationType = "COMMISSION"
	OtTypeInterest   OperationType = "INTEREST"
)
