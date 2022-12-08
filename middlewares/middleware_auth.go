package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang-final-project4-team2/domains/user_domain"
	"golang-final-project4-team2/resources/user_resources"
	"golang-final-project4-team2/utils/error_utils"
	"golang-final-project4-team2/utils/helpers"
	"strconv"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		unauthorizedErr := error_utils.NewUnauthorizedRequest(errors.New("please sign in to proceed").Error())

		data, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		userId := int64(data.(jwt.MapClaims)["id"].(float64))
		exists := user_domain.UserDomain.UserCheckIsExists(userId)
		if !exists {
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Set("user_id", strconv.FormatInt(userId, 10))
		c.Next()
	}
}

func MiddlewareOnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		unauthorizedErr := error_utils.NewUnauthorizedRequest(errors.New("please sign in to proceed").Error())
		wrongRoleErr := error_utils.NewUnauthorizedRequest(errors.New("your account doesn't have access to this data").Error())

		data, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}
		userRole := data.(jwt.MapClaims)["role"].(string)

		if userRole != string(user_resources.RoleAdmin) {
			c.AbortWithStatusJSON(wrongRoleErr.Status(), wrongRoleErr)
			return
		}

		c.Next()
	}
}
