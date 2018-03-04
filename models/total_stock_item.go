package models

import (
	"time"
)

//go:generate goqueryset -in total_stock_item.go

// TotalStockItem struct represent totak stock model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type TotalStockItem struct {
	ID           uint   `json:"id"` // gorm primary
	SKU          string `json:"sku"`
	Name         string `json:"name"`
	Size         string `json:"size"`
	Colour       string `json:"colour"`
	CurrentStock int64  `json:"currentStock"`
	// gorm model
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
