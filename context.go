package goku

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Params   map[string]string
	renderer *Renderer
}

func (context *Context) Param(name string) string { return context.Params[name] }
func (context *Context) Query(name string) string { return context.Request.URL.Query().Get(name) }

func (context *Context) String(status int, value string) error {
	context.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	context.Response.WriteHeader(status)
	_, err := context.Response.Write([]byte(value))
	return err
}

func (context *Context) JSON(status int, value any) error {
	context.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
	context.Response.WriteHeader(status)
	return json.NewEncoder(context.Response).Encode(value)
}

func (context *Context) Render(status int, name string, data any) error {
	if context.renderer == nil {
		return context.String(http.StatusInternalServerError, "Direktori template belum dikonfigurasi.")
	}
	context.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	context.Response.WriteHeader(status)
	return context.renderer.Render(context.Response, name, data)
}
