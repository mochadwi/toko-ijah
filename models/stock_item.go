package models

import (
	"time"
)

//go:generate goqueryset -in stock_item.go

// StockItem struct represent stock model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type StockItem struct {
	ID           uint   `json:"id"` // gorm primary
	SKU          string `json:"sku"`
	Name         string `json:"name"`
	CurrentStock int64  `json:"currentStock"`
	// gorm model
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
