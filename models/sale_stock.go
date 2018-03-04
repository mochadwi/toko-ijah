package models

//go:generate goqueryset -in sale_stock.go

// SaleStock struct represent sale_stock model. Next line (gen:qs) is needed to autogenerate SaleStockQuerySet.
// gen:qs
type SaleStock struct {
	OutcomeStockRequest
	PurchasePrice uint `json:"purchasePrice"`
	Profit        uint `json:"profit"`
}
