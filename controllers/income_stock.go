package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/mochadwi/toko-ijah/models"
	utils "github.com/mochadwi/toko-ijah/utils"
	index "github.com/mochadwi/toko-ijah/views"
	uuid "github.com/satori/go.uuid"
)

// AddIncomeStock will add new income stocks
func AddIncomeStock(c *gin.Context) {

	stock := models.StockItem{
		SKU:  c.PostForm("sku"),
		Name: c.PostForm("name")}

	date := time.Now()
	var amountReceived = utils.StrToUint(c.PostForm("amountReceived"))
	var purchasePrice = utils.StrToUint(c.PostForm("purchasePrice"))

	incomeStock := models.IncomeStockRequest{
		Time:           date,
		StockItem:      stock,
		OrderAmount:    utils.StrToUint(c.PostForm("orderAmount")),
		AmountReceived: amountReceived,
		PurchasePrice:  purchasePrice,
		TotalPrice:     amountReceived * purchasePrice,
		ReceiptNumber:  utils.GenerateReceipt(date.Format("20060102")),
		Note:           c.PostForm("note")}

	fmt.Println(incomeStock)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.AddIncomeStock(&incomeStock); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusCreated
		response.Message = http.StatusText(http.StatusCreated)
		response.Data = incomeStock

		c.JSON(http.StatusCreated, response)
	}
}

// GetIncomeStock record
func GetIncomeStock(c *gin.Context) {
	incomeStock := models.IncomeStockRequest{}
	id := utils.StrToInt(c.Params.ByName("id"))

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowIncomeStock(uint(id), &incomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[incomeStock] results: ")
		fmt.Print(incomeStock)

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = incomeStock

		c.JSON(http.StatusOK, response)
	}
}

// GetAllIncomeStock all available list
func GetAllIncomeStock(c *gin.Context) {
	incomeStock := []models.IncomeStockRequest{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowAllIncomeStock(&incomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[incomeStock] results: ")
		fmt.Print(incomeStock)

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = incomeStock

		c.JSON(http.StatusOK, response)
	}
}

// UpdateIncomeStockByID will update the data by ID
func UpdateIncomeStockByID(c *gin.Context) {
	var currIncomeStock models.IncomeStockRequest
	id := utils.StrToUint(c.Params.ByName("id"))

	var amountReceived = utils.StrToUint(c.PostForm("amountReceived"))

	newIncomeStock := models.IncomeStockRequest{
		AmountReceived: amountReceived,
		TotalPrice:     utils.StrToUint(c.PostForm("totalPrice")),
		ReceiptNumber:  c.PostForm("receiptNumber"),
		Note:           c.PostForm("note")}

	fmt.Println(newIncomeStock)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.UpdateIncomeStockByID(id, &newIncomeStock, &currIncomeStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[totalStock] results: ")
		fmt.Print(currIncomeStock)

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = currIncomeStock

		c.JSON(http.StatusAccepted, response)
	}
}

// DeleteIncomeStockByID will delete the data by ID
func DeleteIncomeStockByID(c *gin.Context) {
	id := utils.StrToUint(c.Params.ByName("id"))

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.DeleteIncomeStockByID(id); err != nil {
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

// GenerateValueReport will delete the data by ID
func GenerateValueReport(c *gin.Context) {
	valueReport := models.ValueReport{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.GenerateValueReport(&valueReport); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = valueReport

		c.JSON(http.StatusAccepted, response)
	}
}

// GenerateValueCSV ...
func GenerateValueCSV(c *gin.Context) {
	valueStocks := []models.ValueStock{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.GenerateValueCSV(&valueStocks); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {

		b := &bytes.Buffer{}
		w := csv.NewWriter(b)

		w.Write([]string{
			"SKU",
			"Name",
			"Quantity",
			"Average Purchases",
			"Total",
		})

		for _, valueStock := range valueStocks {
			w.Write([]string{
				valueStock.SKU,
				valueStock.Name,
				utils.UintToStr(valueStock.FinalStock),
				utils.UintToStr(valueStock.AvgPurchases),
				utils.UintToStr(valueStock.Total)})
		}
		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = valueStocks

		date := time.Now().Local()

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+date.Format("2006-01-02")+"_value_stocks.csv")
		c.Data(http.StatusOK, "text/csv", b.Bytes())
	}
}
