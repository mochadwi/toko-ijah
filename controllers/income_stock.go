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
