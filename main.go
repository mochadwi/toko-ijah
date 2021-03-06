package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	controllers "github.com/mochadwi/toko-ijah/controllers"
	index "github.com/mochadwi/toko-ijah/views"
	uuid "github.com/satori/go.uuid"
)

func main() {
	var baseURL = "/api/v1"
	engine := gin.Default()
	engine.RedirectTrailingSlash = false
	// engine.Use(utils.Middleware)

	v1 := engine.Group(baseURL)
	{
		// Adds total stock
		v1.POST("/total/stock", controllers.AddTotalStock)
		v1.GET("/total/stock/:id", controllers.GetTotalStock)
		v1.GET("/total/stock", controllers.GetAllTotalStock)
		v1.PUT("/total/stock/:id", controllers.UpdateTotalStockByID)
		v1.DELETE("/total/stock/:id", controllers.DeleteTotalStockByID)

		// Adds income stock
		v1.POST("/income/stock", controllers.AddIncomeStock)
		v1.GET("/income/stock/:id", controllers.GetIncomeStock)
		v1.GET("/income/stock", controllers.GetAllIncomeStock)
		v1.PUT("/income/stock/:id", controllers.UpdateIncomeStockByID)
		v1.DELETE("/income/stock/:id", controllers.DeleteIncomeStockByID)

		// Adds outcome stock
		v1.POST("/outcome/stock", controllers.AddOutcomeStock)
		v1.GET("/outcome/stock/:id", controllers.GetOutcomeStock)
		v1.GET("/outcome/stock", controllers.GetAllOutcomeStock)
		v1.PUT("/outcome/stock/:id", controllers.UpdateOutcomeStockByID)
		v1.DELETE("/outcome/stock/:id", controllers.DeleteOutcomeStockByID)

		// Generate value report
		v1.GET("/value/report", controllers.GenerateValueReport)
		v1.GET("/value/csv", controllers.GenerateValueCSV)

		// Generate sale report
		v1.GET("/sale/report", controllers.GenerateSaleReport)
		v1.GET("/sale/csv", controllers.GenerateSaleCSV)
	}

	engine.NoRoute(func(c *gin.Context) {
		var response = &index.DefaultResponseFormat{
			RequestID: uuid.NewV4().String(),
			Now:       time.Now().Unix(),
		}

		response.Code = http.StatusNotFound
		response.Message = "Service resource not found"

		c.JSON(http.StatusNotFound, response)
	})

	engine.Run(":7575")
}
