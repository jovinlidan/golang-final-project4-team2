package product_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project4-team2/resources/product_resources"
	"golang-final-project4-team2/services/product_services"
	"golang-final-project4-team2/utils/error_utils"
	"golang-final-project4-team2/utils/success_utils"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var req product_resources.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	product, err := product_services.ProductService.CreateProduct(&req)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	products, err := product_services.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	var req product_resources.ProductUpdateRequest

	idParam := c.Param("productId")

	if err := c.ShouldBindJSON(&req); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}
	user, err := product_services.ProductService.UpdateProduct(&req, idParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("productId")

	err := product_services.ProductService.DeleteProduct(idParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Product has been successfully deleted"))
}
