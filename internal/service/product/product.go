package product

import (
	"context"

	dto "github.com/mohammaderm/krad/internal/dto/product"
	"github.com/mohammaderm/krad/internal/repository/product"
	"github.com/mohammaderm/krad/log"
)

type (
	Service struct {
		logger            log.Logger
		productRepository product.ProductRepository
	}
	ProductServiceContract interface {
		GLTProduct(ctx context.Context) (dto.ProductRes, error)
		GetByID(ctx context.Context, req dto.FindProductReq) (dto.FindProductRes, error)
		GetByCategoryId(ctx context.Context, req dto.FindByCategoryIdReq) (dto.FindByCategoryIdRes, error)
	}
)

func NewService(logger log.Logger, productrepository product.ProductRepository) ProductServiceContract {
	return &Service{
		logger:            logger,
		productRepository: productrepository,
	}
}

func (s *Service) GLTProduct(ctx context.Context) (dto.ProductRes, error) {
	products, err := s.productRepository.GLTProduct(ctx)
	if err != nil {
		return dto.ProductRes{}, err
	}
	return dto.ProductRes{Products: products}, nil
}

func (s *Service) GetByID(ctx context.Context, req dto.FindProductReq) (dto.FindProductRes, error) {
	product, err := s.productRepository.GetByID(ctx, req.Id)
	if err != nil {
		return dto.FindProductRes{}, err
	}
	return dto.FindProductRes{
		Product: product,
	}, nil
}

func (s *Service) GetByCategoryId(ctx context.Context, req dto.FindByCategoryIdReq) (dto.FindByCategoryIdRes, error) {
	products, err := s.productRepository.GetByCategory(ctx, req.Offset, req.Id, req.Filter, req.Order)
	if err != nil {
		return dto.FindByCategoryIdRes{}, err
	}
	return dto.FindByCategoryIdRes{Products: products}, nil
}
