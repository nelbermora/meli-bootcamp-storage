package products

import (
	"database/sql"
	"log"

	"github.com/nelbermora/meli-bootcamp-storage/db"
	"github.com/nelbermora/meli-bootcamp-storage/pkg/models"
)

type Repository interface {
	Store(models.Product) (models.Product, error)
	GetOne(id int) models.Product
	Update(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
}

const (
	InsertProduct = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	GetProduct    = "SELECT * FROM products WHERE id = ?"
	UpdateProduct = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
)

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	db := db.StorageDB                     // se inicializa la base
	stmt, err := db.Prepare(InsertProduct) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetOne(id int) models.Product {
	var product models.Product
	db := db.StorageDB
	rows, err := db.Query(GetProduct, id)
	if err != nil {
		log.Println(err)
		return product
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return product
		}
	}
	return product
}

func (r *repository) Update(product models.Product) (models.Product, error) {
	db := db.StorageDB                     // se inicializa la base
	stmt, err := db.Prepare(UpdateProduct) // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()                                                                       // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID) // retorna un sql.Result y un error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil

}

func (r *repository) GetAll() ([]models.Product, error) {
	return []models.Product{}, nil
}

func (r *repository) Delete(id int) error {
	return nil
}
