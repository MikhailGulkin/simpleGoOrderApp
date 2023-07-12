package address

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/address/interfaces/command"
)

func NewProductHandler(
	createAddress c.CreateAddress,
) Handler {
	return Handler{
		createAddress: createAddress,
	}
}
