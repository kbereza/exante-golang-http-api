package exante

type OperationType string

func (t OperationType) String() string {
	return string(t)
}

const (
	OperationTypeTrade      OperationType = "TRADE"
	OperationTypeFundWith   OperationType = "FUNDING/WITHDRAWAL"
	OperationTypeCommission OperationType = "COMMISSION"
	OperationTypeInterest   OperationType = "INTEREST"
)
