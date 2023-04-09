package tracer

import (
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"route256/libs/logger"
)

func InitGlobalTracer(serviceName string) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	cfg, _ = config.FromEnv()

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("failed to initialize tracer", zap.Error(err))
	}
}
