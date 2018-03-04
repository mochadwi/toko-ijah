package models

//go:generate goqueryset -in value_report.go

// ValueReport struct represent value model. Next line (gen:qs) is needed to autogenerate ValueReportQuerySet.
// gen:qs
type ValueReport struct {
	SKUCount        uint         `json:"skuCount"`
	StockCount      uint         `json:"stockCount"`
	TotalStockCount uint         `json:"totalStockCount"`
	ValueStock      []ValueStock `json:"valueStock"`
}
