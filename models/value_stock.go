package models

//go:generate goqueryset -in value_stock.go

// ValueStock struct represent value_stock model. Next line (gen:qs) is needed to autogenerate ValueStockQuerySet.
// gen:qs
type ValueStock struct {
	StockItem
	FinalStock   uint `json:"finalStock"`
	AvgPurchases uint `json:"avgPurchases"`
	Total        uint `json:"total"`
}
