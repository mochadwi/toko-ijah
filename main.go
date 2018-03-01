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
		v1.POST("/add/total/stock", controllers.AddTotalStock)
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