package grpc_server

import (
	"context"
	"math/big"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/glp"
)

func NewGlpServer(glpService glp.Service) *GlpServer {
	return &GlpServer{
		glpService: glpService,
	}
}

type GlpServer struct {
	glpService glp.Service
}

func (p *GlpServer) BuyGlp(ctx context.Context, request *generated.BuyGlpRequest) (*generated.BuyGlpResponse, error) {
	response := &generated.BuyGlpResponse{}
	amount := new(big.Int).SetBytes(request.Amount)
	err := p.glpService.BuyGlp(amount)

	if err != nil {
		return response, err
	}

	return response, nil
}
