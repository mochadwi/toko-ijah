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

	data := models.TotalStockItem{
		SKU:          utils.GenerateSKU(c.PostForm("size"), c.PostForm("colour")),
		Name:         utils.PrettifyName(c.PostForm("name"), c.PostForm("size"), c.PostForm("colour")),
		CurrentStock: utils.StrToInt(c.PostForm("currentStock"))}

	fmt.Println(data)

	var response = &index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	// if data != nil {
	response.Code = http.StatusOK
	response.Message = "OK"
	response.Data = data

	c.JSON(http.StatusOK, response)

	// } else {
	// 	response.Code = http.StatusBadRequest
	// 	response.Message = "BAD"
	// 	response.Data = nil

	// 	c.JSON(http.StatusBadRequest, response)
	// }
}
