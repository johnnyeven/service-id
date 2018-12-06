package routes

import (
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/swagger"
)

var RootRouter = courier.NewRouter(GroupRoot{})
var V0Router = courier.NewRouter(GroupV0{})

func init() {
	RootRouter.Register(swagger.SwaggerRouter)
	RootRouter.Register(V0Router)
	V0Router.Register(courier.NewRouter(GetNewId{}))
}

type GroupRoot struct {
	courier.EmptyOperator
}

func (root GroupRoot) Path() string {
	return "/id"
}

type GroupV0 struct {
	courier.EmptyOperator
}

func (v0 GroupV0) Path() string {
	return "/v0"
}