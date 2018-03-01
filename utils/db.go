package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mochadwi/toko-ijah/models"
)

// Manager interface
type Manager interface {
	AddTotalStock(totalStock *models.TotalStockItem) error
	ShowAllTotalStock(totalStock *[]models.TotalStockItem) error
	ShowTotalStock(name string, totalStock *models.TotalStockItem) error
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
		&models.TotalStockItem{})
}

func (mgr *manager) AddTotalStock(totalStock *models.TotalStockItem) (err error) {

	var tempTotalStock models.TotalStockItem
	models.NewTotalStockItemQuerySet(mgr.db).NameEq(totalStock.Name).One(&tempTotalStock)

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

func (mgr *manager) ShowAllTotalStock(totalStock *[]models.TotalStockItem) (err error) {
	if err := models.NewTotalStockItemQuerySet(mgr.db).All(totalStock); err != nil {
		fmt.Print("[error] showallnotifier: ")
		fmt.Println(err)
		return err
	}
	return
}

func (mgr *manager) ShowTotalStock(name string, totalStock *models.TotalStockItem) (err error) {
	if err := models.NewTotalStockItemQuerySet(mgr.db).NameEq(name).One(totalStock); err != nil {
		fmt.Print("[error] shownotifier: ")
		fmt.Println(err)
		return err
	}

	// fmt.Print("[success] shownotifier: ")
	// fmt.Println(err)
	return
}
