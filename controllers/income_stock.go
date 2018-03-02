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
