package interceptors

import (
	"context"
	"fmt"
	"route256/libs/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info("request started", zap.String("method", info.FullMethod), zap.String("request", fmt.Sprintf("%v", req)))

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error("request failed", zap.String("method", info.FullMethod), zap.Error(err))
		return nil, err
	}

	logger.Info("request finished", zap.String("method", info.FullMethod), zap.String("response", fmt.Sprintf("%v", res)))

	return res, nil
}
