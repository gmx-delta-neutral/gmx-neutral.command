package glp

import "math/big"

type Service interface {
	BuyGlp(amount big.Int) error
}

type GlpService struct {
	glpRepository Repository
}

func NewService(glpRepository Repository) *GlpService {
	return &GlpService{glpRepository: glpRepository}
}

func (s *GlpService) BuyGlp(amount big.Int) error {
	err := s.glpRepository.BuyGlp(amount)
	return err
}
