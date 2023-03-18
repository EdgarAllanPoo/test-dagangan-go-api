package repository

import "github.com/EdgarAllanPoo/test-go-api/src/domain"

type DBHandler interface {
	FindAllProducts() ([]*domain.Product, error)
	FindProductById(id int64) (*domain.Product, error)
	SaveProduct(product domain.Product) error
}
