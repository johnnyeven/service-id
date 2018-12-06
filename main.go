package main

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-id/global"
	"github.com/johnnyeven/service-id/routes"
)

func main() {
	servicex.Execute()
	global.Config.Server.Serve(routes.RootRouter)
}
