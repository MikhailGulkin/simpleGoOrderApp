package command

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence"
	outbox "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/repo"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/services"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/vo"
	"reflect"
)

type CreateOrderImpl struct {
	command.CreateOrder
	repo.OrderRepo
	services.Service
	persistence.UoW
	outbox.OutboxRepo
	dao.OrderDAO
}

func (interactor *CreateOrderImpl) Create(command command.CreateOrderCommand) error {
	previousOrder, previousOrderError := interactor.OrderRepo.AcquireLastOrder()
	if previousOrderError != nil {
		return previousOrderError
	}
	serialNumber := 0
	if !reflect.ValueOf(previousOrder).IsZero() {
		serialNumber = previousOrder.GetSerialNumber()
	}

	products, productError := interactor.OrderDAO.GetProductsByIDs(command.ProductsIDs)
	if productError != nil {
		return productError
	}
	orderAggregate, err := interactor.Service.CreateOrder(
		vo.OrderID{Value: command.OrderID},
		command.DeliveryAddress,
		command.UserID,
		serialNumber,
		products,
	)
	if err != nil {
		return err
	}
	err = interactor.OrderRepo.AddOrder(orderAggregate, interactor.UoW.StartTx())
	if err != nil {
		interactor.UoW.Rollback()
		return err
	}
	err = interactor.OutboxRepo.AddEvents(orderAggregate.PullEvents(), interactor.UoW.GetTx())
	if err != nil {
		interactor.UoW.Rollback()
		return err
	}
	err = interactor.UoW.Commit()
	if err != nil {
		return err
	}
	return nil
}
