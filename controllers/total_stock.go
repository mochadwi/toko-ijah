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
		SKU:  utils.GenerateSKU(c.PostForm("size"), c.PostForm("colour")),
		Name: utils.PrettifyName(c.PostForm("name"), c.PostForm("size"), c.PostForm("colour"))}

	totalStock := models.TotalStockRequest{
		StockItem:    stock,
		Size:         c.PostForm("size"),
		Colour:       c.PostForm("colour"),
		CurrentStock: utils.StrToInt64(c.PostForm("currentStock"))}

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
		response.Message = http.StatusText(http.StatusCreated)
		response.Data = totalStock

		c.JSON(http.StatusCreated, response)
	}
}

// GetTotalStock all available list
func GetTotalStock(c *gin.Context) {
	totalStock := models.TotalStockRequest{}
	id := utils.StrToInt(c.Params.ByName("id"))

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.ShowTotalStock(uint(id), &totalStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[totalStock] results: ")
		fmt.Print(totalStock)

		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = totalStock

		c.JSON(http.StatusOK, response)
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
		response.Message = http.StatusText(http.StatusOK)
		response.Data = totalStock

		c.JSON(http.StatusOK, response)
	}
}

/* UpdateTotalStockById all available list
func UpdateTotalStockById(c *gin.Context) {
	var totalStock TotalStockRequest
	id := c.Params.ByName("id")

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err := utils.Mgr.UpdateTotalStockById(id, &totalStock); err != nil {
		response.Code = http.StatusNotFound
		response.Message = err.Error()

		c.JSON(http.StatusNotFound, response)
	} else {
		fmt.Print("[totalStock] results: ")
		fmt.Print(totalStock)

		response.Code = http.StatusAccepted
		response.Message = http.StatusText(http.StatusAccepted)
		response.Data = totalStock

		c.JSON(http.StatusAccepted, response)
	}
}
*/
