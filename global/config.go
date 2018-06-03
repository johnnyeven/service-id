package global

import (
	"profzone/libtools/courier/transport_http"
	"profzone/libtools/log"
	"profzone/libtools/servicex"
)

func init() {
	servicex.SetServiceName("service-id")
	servicex.ConfP(&Config)

}

var Config = struct {
	Log    *log.Log
	Server transport_http.ServeHTTP
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	Server: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
}
