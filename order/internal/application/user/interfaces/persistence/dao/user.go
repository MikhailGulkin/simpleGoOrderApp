package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/user/entities"
	"github.com/google/uuid"
)

type UserDAO interface {
	GetUserAddress(uuid uuid.UUID) (entities.UserAddress, error)
}
