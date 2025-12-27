package controller

import (
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
