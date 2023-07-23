package address

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
)

func NewProductHandler(
	mediator mediator.Mediator,
) Handler {
	return Handler{
		mediator: mediator,
	}
}
