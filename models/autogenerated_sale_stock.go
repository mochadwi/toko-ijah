package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SaleStockQuerySet

// SaleStockQuerySet is an queryset type for SaleStock
type SaleStockQuerySet struct {
	db *gorm.DB
}

// NewSaleStockQuerySet constructs new SaleStockQuerySet
func NewSaleStockQuerySet(db *gorm.DB) SaleStockQuerySet {
	return SaleStockQuerySet{
		db: db.Model(&SaleStock{}),
	}
}

func (qs SaleStockQuerySet) w(db *gorm.DB) SaleStockQuerySet {
	return NewSaleStockQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) All(ret *[]SaleStock) error {
	return qs.db.Find(ret).Error
}

// AmountDeliveredEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredEq(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered = ?", amountDelivered))
}

// AmountDeliveredGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredGt(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered > ?", amountDelivered))
}

// AmountDeliveredGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredGte(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered >= ?", amountDelivered))
}

// AmountDeliveredIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredIn(amountDelivered uint, amountDeliveredRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{amountDelivered}
	for _, arg := range amountDeliveredRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("amount_delivered IN (?)", iArgs))
}

// AmountDeliveredLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredLt(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered < ?", amountDelivered))
}

// AmountDeliveredLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredLte(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered <= ?", amountDelivered))
}

// AmountDeliveredNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredNe(amountDelivered uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("amount_delivered != ?", amountDelivered))
}

// AmountDeliveredNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) AmountDeliveredNotIn(amountDelivered uint, amountDeliveredRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{amountDelivered}
	for _, arg := range amountDeliveredRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("amount_delivered NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *SaleStock) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtEq(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtGt(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtGte(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtLt(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtLte(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) CreatedAtNe(createdAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *SaleStock) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) Delete() error {
	return qs.db.Delete(SaleStock{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtEq(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtGt(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtGte(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtIsNotNull() SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtIsNull() SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtLt(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtLte(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) DeletedAtNe(deletedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) GetUpdater() SaleStockUpdater {
	return NewSaleStockUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDEq(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDGt(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDGte(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDIn(ID uint, IDRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDLt(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDLte(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDNe(ID uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) IDNotIn(ID uint, IDRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) Limit(limit int) SaleStockQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NameEq(name string) SaleStockQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NameIn(name string, nameRest ...string) SaleStockQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NameNe(name string) SaleStockQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NameNotIn(name string, nameRest ...string) SaleStockQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// NoteEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NoteEq(note string) SaleStockQuerySet {
	return qs.w(qs.db.Where("note = ?", note))
}

// NoteIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NoteIn(note string, noteRest ...string) SaleStockQuerySet {
	iArgs := []interface{}{note}
	for _, arg := range noteRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("note IN (?)", iArgs))
}

// NoteNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NoteNe(note string) SaleStockQuerySet {
	return qs.w(qs.db.Where("note != ?", note))
}

// NoteNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) NoteNotIn(note string, noteRest ...string) SaleStockQuerySet {
	iArgs := []interface{}{note}
	for _, arg := range noteRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("note NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SaleStockQuerySet) One(ret *SaleStock) error {
	return qs.db.First(ret).Error
}

// OrderAscByAmountDelivered is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByAmountDelivered() SaleStockQuerySet {
	return qs.w(qs.db.Order("amount_delivered ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByCreatedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByDeletedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByID() SaleStockQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByProfit is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByProfit() SaleStockQuerySet {
	return qs.w(qs.db.Order("profit ASC"))
}

// OrderAscByPurchasePrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByPurchasePrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("purchase_price ASC"))
}

// OrderAscBySellPrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscBySellPrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("sell_price ASC"))
}

// OrderAscByTime is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByTime() SaleStockQuerySet {
	return qs.w(qs.db.Order("time ASC"))
}

// OrderAscByTotalPrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByTotalPrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("total_price ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderAscByUpdatedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByAmountDelivered is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByAmountDelivered() SaleStockQuerySet {
	return qs.w(qs.db.Order("amount_delivered DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByCreatedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByDeletedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByID() SaleStockQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByProfit is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByProfit() SaleStockQuerySet {
	return qs.w(qs.db.Order("profit DESC"))
}

// OrderDescByPurchasePrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByPurchasePrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("purchase_price DESC"))
}

// OrderDescBySellPrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescBySellPrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("sell_price DESC"))
}

// OrderDescByTime is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByTime() SaleStockQuerySet {
	return qs.w(qs.db.Order("time DESC"))
}

// OrderDescByTotalPrice is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByTotalPrice() SaleStockQuerySet {
	return qs.w(qs.db.Order("total_price DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) OrderDescByUpdatedAt() SaleStockQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// ProfitEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitEq(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit = ?", profit))
}

// ProfitGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitGt(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit > ?", profit))
}

// ProfitGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitGte(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit >= ?", profit))
}

// ProfitIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitIn(profit uint, profitRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{profit}
	for _, arg := range profitRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("profit IN (?)", iArgs))
}

// ProfitLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitLt(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit < ?", profit))
}

// ProfitLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitLte(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit <= ?", profit))
}

// ProfitNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitNe(profit uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("profit != ?", profit))
}

// ProfitNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) ProfitNotIn(profit uint, profitRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{profit}
	for _, arg := range profitRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("profit NOT IN (?)", iArgs))
}

// PurchasePriceEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceEq(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price = ?", purchasePrice))
}

// PurchasePriceGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceGt(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price > ?", purchasePrice))
}

// PurchasePriceGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceGte(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price >= ?", purchasePrice))
}

// PurchasePriceIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceIn(purchasePrice uint, purchasePriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{purchasePrice}
	for _, arg := range purchasePriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("purchase_price IN (?)", iArgs))
}

// PurchasePriceLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceLt(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price < ?", purchasePrice))
}

// PurchasePriceLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceLte(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price <= ?", purchasePrice))
}

// PurchasePriceNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceNe(purchasePrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("purchase_price != ?", purchasePrice))
}

// PurchasePriceNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) PurchasePriceNotIn(purchasePrice uint, purchasePriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{purchasePrice}
	for _, arg := range purchasePriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("purchase_price NOT IN (?)", iArgs))
}

// SKUEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SKUEq(sKU string) SaleStockQuerySet {
	return qs.w(qs.db.Where("sku = ?", sKU))
}

// SKUIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SKUIn(sKU string, sKURest ...string) SaleStockQuerySet {
	iArgs := []interface{}{sKU}
	for _, arg := range sKURest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sku IN (?)", iArgs))
}

// SKUNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SKUNe(sKU string) SaleStockQuerySet {
	return qs.w(qs.db.Where("sku != ?", sKU))
}

// SKUNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SKUNotIn(sKU string, sKURest ...string) SaleStockQuerySet {
	iArgs := []interface{}{sKU}
	for _, arg := range sKURest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sku NOT IN (?)", iArgs))
}

// SellPriceEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceEq(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price = ?", sellPrice))
}

// SellPriceGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceGt(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price > ?", sellPrice))
}

// SellPriceGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceGte(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price >= ?", sellPrice))
}

// SellPriceIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceIn(sellPrice uint, sellPriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{sellPrice}
	for _, arg := range sellPriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sell_price IN (?)", iArgs))
}

// SellPriceLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceLt(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price < ?", sellPrice))
}

// SellPriceLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceLte(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price <= ?", sellPrice))
}

// SellPriceNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceNe(sellPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("sell_price != ?", sellPrice))
}

// SellPriceNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) SellPriceNotIn(sellPrice uint, sellPriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{sellPrice}
	for _, arg := range sellPriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sell_price NOT IN (?)", iArgs))
}

// SetAmountDelivered is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetAmountDelivered(amountDelivered uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.AmountDelivered)] = amountDelivered
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetCreatedAt(createdAt time.Time) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetDeletedAt(deletedAt *time.Time) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetID(ID uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetName(name string) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.Name)] = name
	return u
}

// SetNote is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetNote(note string) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.Note)] = note
	return u
}

// SetProfit is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetProfit(profit uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.Profit)] = profit
	return u
}

// SetPurchasePrice is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetPurchasePrice(purchasePrice uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.PurchasePrice)] = purchasePrice
	return u
}

// SetSKU is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetSKU(sKU string) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.SKU)] = sKU
	return u
}

// SetSellPrice is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetSellPrice(sellPrice uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.SellPrice)] = sellPrice
	return u
}

// SetTime is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetTime(time time.Time) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.Time)] = time
	return u
}

// SetTotalPrice is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetTotalPrice(totalPrice uint) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.TotalPrice)] = totalPrice
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) SetUpdatedAt(updatedAt time.Time) SaleStockUpdater {
	u.fields[string(SaleStockDBSchema.UpdatedAt)] = updatedAt
	return u
}

// TimeEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeEq(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time = ?", time))
}

// TimeGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeGt(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time > ?", time))
}

// TimeGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeGte(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time >= ?", time))
}

// TimeLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeLt(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time < ?", time))
}

// TimeLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeLte(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time <= ?", time))
}

// TimeNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TimeNe(time time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("time != ?", time))
}

// TotalPriceEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceEq(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price = ?", totalPrice))
}

// TotalPriceGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceGt(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price > ?", totalPrice))
}

// TotalPriceGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceGte(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price >= ?", totalPrice))
}

// TotalPriceIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceIn(totalPrice uint, totalPriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{totalPrice}
	for _, arg := range totalPriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("total_price IN (?)", iArgs))
}

// TotalPriceLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceLt(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price < ?", totalPrice))
}

// TotalPriceLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceLte(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price <= ?", totalPrice))
}

// TotalPriceNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceNe(totalPrice uint) SaleStockQuerySet {
	return qs.w(qs.db.Where("total_price != ?", totalPrice))
}

// TotalPriceNotIn is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) TotalPriceNotIn(totalPrice uint, totalPriceRest ...uint) SaleStockQuerySet {
	iArgs := []interface{}{totalPrice}
	for _, arg := range totalPriceRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("total_price NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SaleStockUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtEq(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtGt(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtGte(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtLt(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtLte(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs SaleStockQuerySet) UpdatedAtNe(updatedAt time.Time) SaleStockQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set SaleStockQuerySet

// ===== BEGIN of SaleStock modifiers

type saleStockDBSchemaField string

func (f saleStockDBSchemaField) String() string {
	return string(f)
}

// SaleStockDBSchema stores db field names of SaleStock
var SaleStockDBSchema = struct {
	Time            saleStockDBSchemaField
	ID              saleStockDBSchemaField
	SKU             saleStockDBSchemaField
	Name            saleStockDBSchemaField
	CreatedAt       saleStockDBSchemaField
	UpdatedAt       saleStockDBSchemaField
	DeletedAt       saleStockDBSchemaField
	AmountDelivered saleStockDBSchemaField
	SellPrice       saleStockDBSchemaField
	TotalPrice      saleStockDBSchemaField
	Note            saleStockDBSchemaField
	PurchasePrice   saleStockDBSchemaField
	Profit          saleStockDBSchemaField
}{

	Time:            saleStockDBSchemaField("time"),
	ID:              saleStockDBSchemaField("id"),
	SKU:             saleStockDBSchemaField("sku"),
	Name:            saleStockDBSchemaField("name"),
	CreatedAt:       saleStockDBSchemaField("created_at"),
	UpdatedAt:       saleStockDBSchemaField("updated_at"),
	DeletedAt:       saleStockDBSchemaField("deleted_at"),
	AmountDelivered: saleStockDBSchemaField("amount_delivered"),
	SellPrice:       saleStockDBSchemaField("sell_price"),
	TotalPrice:      saleStockDBSchemaField("total_price"),
	Note:            saleStockDBSchemaField("note"),
	PurchasePrice:   saleStockDBSchemaField("purchase_price"),
	Profit:          saleStockDBSchemaField("profit"),
}

// Update updates SaleStock fields by primary key
func (o *SaleStock) Update(db *gorm.DB, fields ...saleStockDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"time":             o.Time,
		"id":               o.ID,
		"sku":              o.SKU,
		"name":             o.Name,
		"created_at":       o.CreatedAt,
		"updated_at":       o.UpdatedAt,
		"deleted_at":       o.DeletedAt,
		"amount_delivered": o.AmountDelivered,
		"sell_price":       o.SellPrice,
		"total_price":      o.TotalPrice,
		"note":             o.Note,
		"purchase_price":   o.PurchasePrice,
		"profit":           o.Profit,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update SaleStock %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SaleStockUpdater is an SaleStock updates manager
type SaleStockUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSaleStockUpdater creates new SaleStock updater
func NewSaleStockUpdater(db *gorm.DB) SaleStockUpdater {
	return SaleStockUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&SaleStock{}),
	}
}

// ===== END of SaleStock modifiers

// ===== END of all query sets