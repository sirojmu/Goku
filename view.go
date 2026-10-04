package goku

import (
	"html/template"
	"io"
	"path/filepath"
)

type Renderer struct{ directory string }

func NewRenderer(directory string) *Renderer { return &Renderer{directory: directory} }

func (renderer *Renderer) Render(writer io.Writer, name string, data any) error {
	path := filepath.Join(renderer.directory, name+".html")
	templateFile, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	return templateFile.ExecuteTemplate(writer, filepath.Base(path), data)
}
