package routes

import (
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/engine"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/middleware"
)

func NewGroupRoutes(config c.APIConfig, handler engine.RequestHandler, middlewares middleware.Middlewares) engine.GroupRoutes {
	return engine.GroupRoutes{RouterGroup: handler.Gin.Group(config.BaseURLPrefix, middlewares...)}
}
