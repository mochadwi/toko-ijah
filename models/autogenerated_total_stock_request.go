package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set TotalStockRequestQuerySet

// TotalStockRequestQuerySet is an queryset type for TotalStockRequest
type TotalStockRequestQuerySet struct {
	db *gorm.DB
}

// NewTotalStockRequestQuerySet constructs new TotalStockRequestQuerySet
func NewTotalStockRequestQuerySet(db *gorm.DB) TotalStockRequestQuerySet {
	return TotalStockRequestQuerySet{
		db: db.Model(&TotalStockRequest{}),
	}
}

func (qs TotalStockRequestQuerySet) w(db *gorm.DB) TotalStockRequestQuerySet {
	return NewTotalStockRequestQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) All(ret *[]TotalStockRequest) error {
	return qs.db.Find(ret).Error
}

// ColourEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) ColourEq(colour string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("colour = ?", colour))
}

// ColourIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) ColourIn(colour string, colourRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{colour}
	for _, arg := range colourRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("colour IN (?)", iArgs))
}

// ColourNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) ColourNe(colour string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("colour != ?", colour))
}

// ColourNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) ColourNotIn(colour string, colourRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{colour}
	for _, arg := range colourRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("colour NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *TotalStockRequest) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtEq(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtGt(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtGte(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtLt(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtLte(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CreatedAtNe(createdAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CurrentStockEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockEq(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock = ?", currentStock))
}

// CurrentStockGt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockGt(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock > ?", currentStock))
}

// CurrentStockGte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockGte(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock >= ?", currentStock))
}

// CurrentStockIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockIn(currentStock int64, currentStockRest ...int64) TotalStockRequestQuerySet {
	iArgs := []interface{}{currentStock}
	for _, arg := range currentStockRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("current_stock IN (?)", iArgs))
}

// CurrentStockLt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockLt(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock < ?", currentStock))
}

// CurrentStockLte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockLte(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock <= ?", currentStock))
}

// CurrentStockNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockNe(currentStock int64) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("current_stock != ?", currentStock))
}

// CurrentStockNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) CurrentStockNotIn(currentStock int64, currentStockRest ...int64) TotalStockRequestQuerySet {
	iArgs := []interface{}{currentStock}
	for _, arg := range currentStockRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("current_stock NOT IN (?)", iArgs))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *TotalStockRequest) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) Delete() error {
	return qs.db.Delete(TotalStockRequest{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtEq(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtGt(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtGte(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtIsNotNull() TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtIsNull() TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtLt(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtLte(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) DeletedAtNe(deletedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) GetUpdater() TotalStockRequestUpdater {
	return NewTotalStockRequestUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDEq(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDGt(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDGte(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDIn(ID uint, IDRest ...uint) TotalStockRequestQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDLt(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDLte(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDNe(ID uint) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) IDNotIn(ID uint, IDRest ...uint) TotalStockRequestQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) Limit(limit int) TotalStockRequestQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) NameEq(name string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) NameIn(name string, nameRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) NameNe(name string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) NameNotIn(name string, nameRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs TotalStockRequestQuerySet) One(ret *TotalStockRequest) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderAscByCreatedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByCurrentStock is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderAscByCurrentStock() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("current_stock ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderAscByDeletedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderAscByID() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderAscByUpdatedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderDescByCreatedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByCurrentStock is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderDescByCurrentStock() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("current_stock DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderDescByDeletedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderDescByID() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) OrderDescByUpdatedAt() TotalStockRequestQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SKUEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SKUEq(sKU string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("sku = ?", sKU))
}

// SKUIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SKUIn(sKU string, sKURest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{sKU}
	for _, arg := range sKURest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sku IN (?)", iArgs))
}

// SKUNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SKUNe(sKU string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("sku != ?", sKU))
}

// SKUNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SKUNotIn(sKU string, sKURest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{sKU}
	for _, arg := range sKURest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("sku NOT IN (?)", iArgs))
}

// SetColour is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetColour(colour string) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.Colour)] = colour
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetCreatedAt(createdAt time.Time) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCurrentStock is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetCurrentStock(currentStock int64) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.CurrentStock)] = currentStock
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetDeletedAt(deletedAt *time.Time) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetID(ID uint) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetName(name string) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.Name)] = name
	return u
}

// SetSKU is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetSKU(sKU string) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.SKU)] = sKU
	return u
}

// SetSize is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetSize(size string) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.Size)] = size
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) SetUpdatedAt(updatedAt time.Time) TotalStockRequestUpdater {
	u.fields[string(TotalStockRequestDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SizeEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SizeEq(size string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("size = ?", size))
}

// SizeIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SizeIn(size string, sizeRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{size}
	for _, arg := range sizeRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("size IN (?)", iArgs))
}

// SizeNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SizeNe(size string) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("size != ?", size))
}

// SizeNotIn is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) SizeNotIn(size string, sizeRest ...string) TotalStockRequestQuerySet {
	iArgs := []interface{}{size}
	for _, arg := range sizeRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("size NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u TotalStockRequestUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtEq(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtGt(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtGte(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtLt(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtLte(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs TotalStockRequestQuerySet) UpdatedAtNe(updatedAt time.Time) TotalStockRequestQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set TotalStockRequestQuerySet

// ===== BEGIN of TotalStockRequest modifiers

type totalStockRequestDBSchemaField string

func (f totalStockRequestDBSchemaField) String() string {
	return string(f)
}

// TotalStockRequestDBSchema stores db field names of TotalStockRequest
var TotalStockRequestDBSchema = struct {
	ID           totalStockRequestDBSchemaField
	SKU          totalStockRequestDBSchemaField
	Name         totalStockRequestDBSchemaField
	CurrentStock totalStockRequestDBSchemaField
	CreatedAt    totalStockRequestDBSchemaField
	UpdatedAt    totalStockRequestDBSchemaField
	DeletedAt    totalStockRequestDBSchemaField
	Size         totalStockRequestDBSchemaField
	Colour       totalStockRequestDBSchemaField
}{

	ID:           totalStockRequestDBSchemaField("id"),
	SKU:          totalStockRequestDBSchemaField("sku"),
	Name:         totalStockRequestDBSchemaField("name"),
	CurrentStock: totalStockRequestDBSchemaField("current_stock"),
	CreatedAt:    totalStockRequestDBSchemaField("created_at"),
	UpdatedAt:    totalStockRequestDBSchemaField("updated_at"),
	DeletedAt:    totalStockRequestDBSchemaField("deleted_at"),
	Size:         totalStockRequestDBSchemaField("size"),
	Colour:       totalStockRequestDBSchemaField("colour"),
}

// Update updates TotalStockRequest fields by primary key
func (o *TotalStockRequest) Update(db *gorm.DB, fields ...totalStockRequestDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":            o.ID,
		"sku":           o.SKU,
		"name":          o.Name,
		"current_stock": o.CurrentStock,
		"created_at":    o.CreatedAt,
		"updated_at":    o.UpdatedAt,
		"deleted_at":    o.DeletedAt,
		"size":          o.Size,
		"colour":        o.Colour,
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

		return fmt.Errorf("can't update TotalStockRequest %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// TotalStockRequestUpdater is an TotalStockRequest updates manager
type TotalStockRequestUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewTotalStockRequestUpdater creates new TotalStockRequest updater
func NewTotalStockRequestUpdater(db *gorm.DB) TotalStockRequestUpdater {
	return TotalStockRequestUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&TotalStockRequest{}),
	}
}

// ===== END of TotalStockRequest modifiers

// ===== END of all query sets
