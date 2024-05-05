package exante

import (
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	h := NewAPI(baseAPIURL,
		"3.0",
		"f9b1cfd2-584c-4bf7-a236-72fc3d20ccfb",
		"289ca5f5-e2b8-4cb6-8356-3fb6c765bc42",
		"is11cEQ4IZf3nL9DrdxCJHGGcBbjJd9t",
		300, "", "",
	)

	/*accounts, err := h.GetUserAccounts()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("accounts", accounts)

	order, err := h.GetOrderV3("e6014e43-23f3-45d9-966b-b91cb793b47f")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("order", order)*/

	fmt.Println(h.GetLastQuote("AAPL.NASDAQ"))
}

// https://api-demo.exante.eu/trade/3.0/orders/e6014e43-23f3-45d9-966b-b91cb793b47f
// https://api-demo.exante.eu/trade/3.0/orders/e6014e43-23f3-45d9-966b-b91cb793b47f/
