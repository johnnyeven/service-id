package global

import (
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
	NodeBits uint8 `conf:"env"`
	StepBits uint8 `conf:"env"`
}

var Config = struct {
	Log               *log.Log
	Server            transport_http.ServeHTTP
	GenerateAlgorithm types.GenerateAlgorithm `conf:"env"`
	SnowFlakeConfig   SnowFlakeConfig
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	Server: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
	GenerateAlgorithm: types.GENERATE_ALGORITHM__SNOWFLAKE,
	SnowFlakeConfig: SnowFlakeConfig{
		Epoch:    1288834974657,
		NodeBits: 10,
		StepBits: 12,
	},
}
