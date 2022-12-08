package router

import (
	"github.com/gin-gonic/gin"
	"golang-final-project4-team2/controllers/category_controllers"
	"golang-final-project4-team2/controllers/product_controllers"
	"golang-final-project4-team2/controllers/user_controllers"
	"golang-final-project4-team2/db"
	"golang-final-project4-team2/middlewares"
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

		productRouter := apiRouter.Group("/products")
		{
			productRouter.Use(middlewares.MiddlewareAuth())
			productRouter.GET("/", product_controllers.GetProducts)
			productRouter.Use(middlewares.MiddlewareOnlyAdmin())
			productRouter.POST("/", product_controllers.CreateProduct)
			productRouter.PUT("/:productId", product_controllers.UpdateProduct)
			productRouter.DELETE("/:productId", product_controllers.DeleteProduct)

		}
	}

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}
