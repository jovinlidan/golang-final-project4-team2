package user_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-final-project3-team2/resources/user_resources"
	"golang-final-project3-team2/services/user_services"
	"golang-final-project3-team2/utils/error_utils"
	"golang-final-project3-team2/utils/success_utils"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userReq user_resources.UserRegisterRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserRegister(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UserLogin(c *gin.Context) {
	var userReq user_resources.UserLoginRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserLogin(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func TopupBalance(c *gin.Context) {
	var userReq user_resources.UserTopupBalanceRequest
	userIdToken := c.MustGet("user_id").(string)

	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	balance, err := user_services.UserService.UserTopup(userIdToken, &userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success(fmt.Sprintf("Your balance has been successfully updated to Rp %d", balance)))
}

func UpdateUser(c *gin.Context) {
	var userReq user_resources.UserUpdateRequest
	userIdToken := c.MustGet("user_id").(string)

	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	user, err := user_services.UserService.UserUpdate(userIdToken, &userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userIdToken := c.MustGet("user_id")

	err := user_services.UserService.UserDelete(userIdToken.(string))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, success_utils.Success("Your account has been successfully deleted"))
}
