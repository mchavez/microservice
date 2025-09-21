package middleware

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func UnaryLoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		latency := time.Since(start)

		fields := logrus.Fields{
			"method":  info.FullMethod,
			"latency": latency.String(),
		}

		if err != nil {
			logger.WithFields(fields).WithError(err).Error("gRPC request failed")
		} else {
			logger.WithFields(fields).Info("gRPC request completed")
		}

		return resp, err
	}
}
