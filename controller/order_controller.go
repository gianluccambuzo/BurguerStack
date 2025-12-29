package controller

import (
	"example/BurgerStack/model"
	"example/BurgerStack/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type orderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) orderController {
	return orderController{
		orderUseCase: orderUseCase,
	}
}

func (orderController *orderController) GetOrderList(ctx *gin.Context) {

	orders, err := orderController.orderUseCase.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, orders)
}

func (orderController *orderController) GetOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")

	if orderId == "" {
		response := model.Response{Message: "order id is empty"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := orderController.orderUseCase.GetOrderById(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if order == nil {
		response := model.Response{Message: "order not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, &order)

}

func (orderController *orderController) CreateOrder(ctx *gin.Context) {
	var order model.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response, err := orderController.orderUseCase.CreateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, response)

}
