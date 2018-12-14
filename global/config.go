package global

import (
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-id/constants/types"
	"github.com/johnnyeven/libtools/courier/transport_grpc"
)

func init() {
	servicex.SetServiceName("service-id")
	servicex.ConfP(&Config)

}

type SnowFlakeConfig struct {
	Epoch    int64 `conf:"env"`
	NodeBits uint8 `conf:"env"`
	StepBits uint8 `conf:"env"`
}

var Config = struct {
	Log               *log.Log
	ServerHTTP        transport_http.ServeHTTP
	ServerGRPC        transport_grpc.ServeGRPC
	GenerateAlgorithm types.GenerateAlgorithm `conf:"env"`
	SnowFlakeConfig   SnowFlakeConfig
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	ServerHTTP: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
	ServerGRPC: transport_grpc.ServeGRPC{
		Port: 9990,
	},
	GenerateAlgorithm: types.GENERATE_ALGORITHM__SNOWFLAKE,
	SnowFlakeConfig: SnowFlakeConfig{
		Epoch:    1288834974657,
		NodeBits: 10,
		StepBits: 12,
	},
}
