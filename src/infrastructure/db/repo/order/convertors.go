package order

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/consts"
	o "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
)

func ConvertOrderModelToAggregate(model models.Order) order.Order {
	products := make([]o.OrderProduct, len(model.Products))
	for index, product := range model.Products {
		products[index] = o.OrderProduct{
			ProductID: product.ID,
			Price:     product.Price,
		}
	}
	return order.Order{
		OrderID:  vo.OrderID{Value: model.ID},
		Products: products,
		Client: o.OrderClient{
			ClientID: model.Client.ID,
			Username: model.Client.Username,
		},
		OrderStatus:   consts.OrderStatus(model.OrderStatus),
		PaymentMethod: consts.PaymentMethod(model.PaymentMethod),
		DeliveryAddress: o.OrderAddress{
			AddressID:   model.Address.ID,
			FullAddress: model.Address.GetFullAddress(),
		},
		Date:         model.Date,
		SerialNumber: model.SerialNumber,
		Closed:       model.Closed,
	}
}
func ConvertOrderAggregateToModel(order order.Order) models.Order {
	return models.Order{
		Base:          models.Base{ID: order.OrderID.Value},
		OrderStatus:   string(order.OrderStatus),
		ClientID:      order.Client.ClientID,
		PaymentMethod: string(order.PaymentMethod),
		AddressID:     order.DeliveryAddress.AddressID,
		Date:          order.Date,
		Closed:        order.Closed,
		SerialNumber:  order.SerialNumber,
	}

}
