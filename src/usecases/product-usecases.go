package usecases

import (
	"log"

	"github.com/EdgarAllanPoo/test-go-api/src/domain"
)

type ProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewProductInteractor(repository domain.ProductRepository) ProductInteractor {
	return ProductInteractor{repository}
}

func (interactor *ProductInteractor) CreateProduct(product domain.Product) error {
	err := interactor.ProductRepository.SaveProduct(product)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *ProductInteractor) FindAll(category string, limit, offset int) ([]*domain.Product, int64, error) {
	results, totalRows, err := interactor.ProductRepository.FindAll(category, limit, offset)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, err
	}

	if results == nil {
		results = []*domain.Product{}
	}

	return results, totalRows, nil
}

func (interactor *ProductInteractor) FindById(id int64) (*domain.Product, error) {
	result, err := interactor.ProductRepository.FindById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func (interactor *ProductInteractor) DeleteProduct(id int64) error {
	err := interactor.ProductRepository.DeleteProduct(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *ProductInteractor) UpdateProduct(id int64, product domain.Product) error {
	err := interactor.ProductRepository.UpdateProduct(id, product)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
