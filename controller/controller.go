package controller

import (
	"GoEBiznes/config"
	"GoEBiznes/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateProduct(c echo.Context) error {
	b := new(model.Product)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Create(&b).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	b := new(model.Product)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_product := new(model.Product)

	if err := db.First(&existing_product, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_product.Name = b.Name
	existing_product.Price = b.Price
	if err := db.Save(&existing_product).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_product,
	}

	return c.JSON(http.StatusOK, response)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var products []*model.Product

	if res := db.Find(&products, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": products[0],
	}

	return c.JSON(http.StatusOK, response)
}
func GetProducts(c echo.Context) error {
	db := config.DB()

	var products []*model.Product

	if res := db.Find(&products); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": products,
	}

	return c.JSON(http.StatusOK, response)
}
func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	product := new(model.Product)

	err := db.Delete(&product, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a product has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
