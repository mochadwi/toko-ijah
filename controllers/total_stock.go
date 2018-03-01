package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	models "github.com/mochadwi/toko-ijah/models"
	index "github.com/mochadwi/toko-ijah/views"
	uuid "github.com/satori/go.uuid"
)

// AddTotalStock will add new total stocks
func AddTotalStock(c *gin.Context) {
	// "name": "Zalekia Plain Casual Blouse",
	// "size": "L",
	// "colour": "Broken White",
	// "currentStock": 100

	currStock, _ := strconv.ParseInt(c.PostForm("currentStock"), 0, 64)
	var size = ""
	colour := strings.ToUpper(c.PostForm("colour")[0:3])
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	skuGenerated := strconv.Itoa(r1.Intn(99999999))

	switch c.PostForm("size") {
	case "XXXL":
		size = "X3"
	case "XXL":
		size = "XX"
	case "XL":
		size = "XL"
	default:
		size = c.PostForm("size") + c.PostForm("size")
	}

	if strings.Contains(c.PostForm("colour"), " ") {
		slice := strings.Split(c.PostForm("colour"), " ")
		colour = strings.ToUpper(slice[0][0:1] + slice[1][0:2])
	}

	data := models.TotalStockItem{
		SKU:          "SSI-D" + skuGenerated + "-" + size + "-" + colour,
		Name:         c.PostForm("name") + " (" + c.PostForm("size") + ", " + c.PostForm("colour") + ")",
		CurrentStock: currStock}

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
