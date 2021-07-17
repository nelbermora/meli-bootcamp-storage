package employees

import (
	"github.com/nelbermora/meli-bootcamp-storage/pkg/models"
)

type Repository interface {
	Store(name, productType string, count int, price float64) (models.Product, error)
	GetOne(id int)
	Update(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
}

type repository struct {
}
