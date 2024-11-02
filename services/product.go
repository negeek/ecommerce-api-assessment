package services

import (
	"errors"

	"github.com/negeek/ecommerce-api-assessment/repositories"
)

type ProductService struct{}

func (s *ProductService) Create(product *repositories.Product) error {
	return product.Create()
}

func (s *ProductService) FindByID(id int) (*repositories.Product, error) {
	product := &repositories.Product{}
	err := product.FindByID(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}
	return product, nil
}

func (s *ProductService) Update(product *repositories.Product, partial bool) error {
	if partial {
		return product.Patch()
	}
	return product.Put()
}

func (s *ProductService) Delete(id int) error {
	product := &repositories.Product{}
	return product.Delete(id)
}

func (ps *ProductService) List() ([]repositories.Product, error) {
	products, err := repositories.ProductList()
	if err != nil {
		return nil, err
	}
	return products, nil
}
