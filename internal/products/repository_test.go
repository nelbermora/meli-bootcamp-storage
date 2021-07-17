package products

import (
	"log"
	"testing"

	"github.com/nelbermora/meli-bootcamp-storage/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	product := models.Product{
		Name: "test",
	}
	myRepo := NewRepo()
	productResult, err := myRepo.Store(product)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, product.Name, productResult.Name)
}

func TestGetOne(t *testing.T) {
	id := 4
	product := models.Product{
		Name: "test",
	}
	myRepo := NewRepo()
	productResult := myRepo.GetOne(id)
	assert.Equal(t, product.Name, productResult.Name)

}

func TestUpdate(t *testing.T) {
	product := models.Product{
		Name:  "Notebook",
		Type:  "Tech",
		ID:    4,
		Count: 3,
		Price: 120000,
	}
	myRepo := NewRepo()
	productResult, _ := myRepo.Update(product)
	assert.Equal(t, product, productResult)
}
