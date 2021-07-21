package products

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/nelbermora/meli-bootcamp-storage/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	product := models.Product{
		Name: "test",
	}
	product2 := models.Product{
		Name: "test2",
	}
	myRepo := NewRepo()
	productResult, err := myRepo.Store(product)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, product.Name, productResult.Name)
	productResult2, err := myRepo.Store(product2)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, product2.Name, productResult2.Name)
}

func TestGetOne(t *testing.T) {
	id := 1
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
		ID:    1,
		Count: 3,
		Price: 120000,
	}
	myRepo := NewRepo()
	productResult, _ := myRepo.Update(product)
	assert.Equal(t, product, productResult)
}

func TestGetAll(t *testing.T) {
	myRepo := NewRepo()
	productsResult, _ := myRepo.GetAll()
	assert.LessOrEqual(t, 1, len(productsResult))
}

func TestDelete(t *testing.T) {
	myRepo := NewRepo()
	idForDelete := 2
	errResult := myRepo.Delete(idForDelete)
	assert.Nil(t, errResult)
}

func TestGetFullData(t *testing.T) {
	expectedWarehouse := "Main Warehouse"
	myRepo := NewRepo()
	idForSearch := 1
	productResult := myRepo.GetFullData(idForSearch)
	assert.Equal(t, productResult.Warehouse, expectedWarehouse)
}

func TestGetOneWithContext(t *testing.T) {
	// usamos un Id que exista en la DB
	id := 1
	// definimos un Product cuyo nombre sea igual al registro de la DB
	product := models.Product{
		Name: "test",
	}
	myRepo := NewRepo()
	// se define un context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	productResult, err := myRepo.GetOneWithcontext(ctx, id)
	fmt.Println(err)
	assert.Equal(t, product.Name, productResult.Name)
}
