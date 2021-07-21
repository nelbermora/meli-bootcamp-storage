package products

import (
	"context"
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
	GetFullData(id int) models.Product
	GetOneWithcontext(ctx context.Context, id int) (models.Product, error)
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

const (
	GetAllProducts = "SELECT * FROM products"
)

func (r *repository) GetAll() ([]models.Product, error) {
	var products []models.Product
	db := db.StorageDB
	rows, err := db.Query(GetAllProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Delete(id int) error {
	db := db.StorageDB                                           // se inicializa la base
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()     // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(id) // retorna un sql.Result y un error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetFullData(id int) models.Product {
	var product models.Product
	db := db.StorageDB
	innerJoin := "SELECT products.id, products.name, products.type, products.count, products.price, warehouses.name, warehouses.adress " +
		"FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id " +
		"WHERE products.id = ?"
	rows, err := db.Query(innerJoin, id)
	if err != nil {
		log.Println(err)
		return product
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse, &product.WarehouseAdress); err != nil {
			log.Fatal(err)
			return product
		}
	}
	return product
}

func (r *repository) GetOneWithcontext(ctx context.Context, id int) (models.Product, error) {
	var product models.Product
	db := db.StorageDB
	// se especifican los campos en la query
	//getQuery := "SELECT p.id, p.name, p.type, p.count, p.price FROM products p WHERE p.id = ?"
	// se utiliza una query que demore 30 segundos en ejecutarse
	getQuery := "SELECT SLEEP(30) FROM DUAL where 0 < ?"
	// ya no se usa db.Query sino db.QueryContext
	rows, err := db.QueryContext(ctx, getQuery, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return product, err
		}
	}
	return product, nil
}
