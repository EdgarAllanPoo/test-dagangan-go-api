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

func (interactor *ProductInteractor) FindAll() ([]*domain.Product, error) {
	results, err := interactor.ProductRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
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

func (interactor *ProductInteractor) FilterByCategory(category string) ([]*domain.Product, error) {
	results, err := interactor.ProductRepository.FilterByCategory(category)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
