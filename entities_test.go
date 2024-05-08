package exante

import "testing"

func TestTransactionV3_TradeIsBySymbol(t1 *testing.T) {
	type fields struct {
		SymbolID      string
		Timestamp     int
		UUID          string
		OrderPos      int
		AccountID     string
		ValueDate     string
		ID            int
		Sum           string
		OrderID       string
		OperationType OperationType
		ParentUuid    string
		Asset         string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"#1", fields{SymbolID: "BTC", Asset: "BTC", OperationType: OperationTypeTrade}, true},
		{"#1.1", fields{SymbolID: "BTC", Asset: "USD", OperationType: OperationTypeTrade}, false},
		{"#2", fields{SymbolID: "BTC", Asset: "", OperationType: OperationTypeTrade}, false},
		{"#2.1", fields{SymbolID: "", Asset: "USD", OperationType: OperationTypeTrade}, false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TransactionV3{
				SymbolID:      tt.fields.SymbolID,
				Timestamp:     tt.fields.Timestamp,
				UUID:          tt.fields.UUID,
				OrderPos:      tt.fields.OrderPos,
				AccountID:     tt.fields.AccountID,
				ValueDate:     tt.fields.ValueDate,
				ID:            tt.fields.ID,
				Sum:           tt.fields.Sum,
				OrderID:       tt.fields.OrderID,
				OperationType: tt.fields.OperationType,
				ParentUuid:    tt.fields.ParentUuid,
				Asset:         tt.fields.Asset,
			}
			if got := t.TradeIsBySymbol(); got != tt.want {
				t1.Errorf("TradeIsBySymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
