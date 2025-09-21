package middleware

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// UnaryLoggingInterceptor logs the details of gRPC requests and their latency.
func UnaryLoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		latency := time.Since(start)

		logFields := logrus.Fields{
			"method":  info.FullMethod,
			"latency": latency.String(),
		}

		if err != nil {
			logger.WithFields(logFields).WithError(err).Error("gRPC request failed")
		} else {
			logger.WithFields(logFields).Info("gRPC request completed")
		}

		return resp, err
	}
}
