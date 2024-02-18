package db

import (
	"database/sql"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (d *ProductDB) Get(id string) (application.ProductInterface, error) {

	row := d.db.QueryRow("SELECT id, name, price, status FROM products WHERE id = ?", id)
	var p application.Product

	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Status)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (d *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {

	if d.checkIfExists(product.GetID()) {

		err := d.update(product)

		if err != nil {

			return nil, err
		}
	} else {

		err := d.create(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (c *ProductDB) checkIfExists(id string) bool {

	err := c.db.QueryRow("SELECT id FROM products WHERE id = ?", id).Scan()

	return err != sql.ErrNoRows
}

func (d *ProductDB) update(p application.ProductInterface) error {

	_, err := d.db.Exec("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		p.GetName(),
		p.GetPrice(),
		p.GetStatus(),
		p.GetID())

	return err
}

func (d *ProductDB) create(p application.ProductInterface) error {

	_, err := d.db.Exec("INSERT INTO products(id, name, price, status) VALUES(?,?,?,?)",
		p.GetID(),
		p.GetName(),
		p.GetPrice(),
		p.GetStatus())

	return err
}
