package models

import "time"

//go:generate goqueryset -in sale_report.go

// SaleReport struct represent sale model. Next line (gen:qs) is needed to autogenerate SaleReportQuerySet.
// gen:qs
type SaleReport struct {
	FromDate     *time.Time  `json:"fromDate"`
	ToDate       *time.Time  `json:"toDate"`
	TotalRevenue uint        `json:"totalRevenue"`
	TotalProfit  uint        `json:"totalProfit"`
	TotalSale    uint        `json:"totalSale"`
	TotalStock   uint        `json:"totalStock"`
	SaleStock    []SaleStock `json:"saleStock"`
}
