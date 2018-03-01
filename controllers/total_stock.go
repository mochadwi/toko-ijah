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

// AddTotalStock will add new total stocks
func AddTotalStock(c *gin.Context) {
	// "name": "Zalekia Plain Casual Blouse",
	// "size": "L",
	// "colour": "Broken White",
	// "currentStock": 100

	stock := models.StockItem{
		SKU:          utils.GenerateSKU(c.PostForm("size"), c.PostForm("colour")),
		Name:         c.PostForm("name"),
		CurrentStock: utils.StrToInt(c.PostForm("currentStock"))}

	totalStock := models.TotalStockRequest{
		StockItem: stock,
		Size:      c.PostForm("size"),
		Colour:    c.PostForm("colour")}

	fmt.Println(totalStock)

	err := utils.Mgr.AddTotalStock(&totalStock)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusCreated
		response.Message = "OK"
		response.Data = totalStock

		c.JSON(http.StatusCreated, response)
	}
}

// GetAllTotalStock all available list
func GetAllTotalStock(c *gin.Context) {
	totalStock := []models.TotalStockRequest{}

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowAllTotalStock(&totalStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[totalStock] results: ")
		fmt.Print(totalStock)

		response.Code = http.StatusOK
		response.Message = "OK"
		response.Data = totalStock

		c.JSON(http.StatusOK, response)
	}
}
