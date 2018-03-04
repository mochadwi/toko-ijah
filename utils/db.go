package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

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
	DeleteIncomeStockByID(id uint) (err error)
	GenerateValueReport(valueReport *models.ValueReport) (err error)
	GenerateValueCSV(valueStocks *[]models.ValueStock) error

	// Outcome Stock
	AddOutcomeStock(incomeStock *models.OutcomeStockRequest) error
	ShowOutcomeStock(id uint, incomeStock *models.OutcomeStockRequest) error
	ShowAllOutcomeStock(incomeStock *[]models.OutcomeStockRequest) error
	UpdateOutcomeStockByID(id uint, newOutcomeStock *models.OutcomeStockRequest, currOutcomeStock *models.OutcomeStockRequest) (err error)
	DeleteOutcomeStockByID(id uint) (err error)
	GenerateSaleReport(fromDate string, toDate string, saleReport *models.SaleReport) (err error)
	GenerateSaleCSV(saleStocks *[]models.SaleStock) error
}

type manager struct {
	db            *gorm.DB
	isIniatilized bool
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

	Mgr = &manager{db: db, isIniatilized: true}

	// db.Model(&user).Related(&emails)
	db.Debug().AutoMigrate(
		&models.TotalStockRequest{},
		&models.IncomeStockRequest{},
		&models.OutcomeStockRequest{},
		&models.ValueReport{},
		&models.ValueStock{},
		&models.SaleReport{},
		&models.SaleStock{})
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

	// TODO: update the actual total stock data
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

	// TODO: update the actual total stock data
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

func (mgr *manager) DeleteIncomeStockByID(id uint) (err error) {

	var tempIncomeStock models.IncomeStockRequest

	if err := models.NewIncomeStockRequestQuerySet(mgr.db).IDEq(id).One(&tempIncomeStock); err != nil {
		fmt.Print("[error] showincomestock: ")
		fmt.Println(err)
		return err
	}

	// TODO: update the actual total stock data
	// This method doesn't delete the data, for backup purpose
	// it only update the deleted_at fields
	tempIncomeStock.Delete(mgr.db)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[error] deleteincomestock - delete query: ")
		fmt.Println(err)
		return err
	} // end Delete

	return
}

func (mgr *manager) GenerateValueReport(valueReport *models.ValueReport) (err error) {

	totalStocks := []models.TotalStockRequest{}
	if err = models.NewTotalStockRequestQuerySet(mgr.db).OrderAscByID().All(&totalStocks); err != nil {

		return err
	}

	valueStocks := make([]models.ValueStock, len(totalStocks))

	for i, totalStock := range totalStocks {
		// fmt.Print(strconv.Itoa(i) + "[i]: ")
		// fmt.Println(totalStock)

		var reportAmountReceived = 0
		var reportTotal = 0

		incomeStocks := []models.IncomeStockRequest{}
		if err = models.NewIncomeStockRequestQuerySet(mgr.db).SKUEq(totalStock.SKU).OrderAscByID().All(&incomeStocks); err != nil {

			return err
		}

		for _, incomeStock := range incomeStocks {
			// fmt.Print(strconv.Itoa(j) + "[j] : ")
			// fmt.Println(incomeStock)

			reportAmountReceived += int(incomeStock.AmountReceived)
			reportTotal += int(incomeStock.TotalPrice)
		}

		if len(incomeStocks) > 0 {
			fmt.Print(valueStocks[i].StockItem.SKU)
			fmt.Println(": ")
			fmt.Print("- Final Stock")
			fmt.Println(reportAmountReceived)
			fmt.Print("- Total Prices")
			fmt.Println(reportTotal)

			valueStocks[i].FinalStock = uint(reportAmountReceived)
			valueStocks[i].AvgPurchases = uint(reportTotal / reportAmountReceived)
			valueStocks[i].Total = valueStocks[i].FinalStock * valueStocks[i].AvgPurchases
			valueStocks[i].StockItem = totalStock.StockItem

			valueReport.SKUCount++
			valueReport.StockCount += uint(reportAmountReceived)
			valueReport.TotalStockCount += valueStocks[i].Total
		}

		valueStocks[i].Create(mgr.db)
	}

	valueReport.ValueStock = valueStocks

	// Create
	valueReport.Create(mgr.db)

	// fmt.Print("Value Report: ")
	// fmt.Println(valueReport)

	return
}

func (mgr *manager) GenerateValueCSV(valueStocks *[]models.ValueStock) (err error) {

	if err = models.NewValueStockQuerySet(mgr.db).All(valueStocks); err != nil {

		return err
	}

	return
}

func (mgr *manager) AddOutcomeStock(outcomeStock *models.OutcomeStockRequest) (err error) {

	// Create
	outcomeStock.Create(mgr.db)

	// TODO: update the actual total stock data
	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[error] addoutcomestock - create query: ")
		fmt.Println(err)
		return err
	} // end Create

	return
}

func (mgr *manager) ShowOutcomeStock(id uint, outcomeStock *models.OutcomeStockRequest) (err error) {
	if err := models.NewOutcomeStockRequestQuerySet(mgr.db).IDEq(id).One(outcomeStock); err != nil {
		fmt.Print("[error] showoutcomestock: ")
		fmt.Println(err)
		return err
	}

	// fmt.Print("[success] showoutcomestock: ")
	// fmt.Println(err)
	return
}

func (mgr *manager) ShowAllOutcomeStock(outcomeStock *[]models.OutcomeStockRequest) (err error) {
	if err := models.NewOutcomeStockRequestQuerySet(mgr.db).All(outcomeStock); err != nil {
		fmt.Print("[error] showalloutcomestock: ")
		fmt.Println(err)
		return err
	}
	return
}

func (mgr *manager) UpdateOutcomeStockByID(id uint, newOutcomeStock *models.OutcomeStockRequest, currOutcomeStock *models.OutcomeStockRequest) (err error) {

	if err := models.NewOutcomeStockRequestQuerySet(mgr.db).IDEq(id).One(currOutcomeStock); err != nil {
		fmt.Print("[error] showtotalstock: ")
		fmt.Println(err)
		return err
	}

	fmt.Print("[totalstock]: ")
	fmt.Println(newOutcomeStock)

	fmt.Print("[temptotalstock]: ")
	fmt.Println(currOutcomeStock)

	// TODO: update the actual total stock data
	if !cmp.Equal(currOutcomeStock, newOutcomeStock) {
		// Update
		currOutcomeStock = &models.OutcomeStockRequest{
			StockItem: models.StockItem{
				ID: id},
			AmountDelivered: newOutcomeStock.AmountDelivered,
			SellPrice:       newOutcomeStock.SellPrice,
			TotalPrice:      newOutcomeStock.TotalPrice,
			Note:            newOutcomeStock.Note}

		currOutcomeStock.Update(mgr.db,
			models.OutcomeStockRequestDBSchema.AmountDelivered,
			models.OutcomeStockRequestDBSchema.SellPrice,
			models.OutcomeStockRequestDBSchema.TotalPrice,
			models.OutcomeStockRequestDBSchema.Note)

		if errs := mgr.db.GetErrors(); len(errs) > 0 {
			err = errs[0]
			fmt.Print("[error] updatetotalstock - update query: ")
			fmt.Println(err)
			return err
		} // end Update
	}

	return
}

func (mgr *manager) DeleteOutcomeStockByID(id uint) (err error) {

	var tempOutcomeStock models.OutcomeStockRequest

	if err := models.NewOutcomeStockRequestQuerySet(mgr.db).IDEq(id).One(&tempOutcomeStock); err != nil {
		fmt.Print("[error] showoutcomestock: ")
		fmt.Println(err)
		return err
	}

	// TODO: update the actual total stock data
	// This method doesn't delete the data, for backup purpose
	// it only update the deleted_at fields
	tempOutcomeStock.Delete(mgr.db)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
		fmt.Print("[error] deleteoutcomestock - delete query: ")
		fmt.Println(err)
		return err
	} // end Delete

	return
}

func (mgr *manager) GenerateSaleReport(fromDate string, toDate string, saleReport *models.SaleReport) (err error) {

	if mgr.isIniatilized {
		// mgr.db.Begin()

		// var sku string
		err := mgr.db.
			// rows, err := mgr.db.
			Table("outcome_stock_requests").
			Select("outcome_stock_requests.note, outcome_stock_requests.\"time\", outcome_stock_requests.sku, outcome_stock_requests.name, outcome_stock_requests.amount_delivered, outcome_stock_requests.sell_price, outcome_stock_requests.total_price, income_stock_requests.purchase_price, outcome_stock_requests.total_price - (income_stock_requests.purchase_price * outcome_stock_requests.amount_delivered) AS profit").
			Joins("INNER JOIN income_stock_requests ON outcome_stock_requests.sku = income_stock_requests.sku").
			Where("outcome_stock_requests.time BETWEEN  \"" + fromDate + "\" AND  \"" + toDate + "\"").
			Find(&saleReport.SaleStock).Error
		// Rows()

		if err != nil {
			return err
		}

		for _, sale := range saleReport.SaleStock {
			if sale.OrderID != "" {
				saleReport.TotalSale++
			}
			saleReport.TotalRevenue += sale.TotalPrice
			saleReport.TotalProfit += sale.Profit
			saleReport.TotalStock += sale.AmountDelivered

			sale.Create(mgr.db)
		}

		// dateString := "2016-09-01"
		//convert string to time.Time type
		layOut := "2006-01-02" // yyyy-dd-MM
		// dateStamp, err := time.Parse(layOut, dateString)

		// // we want same format as the dateString
		// // but time.Parse function provides extra information such as time zone, minutes
		// // that we don't need.

		// fmt.Printf("Output(local date) : %s\n", dateStamp.Local())
		// fmt.Printf("Output(UTC) : %s\n", dateStamp)

		// // additional step to format to dd-MMM-yyyy
		// convertedDateString := dateStamp.Format("2-Jan-2006")
		// fmt.Printf("Output : %s\n", convertedDateString)

		// // or if you prefer the full month name
		// fmt.Printf("Full output : %s\n", dateStamp.Format("2-January-2006"))
		saleReport.FromDate, err = time.Parse(layOut, fromDate)
		saleReport.ToDate, err = time.Parse(layOut, toDate)

		if err != nil {
			fmt.Println(err)
			return err
		}

		saleReport.Create(mgr.db)
	}

	return
}

func (mgr *manager) GenerateSaleCSV(saleStocks *[]models.SaleStock) (err error) {

	if err = models.NewSaleStockQuerySet(mgr.db).All(saleStocks); err != nil {

		return err
	}

	return
}
