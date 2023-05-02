package main

import (
	"GoEBiznes/config"
	"GoEBiznes/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create HTTP server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	bookRoute := e.Group("/product")
	bookRoute.POST("/", controller.CreateProduct)
	bookRoute.GET("/:id", controller.GetProduct)
	bookRoute.PUT("/:id", controller.UpdateProduct)
	bookRoute.DELETE("/:id", controller.DeleteProduct)
	bookRoute.GET("/", controller.GetProducts)
	e.Logger.Fatal(e.Start(":8080"))
}
