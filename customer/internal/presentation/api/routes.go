package api

import "github.com/gofiber/fiber/v2"

type Routes struct {
	FiberGroup
	controller Handler
}

type Route interface {
	Setup()
}

func (r Routes) Setup() {
	r.Add(fiber.MethodPost, "/upload-avatar/:id", r.controller.UploadNewAvatar)
}

func NewCustomerRouter(
	group FiberGroup,
	controller Handler,
) Route {
	return Routes{controller: controller, FiberGroup: group}
}
