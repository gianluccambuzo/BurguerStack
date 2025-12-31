package main

import (
	"example/BurgerStack/config/db"
	"example/BurgerStack/controller"
	"example/BurgerStack/repository"
	"example/BurgerStack/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	OrderRepository := repository.NewOrderRepository(dbConnection)
	OrderUseCase := usecase.NewOrderUseCase(OrderRepository)
	OrderController := controller.NewOrderController(OrderUseCase)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	server.GET("/orders", OrderController.GetOrderList)
	server.POST("/orders", OrderController.CreateOrder)
	server.GET("/orders/:id", OrderController.GetOrderById)

	server.Run(":8080")
}
