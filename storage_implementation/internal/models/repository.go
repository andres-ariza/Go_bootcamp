package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/andres-ariza/Go_bootcamp/storage_implementation/internal/db"
)

type Repository interface {
	Store(product Product) (Product, error)
	GetOne(id int) Product
	Update(product Product) (Product, error)
	GetByName(name string) Product // ejercicio 1
	GetAll() ([]Product, error)
	Delete(id int) error
	GetFullData(id int) Product
	GetOneWithcontext(ctx context.Context, id int) (Product, error)
}
type repository struct{}

func NewRepo() Repository {
	return &repository{}
}

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"nombre"`
	Type            string  `json:"tipo"`
	Count           int     `json:"cantidad"`
	Price           float64 `json:"precio"`
	WarehouseID     int     `json:"numero_deposito"`
	Warehouse       string  `json:"deposito"`
	WarehouseAdress string  `json:"direccion_deposito"`
}

const (
	GetProduct = "SELECT * FROM products WHERE id = ?"
	GetByName  = "SELECT * FROM products WHERE name = ?"
)

func (r *repository) Store(product Product) (Product, error) {
	db := db.StorageDB                                                                             // se inicializa la base
	stmt, err := db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara el SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) GetOne(id int) Product {
	var product Product
	db := db.StorageDB
	rows, err := db.Query(GetProduct, id)
	if err != nil {
		log.Println(err)
		return product
	}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			log.Fatal(err)
			return product
		}
	}
	return product
}

func (r *repository) Update(product Product) (Product, error) {
	db := db.StorageDB                                                                                   // se inicializa la base
	stmt, err := db.Prepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return Product{}, err
	}
	return product, nil

}

// ej 1 implementar geybyname()
func (r *repository) GetByName(name string) Product {
	var product Product
	db := db.StorageDB
	rows, err := db.Query(GetByName, name)
	if err != nil {
		log.Println(err)
		return product
	}
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.WarehouseID) // product.ID, &product.Name, &product.Type, &product.Count, &product.Price
		if err != nil {
			log.Fatal(err)
			return product
		}
	}
	return product
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product Product
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
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	_, err = stmt.Exec(id) // retorna un sql.Result y un error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetFullData(id int) Product {
	var product Product
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
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.Warehouse,
			&product.WarehouseAdress); err != nil {
			log.Fatal(err)
			return product
		}
	}
	return product
}

func (r *repository) GetOneWithcontext(ctx context.Context, id int) (Product, error) {
	var product Product
	db := db.StorageDB
	// // se especifican estrictamente los campos necesarios en la query
	// getQuery := "SELECT p.id, p.name, p.type, p.count, p.price FROM products p WHERE p.id = ?"
	// // ya no se usa db.Query sino db.QueryContext
	// rows, err := db.QueryContext(ctx, getQuery, id)
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
