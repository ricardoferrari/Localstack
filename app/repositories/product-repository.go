package productRepository

import (
	productModels "github.com/ricardoferrari/gorest/models"
)

type ProductRepositoryInterface interface {
	AddProduct(p productModels.Product)
	GetProduct(id int) productModels.Product
	GetProducts() []productModels.Product
	RemoveProduct(id int)
}

type ProductRepository struct {
	products []productModels.Product
}

func (pl *ProductRepository) AddProduct(p productModels.Product) {
	pl.products = append(pl.products, p)
}

func (pl *ProductRepository) GetProduct(id int) productModels.Product {
	for _, p := range pl.products {
		if p.Id == id {
			return p
		}
	}
	return productModels.Product{}
}

func (pl *ProductRepository) GetProducts() []productModels.Product {
	return pl.products
}

func (pl *ProductRepository) RemoveProduct(id int) {
	for i, p := range pl.products {
		if p.Id == id {
			pl.products = append(pl.products[:i], pl.products[i+1:]...)
			break
		}
	}
}

func NewProductRepository() ProductRepositoryInterface {
	return &ProductRepository{}
}
