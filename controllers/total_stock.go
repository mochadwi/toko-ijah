package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/mochadwi/toko-ijah/models"
	db "github.com/mochadwi/toko-ijah/utils"
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

	totalStock := models.TotalStockItem{
		SKU:          utils.GenerateSKU(c.PostForm("size"), c.PostForm("colour")),
		Name:         c.PostForm("name"),
		Size:         c.PostForm("size"),
		Colour:       c.PostForm("colour"),
		CurrentStock: utils.StrToInt(c.PostForm("currentStock"))}

	fmt.Println(totalStock)

	err := db.Mgr.AddTotalStock(&totalStock)

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
