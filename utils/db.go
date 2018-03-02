package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mochadwi/toko-ijah/models"
)

// Manager interface
type Manager interface {
	// Total Stock
	AddTotalStock(totalStock *models.TotalStockRequest) error
	ShowAllTotalStock(totalStock *[]models.TotalStockRequest) error
	ShowTotalStock(id uint, totalStock *models.TotalStockRequest) error
	UpdateTotalStockByID(id uint, newTotalStock *models.TotalStockRequest, currTotalStock *models.TotalStockRequest) (err error)
	DeleteTotalStockByID(id uint) (err error)

	// Income Stock
	AddIncomeStock(incomeStock *models.IncomeStockRequest) error
	ShowIncomeStock(id uint, incomeStock *models.IncomeStockRequest) error
	ShowAllIncomeStock(incomeStock *[]models.IncomeStockRequest) error
	UpdateIncomeStockByID(id uint, newIncomeStock *models.IncomeStockRequest, currIncomeStock *models.IncomeStockRequest) (err error)
}

type manager struct {
	db *gorm.DB
}

// Mgr to manage database
var Mgr Manager

func init() {
	db, err := gorm.Open("sqlite3", "./toko_ijah.db")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	db.LogMode(true)
	// defer db.Close()

	Mgr = &manager{db: db}

	db.Debug().AutoMigrate(
		&models.TotalStockRequest{},
		&models.IncomeStockRequest{})
}

func (mgr *manager) AddTotalStock(totalStock *models.TotalStockRequest) (err error) {

	var tempTotalStock models.TotalStockRequest
	models.NewTotalStockRequestQuerySet(mgr.db).NameEq(totalStock.Name).One(&tempTotalStock)

	if tempTotalStock.Name != totalStock.Name {
		// Create
		totalStock.Create(mgr.db)

		if errs := mgr.db.GetErrors(); len(errs) > 0 {
			err = errs[0]
			fmt.Print("[error] addtotalstock - create query: ")
			fmt.Println(err)
			return err
		} // end Create

		return
	}

	fmt.Print("[error] addtotalstock - duplicate found: ")
	fmt.Println("duplicate entry")
	return fmt.Errorf("%s", "duplicate entry")
}

func (mgr *manager) ShowTotalStock(id uint, totalStock *models.TotalStockRequest) (err error) {
	if err := models.NewTotalStockRequestQuerySet(mgr.db).IDEq(id).One(totalStock); err != nil {
		fmt.Print("[error] showtotalstock: ")
		fmt.Println(err)
		return err
	}

	// fmt.Print("[success] showtotalstock: ")
	// fmt.Println(err)
	return
}

func (mgr *manager) ShowAllTotalStock(totalStock *[]models.TotalStockRequest) (err error) {
	if err := models.NewTotalStockRequestQuerySet(mgr.db).All(totalStock); err != nil {
		fmt.Print("[error] showalltotalstock: ")
		fmt.Println(err)
		return err
	}
	return
}

func (mgr *manager) UpdateTotalStockByID(id uint, newTotalStock *models.TotalStockRequest, currTotalStock *models.TotalStockRequest) (err error) {

	if err := models.NewTotalStockRequestQuerySet(mgr.db).IDEq(id).One(currTotalStock); err != nil {
		fmt.Print("[error] showtotalstock: ")
		fmt.Println(err)
		return err
	}

	fmt.Print("[totalstock]: ")
	fmt.Println(newTotalStock)

	fmt.Print("[temptotalstock]: ")
	fmt.Println(currTotalStock)

	if !cmp.Equal(currTotalStock, newTotalStock) {
		// Update
		var oldStr = currTotalStock.SKU[13:len(currTotalStock.SKU)]
		var newStr = newTotalStock.SKU[13:len(newTotalStock.SKU)]

		currTotalStock = &models.TotalStockRequest{
			StockItem: models.StockItem{
				ID:   id,
				SKU:  strings.Replace(currTotalStock.SKU, oldStr, newStr, -1),
				Name: newTotalStock.Name},
			Size:         newTotalStock.Size,
			Colour:       newTotalStock.Colour,
			CurrentStock: newTotalStock.CurrentStock}

		currTotalStock.Update(mgr.db,
			models.TotalStockRequestDBSchema.SKU,
			models.TotalStockRequestDBSchema.Name,
			models.TotalStockRequestDBSchema.Size,
			models.TotalStockRequestDBSchema.Colour,
			models.TotalStockRequestDBSchema.CurrentStock)

		if errs := mgr.db.GetErrors(); len(errs) > 0 {
			err = errs[0]
			fmt.Print("[error] updatetotalstock - update query: ")
			fmt.Println(err)
			return err
		} // end Update
	}

	return
}

func (mgr *manager) DeleteTotalStockByID(id uint) (err error) {

	var tempTotalStock models.TotalStockRequest

	if err := models.NewTotalStockRequestQuerySet(mgr.db).IDEq(id).One(&tempTotalStock); err != nil {
		fmt.Print("[error] showtotalstock: ")
		fmt.Println(err)
		return err
	}

	// This method doesn't delete the data, for backup purpose
	// it only update the deleted_at fields
	tempTotalStock.Delete(mgr.db)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[error] deletetotalstock - delete query: ")
		fmt.Println(err)
		return err
	} // end Delete

	return
}

func (mgr *manager) AddIncomeStock(incomeStock *models.IncomeStockRequest) (err error) {

	// Create
	incomeStock.Create(mgr.db)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[error] addincomestock - create query: ")
		fmt.Println(err)
		return err
	} // end Create

	return
}

func (mgr *manager) ShowIncomeStock(id uint, incomeStock *models.IncomeStockRequest) (err error) {
	if err := models.NewIncomeStockRequestQuerySet(mgr.db).IDEq(id).One(incomeStock); err != nil {
		fmt.Print("[error] showincomestock: ")
		fmt.Println(err)
		return err
	}

	// fmt.Print("[success] showincomestock: ")
	// fmt.Println(err)
	return
}

func (mgr *manager) ShowAllIncomeStock(incomeStock *[]models.IncomeStockRequest) (err error) {
	if err := models.NewIncomeStockRequestQuerySet(mgr.db).All(incomeStock); err != nil {
		fmt.Print("[error] showallincomestock: ")
		fmt.Println(err)
		return err
	}
	return
}

func (mgr *manager) UpdateIncomeStockByID(id uint, newIncomeStock *models.IncomeStockRequest, currIncomeStock *models.IncomeStockRequest) (err error) {

	if err := models.NewIncomeStockRequestQuerySet(mgr.db).IDEq(id).One(currIncomeStock); err != nil {
		fmt.Print("[error] showtotalstock: ")
		fmt.Println(err)
		return err
	}

	fmt.Print("[totalstock]: ")
	fmt.Println(newIncomeStock)

	fmt.Print("[temptotalstock]: ")
	fmt.Println(currIncomeStock)

	if !cmp.Equal(currIncomeStock, newIncomeStock) {
		// Update
		currIncomeStock = &models.IncomeStockRequest{
			StockItem: models.StockItem{
				ID: id},
			AmountReceived: newIncomeStock.AmountReceived,
			TotalPrice:     newIncomeStock.TotalPrice,
			ReceiptNumber:  newIncomeStock.ReceiptNumber,
			Note:           newIncomeStock.Note}

		currIncomeStock.Update(mgr.db,
			models.IncomeStockRequestDBSchema.AmountReceived,
			models.IncomeStockRequestDBSchema.TotalPrice,
			models.IncomeStockRequestDBSchema.ReceiptNumber,
			models.IncomeStockRequestDBSchema.Note)

		if errs := mgr.db.GetErrors(); len(errs) > 0 {
			err = errs[0]
			fmt.Print("[error] updatetotalstock - update query: ")
			fmt.Println(err)
			return err
		} // end Update
	}

	return
}
