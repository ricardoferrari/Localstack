package productUseCase

import (
	productModels "github.com/ricardoferrari/gorest/models"
	productRepository "github.com/ricardoferrari/gorest/repositories"
)

type ProductListInterface interface {
	AddProduct(p productModels.Product)
	GetProduct(id int) productModels.Product
	GetProducts() []productModels.Product
	RemoveProduct(id int)
}

type ProductUseCase struct {
	products productRepository.ProductRepositoryInterface
}

func (productUseCase *ProductUseCase) AddProduct(p productModels.Product) {
	productUseCase.products.AddProduct(p)
}

func (productUseCase *ProductUseCase) GetProduct(id int) productModels.Product {
	return productUseCase.products.GetProduct(id)
}

func (productUseCase *ProductUseCase) GetProducts() []productModels.Product {
	return productUseCase.products.GetProducts()
}

func (productUseCase *ProductUseCase) RemoveProduct(id int) {
	productUseCase.products.RemoveProduct(id)
}

func NewProductUseCase(repository productRepository.ProductRepositoryInterface) ProductListInterface {
	return &ProductUseCase{products: repository}
}
