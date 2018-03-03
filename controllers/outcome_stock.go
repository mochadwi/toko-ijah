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
