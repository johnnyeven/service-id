package main

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-id/global"
	"github.com/johnnyeven/service-id/routes"
)

func main() {
	servicex.Execute()
	go global.Config.ServerGRPC.Serve(routes.RootRouter)
	global.Config.ServerHTTP.Serve(routes.RootRouter)
}
