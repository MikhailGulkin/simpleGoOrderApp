package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/filters"
	q "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/presentation/api/controllers/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	mediator mediator.Mediator
}

func (c *Handler) CreateProduct(context *gin.Context) {
	var requestBody command.CreateProductCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := c.mediator.Send(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
func (c *Handler) UpdateProductName(context *gin.Context) {
	productID := context.Param("productID")

	var requestBody command.UpdateProductNameCommand
	if err := context.BindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	parse, err := uuid.Parse(productID)
	if err != nil {
		context.Error(err)
	}
	requestBody.ProductID = parse
	err = c.mediator.Send(requestBody)
	if err != nil {
		context.Error(err)
		return
	}
	context.Status(http.StatusNoContent)
}
func (c *Handler) GetAllProducts(context *gin.Context) {
	Limit, Offset, Order := handlers.GetQueryParams(context)
	products, err := c.mediator.Query(
		query.GetAllProductsQuery{
			BaseListQueryParams: q.BaseListQueryParams{
				Limit:  uint(Limit),
				Offset: uint(Offset),
				Order:  filters.BaseOrder(Order),
			},
		},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, products)
}
func (c *Handler) GetProductByName(context *gin.Context) {
	productName := context.Param("productName")
	productByName, err := c.mediator.Query(
		query.GetProductByNameQuery{Name: productName},
	)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, productByName)
}
