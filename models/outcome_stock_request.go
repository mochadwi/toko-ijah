package models

import "time"

//go:generate goqueryset -in outcome_stock_request.go

// OutcomeStockRequest struct represent outcome stock model. Next line (gen:qs) is needed to autogenerate OutcomeStockQuerySet.
// gen:qs
type OutcomeStockRequest struct {
	Time time.Time `json:"time"`
	StockItem
	AmountDelivered uint   `json:"amountDelivered"`
	SellPrice       uint   `json:"sellPrice"`
	TotalPrice      uint   `json:"totalPrice"`
	Note            string `json:"note"`
}
