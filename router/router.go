package router

import (
	"github.com/gin-gonic/gin"
	"golang-final-project3-team2/controllers/category_controllers"
	"golang-final-project3-team2/controllers/user_controllers"
	"golang-final-project3-team2/db"
	"golang-final-project3-team2/middlewares"
	"log"
)

const PORT = ":8080"

func init() {
	db.InitializeDB()

}

func StartRouter() {
	router := gin.Default()
	apiRouter := router.Group("/api")
	{
		userRouter := apiRouter.Group("/users")
		{
			userRouter.POST("/register", user_controllers.CreateUser)
			userRouter.POST("/login", user_controllers.UserLogin)
			userRouter.Use(middlewares.MiddlewareAuth())
			userRouter.PATCH("/topup", user_controllers.TopupBalance)
			//userRouter.PUT("/update-account", user_controllers.UpdateUser)
			//userRouter.DELETE("/delete-account", user_controllers.DeleteUser)
		}

		categoryRouter := apiRouter.Group("/categories")
		{
			categoryRouter.Use(middlewares.MiddlewareAuth())
			categoryRouter.GET("/", category_controllers.GetCategories)
			categoryRouter.Use(middlewares.MiddlewareOnlyAdmin())
			categoryRouter.POST("/", category_controllers.CreateCategory)
			categoryRouter.PATCH("/:categoryId", category_controllers.UpdateCategory)
			categoryRouter.DELETE("/:categoryId", category_controllers.DeleteCategory)
		}
	}

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
