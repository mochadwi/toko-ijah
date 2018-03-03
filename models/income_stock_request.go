package models

import "time"

//go:generate goqueryset -in income_stock_request.go

// IncomeStockRequest struct represent income stock model. Next line (gen:qs) is needed to autogenerate IncomeStockQuerySet.
// gen:qs
type IncomeStockRequest struct {
	Time time.Time `json:"time"`
	StockItem
	OrderAmount    uint   `json:"orderAmount"`
	AmountReceived uint   `json:"amountReceived"`
	PurchasePrice  uint   `json:"purchasePrice"`
	TotalPrice     uint   `json:"totalPrice"`
	ReceiptNumber  string `json:"receiptNumber"`
	Note           string `json:"note"`
}
