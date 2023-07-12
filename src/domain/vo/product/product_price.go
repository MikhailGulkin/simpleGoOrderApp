package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/exceptions"
	"strconv"
)

type ProductPrice struct {
	Value float64
}

func (ProductPrice) Create(price float64) (ProductPrice, error) {
	if price < 0 {
		priceError := exceptions.InvalidPriceProductCreation{}.Exception(strconv.Itoa(int(price)))
		return ProductPrice{}, &priceError
	}
	return ProductPrice{Value: price}, nil
}
