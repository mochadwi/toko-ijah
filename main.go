package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	index "github.com/mochadwi/toko-ijah/views"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// var baseURL = "/api/v1"
	engine := gin.Default()
	engine.RedirectTrailingSlash = false
	// engine.Use(utils.Middleware)

	engine.NoRoute(func(c *gin.Context) {
		var response = &index.DefaultResponseFormat{
			RequestID: uuid.NewV4().String(),
			Now:       time.Now().Unix(),
		}

		response.Code = strconv.Itoa(http.StatusNotFound)
		response.Message = "Service resource not found"

		c.JSON(http.StatusNotFound, response)
	})

	engine.Run(":7575")
}
