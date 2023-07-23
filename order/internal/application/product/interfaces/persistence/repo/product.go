package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/product/aggregate"
)

type ProductRepo interface {
	AcquireProductByID(productID id.ID) (aggregate.Product, error)
	AddProduct(product aggregate.Product, tx interface{}) error
	UpdateProduct(product aggregate.Product, tx interface{}) error
}
