package grpc_server

import (
	"context"
	"math/big"

	"github.com/gmx-delta-neutral/gmx-neutral.command/internal/glp"
	"github.com/gmx-delta-neutral/gmx-neutral.command/pkg/command/api"
)

func NewGlpServer(glpService glp.Service) *GlpServer {
	return &GlpServer{
		glpService: glpService,
	}
}

type GlpServer struct {
	glpService glp.Service
}

func (p *GlpServer) BuyGlp(ctx context.Context, request *api.BuyGlpRequest) (*api.BuyGlpResponse, error) {
	response := &api.BuyGlpResponse{}
	amount := new(big.Int).SetBytes(request.Amount)
	err := p.glpService.BuyGlp(amount)

	if err != nil {
		return response, err
	}

	return response, nil
}
