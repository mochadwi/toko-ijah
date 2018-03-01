package models

//go:generate goqueryset -in total_stock_request.go

// TotalStockRequest struct represent total stock model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type TotalStockRequest struct {
	StockItem
	Size   string `json:"size"`
	Colour string `json:"colour"`
}
