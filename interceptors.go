package go_grpc_logging_middleware

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Infof("Received request: %s", info.FullMethod)

		// Log the request body
		logger.Infof("Request body: %v", req)

		// Invoke the gRPC handler to process the request.
		resp, err := handler(ctx, req)

		if err != nil {
			logger.WithError(err).Errorf("Error handling request: %s", err.Error())
		} else {
			// Log the response body
			logger.Infof("Response body: %v", resp)
			logger.Infof("Sent response: %s", info.FullMethod)
		}

		return resp, err
	}
}
