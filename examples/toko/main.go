package main

import (
	"log"
	"net/http"

	"github.com/sirojmu/Goku"
)

func main() {
	app := goku.New(goku.WithViews("examples/toko/templates"))
	app.Use(goku.Recover)
	app.Get("/", func(context *goku.Context) error {
		return context.Render(http.StatusOK, "home", map[string]string{"Nama": "Toko Contoh"})
	})
	app.Get("/api/produk/:id", func(context *goku.Context) error {
		return context.JSON(http.StatusOK, map[string]string{"id": context.Param("id")})
	})

	log.Println("Buka http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
