package main

import (
	"context"
	"log"
	"math/big"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/infrastructure"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const ServiceName = "gmx-neutral.command-client"

type asset struct {
	Symbol     string
	Weight     *big.Int
	PoolAmount *big.Int
}

func main() {
	logger := infrastructure.NewLogger(ServiceName)
	interceptors := infrastructure.NewInterceptors(logger)

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()), grpc.WithUnaryInterceptor(interceptors.TracingClientInterceptor(ServiceName)))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := generated.NewGlpServiceClient(conn)

	// create request
	req := generated.BuyGlpRequest{
		Amount: big.NewInt(1).Bytes(),
	}

	// call Greet service
	_, err = c.BuyGlp(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	if err != nil {
		panic(err)
	}
}
