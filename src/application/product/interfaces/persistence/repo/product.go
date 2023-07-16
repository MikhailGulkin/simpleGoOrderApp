package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/aggregate"
)

type ProductRepo interface {
	AcquireProductByID(productID id.ID) (aggregate.Product, error)
	AddProduct(product aggregate.Product, tx interface{}) error
	UpdateProduct(product aggregate.Product, tx interface{}) error
}
