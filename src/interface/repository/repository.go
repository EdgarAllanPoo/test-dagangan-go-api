package repository

import "github.com/EdgarAllanPoo/test-go-api/src/domain"

type DBHandler interface {
	FindAllProducts(category string, limit, offset int) ([]*domain.Product, int64, error)
	FindProductById(id int64) (*domain.Product, error)
	SaveProduct(product domain.Product) error
	DeleteProduct(id int64) error
	UpdateProduct(id int64, product domain.Product) error
}
