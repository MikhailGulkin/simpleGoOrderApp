package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	outboxRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/repo"
	q "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/query"
	s "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/saga"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/services"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"go.uber.org/fx"
)

func NewCreateOrder(
	uow persistence.UoW,
	orderRepo repo.OrderRepo,
	orderDAO dao.OrderDAO,
	outboxRepo outboxRepo.OutboxRepo,
) c.CreateOrder {
	return &command.CreateOrderImpl{
		UoW:        uow,
		OrderRepo:  orderRepo,
		Service:    services.Service{},
		OutboxRepo: outboxRepo,
		OrderDAO:   orderDAO,
	}
}
func NewGetAllOrders(dao reader.OrderReader) q.GetAllOrders {
	return &query.GetAllOrderImpl{
		OrderReader: dao,
	}
}
func NewGetOrderByID(dao reader.OrderReader) q.GetOrderByID {
	return &query.GetOrderByIDImpl{
		OrderReader: dao,
	}
}

func NewSagaCreateOrder(dao dao.OrderSagaDAO, uow persistence.UoW, logger logger.Logger) saga.CreateOrder {
	return &s.CreateOrderImpl{
		OrderSagaDAO: dao,
		UoW:          uow,
		Logger:       logger,
	}
}

var Module = fx.Provide(
	NewCreateOrder,
	NewGetAllOrders,
	NewGetOrderByID,

	NewSagaCreateOrder,
)
