package main

import (
	"database/sql"
	"log"

	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/adapters/db"
	"github.com/IcaroSilvaFK/fc2-arquitetura-hexagonal/application"
)

func main() {

	DB, _ := sql.Open("sqlite3", "db.sqlite3")

	pdbAdapter := db.NewProductDB(DB)

	pService := application.NewProductService(pdbAdapter)

	p, err := pService.Create("Product", 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	pService.Enable(p)

}
