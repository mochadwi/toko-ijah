package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/mochadwi/toko-ijah/models"
	utils "github.com/mochadwi/toko-ijah/utils"
	index "github.com/mochadwi/toko-ijah/views"
	uuid "github.com/satori/go.uuid"
)

// AddOutcomeStock will add new outcome stocks
func AddOutcomeStock(c *gin.Context) {

	// AmountDelivered uint   `json:"amountDelivered"`
	// SellPrice       uint   `json:"sellPrice"`
	// TotalPrice      uint   `json:"totalPrice"`
	// Note            string `json:"note"`

	stock := models.StockItem{
		SKU:  c.PostForm("sku"),
		Name: c.PostForm("name")}

	date := time.Now()
	var amountDelivered = utils.StrToUint(c.PostForm("amountDelivered"))
	var sellPrice = utils.StrToUint(c.PostForm("sellPrice"))

	outcomeStock := models.OutcomeStockRequest{
		Time:            date,
		StockItem:       stock,
		AmountDelivered: amountDelivered,
		SellPrice:       sellPrice,
		TotalPrice:      amountDelivered * sellPrice,
		Note:            c.PostForm("note")}

	fmt.Println(outcomeStock)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.AddOutcomeStock(&outcomeStock); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusCreated
		response.Message = http.StatusText(http.StatusCreated)
		response.Data = outcomeStock

		c.JSON(http.StatusCreated, response)
	}
}

// GetOutcomeStock record
func GetOutcomeStock(c *gin.Context) {
	outcomeStock := models.OutcomeStockRequest{}
	id := utils.StrToUint(c.Params.ByName("id"))

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowOutcomeStock(id, &outcomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[outcomeStock] results: ")
		fmt.Print(outcomeStock)

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = outcomeStock

		c.JSON(http.StatusOK, response)
	}
}

// GetAllOutcomeStock all available list
func GetAllOutcomeStock(c *gin.Context) {
	outcomeStock := []models.OutcomeStockRequest{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowAllOutcomeStock(&outcomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[outcomeStock] results: ")
		fmt.Print(outcomeStock)

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = outcomeStock

		c.JSON(http.StatusOK, response)
	}
}

// UpdateOutcomeStockByID will update the data by ID
func UpdateOutcomeStockByID(c *gin.Context) {
	var currOutcomeStock models.OutcomeStockRequest
	id := utils.StrToUint(c.Params.ByName("id"))

	newOutcomeStock := models.OutcomeStockRequest{
		AmountDelivered: utils.StrToUint(c.PostForm("amountDelivered")),
		SellPrice:       utils.StrToUint(c.PostForm("sellPrice")),
		TotalPrice:      utils.StrToUint(c.PostForm("totalPrice")),
		Note:            c.PostForm("note")}

	fmt.Println(newOutcomeStock)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.UpdateOutcomeStockByID(id, &newOutcomeStock, &currOutcomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[totalStock] results: ")
		fmt.Print(currOutcomeStock)

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = currOutcomeStock

		c.JSON(http.StatusAccepted, response)
	}
}

// DeleteOutcomeStockByID will delete the data by ID
func DeleteOutcomeStockByID(c *gin.Context) {
	id := utils.StrToUint(c.Params.ByName("id"))

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.DeleteOutcomeStockByID(id); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = nil

		c.JSON(http.StatusAccepted, response)
	}
}

// GenerateSaleReport ...
func GenerateSaleReport(c *gin.Context) {
	saleReport := models.SaleReport{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	fromDate := "2018-03-04 09:38:44.849835114+07:00"
	toDate := "2018-03-14 09:38:44.849835114+07:00"

	if err := utils.Mgr.GenerateSaleReport(fromDate, toDate, &saleReport); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = saleReport

		c.JSON(http.StatusAccepted, response)
	}
}

// // GenerateSaleCSV ...
// func GenerateSaleCSV(c *gin.Context) {
// 	valueStocks := []models.SaleStock{}

// 	var response = &index.DefaultResponseFormat{
// 		RequestID: uuid.NewV4().String(),
// 		Now:       time.Now().Unix(),
// 	}

// 	if err := utils.Mgr.GenerateSaleCSV(&valueStocks); err != nil {
// 		response.Code = http.StatusNotFound
// 		response.Message = err.Error()

// 		c.JSON(http.StatusNotFound, response)
// 	} else {

// 		b := &bytes.Buffer{}
// 		w := csv.NewWriter(b)

// 		w.Write([]string{
// 			"SKU",
// 			"Name",
// 			"Quantity",
// 			"Average Purchases",
// 			"Total",
// 		})

// 		for _, valueStock := range valueStocks {
// 			w.Write([]string{
// 				valueStock.SKU,
// 				valueStock.Name,
// 				utils.UintToStr(valueStock.FinalStock),
// 				utils.PrettifyPrice("IDR", valueStock.AvgPurchases),
// 				utils.PrettifyPrice("IDR", valueStock.Total)})
// 		}
// 		w.Flush()

// 		if err := w.Error(); err != nil {
// 			log.Fatal(err)
// 		}

// 		response.Code = http.StatusOK
// 		response.Message = http.StatusText(http.StatusOK)
// 		response.Data = valueStocks

// 		date := time.Now().Local()

// 		c.Header("Content-Description", "File Transfer")
// 		c.Header("Content-Disposition", "attachment; filename="+date.Format("2006-01-02")+"_value_stocks.csv")
// 		c.Data(http.StatusOK, "text/csv", b.Bytes())
// 	}
// }
