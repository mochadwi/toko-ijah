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
	AddTotalStock(totalStock *models.TotalStockRequest) error
	ShowAllTotalStock(totalStock *[]models.TotalStockRequest) error
	ShowTotalStock(id uint, totalStock *models.TotalStockRequest) error
	// UpdateTotalStockById(id uint, totalStock *models.TotalStockRequest) (err error)
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
		&models.TotalStockRequest{})
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
		fmt.Print("[error] showallnotifier: ")
		fmt.Println(err)
		return err
	}
	return
}

// func (mgr *manager) UpdateTotalStockById(id uint, totalStock *models.TotalStockRequest) (err error) {

// 	if err := models.NewTotalStockRequestUpdater(mgr.db).

// 	mgr.db.Save()
// 	if err := mgr.db.Where("id = ?", id).First(totalStock).Error; err != nil {

// 		c.AbortWithStatus(404)
// 		fmt.Println(err)

// 		return err
// 	}

// 	c.BindJSON(&person)
// 	db.Save(&person)
// 	c.JSON(200, person)
// }
