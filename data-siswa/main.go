package main

import (
	"data-siswa/db"
	"data-siswa/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
