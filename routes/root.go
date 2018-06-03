package routes

import (
	"profzone/libtools/courier"
	"profzone/libtools/courier/swagger"
)

var RootRouter = courier.NewRouter(GroupRoot{})

func init() {
	RootRouter.Register(swagger.SwaggerRouter)
}

type GroupRoot struct {
	courier.EmptyOperator
}

func (root GroupRoot) Path() string {
	return "/service-id"
}
