package outbox

import (
	appRepo "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/relay/interfaces/persistence/dao"
	relay "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/dao/outbox"
	repo "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/repo/outbox"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func BuildOutboxRepo(conn *gorm.DB) appRepo.OutboxRepo {
	return &outbox.RepoImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}
func BuildOutboxDAO(conn *gorm.DB) dao.OutboxDAO {
	return &relay.DAOImpl{BaseGormRepo: repo.BaseGormRepo{Session: conn}}
}

var Module = fx.Provide(
	BuildOutboxRepo,
	BuildOutboxDAO,
)
