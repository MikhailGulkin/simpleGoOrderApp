package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"go.uber.org/fx"
)

func NewCreateProduct(repo repo.ProductRepo, uow persistence.UoW) c.CreateProduct {
	return &command.CreateProductImpl{
		ProductRepo: repo,
		UoW:         uow,
	}
}
func NewUpdateProductName(dao repo.ProductRepo, uow persistence.UoW) c.UpdateProductName {
	return &command.UpdateProductNameImpl{
		ProductRepo: dao,
		UoW:         uow,
	}
}

func NewGetALlProducts(dao reader.ProductReader, logger logger.Logger) q.GetAllProducts {
	return &query.GetAllProductsImpl{
		DAO:    dao,
		Logger: logger,
	}
}
func NewGetProductByName(dao reader.ProductReader) q.GetProductByName {
	return &query.GetProductByNameImpl{
		DAO: dao,
	}
}

var Module = fx.Provide(
	NewCreateProduct,
	NewGetALlProducts,
	NewUpdateProductName,
	NewGetProductByName,
)
