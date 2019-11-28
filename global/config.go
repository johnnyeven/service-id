package global

import (
	"github.com/johnnyeven/libtools/courier/transport_grpc"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-id/constants/types"
)

func init() {
	servicex.SetServiceName("service-id")
	servicex.ConfP(&Config)

}

type SnowFlakeConfig struct {
	Epoch    int64 `conf:"env"`
	NodeID   int64 `conf:"env"`
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
		Port:     8001,
	},
	ServerGRPC: transport_grpc.ServeGRPC{
		Port: 9991,
	},
	GenerateAlgorithm: types.GENERATE_ALGORITHM__SNOWFLAKE,
	SnowFlakeConfig: SnowFlakeConfig{
		Epoch:    1288834974657,
		NodeID:   1,
		NodeBits: 10,
		StepBits: 12,
	},
}
