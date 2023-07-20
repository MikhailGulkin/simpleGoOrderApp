package convertors

import (
	o "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/models"
	"reflect"
)

func OrdersEventsHandler(outboxes *[]models.Outbox, payload PayloadEnhanced) bool {
	if payload.reflect == reflect.TypeOf(new(o.OrderCreated)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange: "Orders",
			Route:    "Order.Create",
			Payload:  payload.payload,
		})
		return true
	}
	if payload.reflect == reflect.TypeOf(new(o.OrderAddProduct)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange: "Orders",
			Route:    "Order.AddProduct",
			Payload:  payload.payload,
		})
		return true
	}
	return false
}
