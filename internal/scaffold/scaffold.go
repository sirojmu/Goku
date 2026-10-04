package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var projectName = regexp.MustCompile(`^[a-z][a-z0-9-]{1,62}$`)

func Create(directory, name string) error {
	if !projectName.MatchString(name) {
		return fmt.Errorf("nama proyek harus huruf kecil, angka, atau tanda hubung")
	}
	projectDirectory := filepath.Join(directory, name)
	if _, err := os.Stat(projectDirectory); err == nil {
		return fmt.Errorf("folder %q sudah ada", projectDirectory)
	}
	if err := os.MkdirAll(filepath.Join(projectDirectory, "templates"), 0o755); err != nil {
		return err
	}

	files := map[string]string{
		"go.mod": "module " + name + "\n\ngo 1.22\n\nrequire github.com/sirojmu/Goku v0.0.0\n",
		"main.go": `package main

import (
	"log"
	"net/http"

	"github.com/sirojmu/Goku"
)

func main() {
	app := goku.New(goku.WithViews("templates"))
	app.Use(goku.Recover)
	app.Get("/", func(context *goku.Context) error {
		return context.Render(http.StatusOK, "home", map[string]string{"Judul": "Halo dari Goku"})
	})
	app.Get("/api/salam/:nama", func(context *goku.Context) error {
		return context.JSON(http.StatusOK, map[string]string{"pesan": "Halo, " + context.Param("nama")})
	})

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
`,
		"templates/home.html": `<!doctype html>
<html lang="id">
  <head><meta charset="utf-8"><title>{{ .Judul }}</title></head>
  <body><h1>{{ .Judul }}</h1><p>Ubah <code>main.go</code> untuk mulai membuat aplikasi.</p></body>
</html>
`,
	}
	for path, content := range files {
		if err := os.WriteFile(filepath.Join(projectDirectory, path), []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}
