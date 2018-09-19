package main

import (
	"github.com/profzone/libtools/servicex"
	"github.com/profzone/service-id/global"
	"github.com/profzone/service-id/routes"
)

func main() {
	servicex.Execute()
	global.Config.Server.Serve(routes.RootRouter)
}
