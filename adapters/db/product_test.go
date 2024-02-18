package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/adapters/db"
	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	_, err := DB.Exec("CREATE TABLE products (id string, name string, price float, status string, PRIMARY KEY (id))")

	if err != nil {
		log.Fatal(err.Error())
	}

	createProduct(DB)

}

func createProduct(db *sql.DB) {

	_, err := db.Exec("INSERT INTO products(id, name, price, status) VALUES(?,?,?,?)", "1", "Product 1", 0, "disabled")

	if err != nil {
		log.Fatal(err.Error())
	}

}

func TestShouldGetAProduct(t *testing.T) {

	setup()

	defer DB.Close()

	db := db.NewProductDB(DB)

	p, err := db.Get("1")

	assert.Nil(t, err)
	assert.Equal(t, "Product 1", p.GetName())
	assert.Equal(t, 0.0, p.GetPrice())
	assert.Equal(t, "disabled", p.GetStatus())
}

func TestShouldCreateNewProductWhenProcutNotExists(t *testing.T) {

	setup()

	defer DB.Close()

	db := db.NewProductDB(DB)

	p := application.NewProduct("Novo Produto", 10.0)

	p.Enable()

	up, err := db.Save(p)

	assert.Nil(t, err)
	assert.Equal(t, p, up)

}

func TestShouldUpdateProductWhenProductExists(t *testing.T) {

	setup()

	defer DB.Close()

	db := db.NewProductDB(DB)

	p := application.NewProduct("Novo Produto", 10.0)

	p.Enable()

	db.Save(p)

	p.Name = "Atualualizando"

	up, err := db.Save(p)

	assert.Nil(t, err)
	assert.Equal(t, up, p)
}
