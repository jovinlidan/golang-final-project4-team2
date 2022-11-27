package category_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project3-team2/resources/category_resources"
	"golang-final-project3-team2/services/category_services"
	"golang-final-project3-team2/utils/error_utils"
	"golang-final-project3-team2/utils/success_utils"
	"net/http"
)

func CreateCategory(c *gin.Context) {
	var req category_resources.CategoryCreateRequest
	userIdToken := c.MustGet("user_id")
	if err := c.ShouldBindJSON(&req); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	category, err := category_services.CategoryService.CreateCategory(&req, userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, category)
}

func GetCategories(c *gin.Context) {
	user, err := category_services.CategoryService.GetCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateCategory(c *gin.Context) {
	var req category_resources.CategoryUpdateRequest

	idParam := c.Param("categoryId")

	if err := c.ShouldBindJSON(&req); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}
	user, err := category_services.CategoryService.UpdateCategory(&req, idParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteCategory(c *gin.Context) {
	idParam := c.Param("categoryId")

	err := category_services.CategoryService.DeleteCategory(idParam)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Category has been successfully deleted"))
}
