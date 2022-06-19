package infrastructure

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type interceptors struct {
	logger *log.Entry
}

func NewInterceptors(logger *log.Entry) *interceptors {
	return &interceptors{logger}
}

func (i *interceptors) LoggingServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		span := trace.SpanFromContext(ctx).SpanContext().SpanID()
		trace := trace.SpanFromContext(ctx).SpanContext().TraceID()
		fields := log.Fields{
			"method": info.FullMethod,
		}

		if span.IsValid() {
			fields["span_id"] = span.String()
		}

		if trace.IsValid() {
			fields["trace_id"] = trace.String()
		}

		i.logger.WithContext(ctx).
			WithFields(fields).Info("Request started")

		h, err := handler(ctx, req)

		i.logger.WithContext(ctx).
			WithFields(log.Fields{
				"method":   info.FullMethod,
				"duration": time.Since(start),
			}).Info("Request finished")

		return h, err
	}
}

func (i *interceptors) LoggingClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()

	i.logger.WithContext(ctx).
		WithFields(log.Fields{
			"method": method,
		}).Info("Request started")

	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	i.logger.WithContext(ctx).
		WithFields(log.Fields{
			"method":   method,
			"duration": time.Since(start),
		}).Info("Request finished")

	return err
}

func (i *interceptors) TracingClientInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		// Logic before invoking the invoker
		newCtx, span := otel.Tracer(serviceName).Start(ctx, "grpc")
		defer span.End()

		// Calls the invoker to execute RPC
		err := invoker(newCtx, method, req, reply, cc, opts...)

		return err
	}
}

func (i *interceptors) TracingServerInterceptor(serviceName string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		newCtx, span := otel.Tracer(serviceName).Start(ctx, "grpc")
		defer span.End()
		h, err := handler(newCtx, req)

		return h, err
	}
}
