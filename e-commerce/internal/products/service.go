package products

import (
	"context"

	repo "github.com/agamyo168/e-commerce/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context)([]repo.Product, error)
	CreateProduct(ctx context.Context, productDto createProductParams )(repo.Product, error)
}

type svc struct {
	repo *repo.Queries
	}

func NewService(repo *repo.Queries) Service {
	return &svc{repo}
}
func (s *svc) ListProducts(ctx context.Context)([]repo.Product, error){
return s.repo.ListProducts(ctx)
}
func (s *svc) CreateProduct (ctx context.Context, productDto createProductParams)(repo.Product, error){
	return s.repo.CreateProduct(ctx, repo.CreateProductParams(productDto))
}