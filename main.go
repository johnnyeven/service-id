package main

import (
	"profzone/libtools/servicex"
	"profzone/service-id/global"
	"profzone/service-id/routes"
)

func main() {
	servicex.Execute()
	global.Config.Server.Serve(routes.RootRouter)
}
