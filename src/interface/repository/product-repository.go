package repository

import (
	"github.com/EdgarAllanPoo/test-go-api/src/domain"
)

type ProductRepo struct {
	handler DBHandler
}

func NewProductRepo(handler DBHandler) ProductRepo {
	return ProductRepo{handler}
}

func (repo ProductRepo) SaveProduct(product domain.Product) error {
	err := repo.handler.SaveProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (repo ProductRepo) FindAll(category string, limit, offset int) ([]*domain.Product, int64, error) {
	results, totalRows, err := repo.handler.FindAllProducts(category, limit, offset)
	if err != nil {
		return results, 0, err
	}
	return results, totalRows, nil
}

func (repo ProductRepo) FindById(id int64) (*domain.Product, error) {
	result, err := repo.handler.FindProductById(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (repo ProductRepo) DeleteProduct(id int64) error {
	err := repo.handler.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

func (repo ProductRepo) UpdateProduct(id int64, product domain.Product) error {
	err := repo.handler.UpdateProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}
