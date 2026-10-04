// Package goku menyediakan fondasi aplikasi web Go yang kecil dan tanpa dependency eksternal.
package goku

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Handler func(*Context) error
type Middleware func(Handler) Handler
type Option func(*App)

type App struct {
	routes      []route
	middlewares []Middleware
	views       *Renderer
}

type route struct {
	method  string
	pattern string
	handler Handler
}

func New(options ...Option) *App {
	app := &App{}
	for _, option := range options {
		option(app)
	}
	return app
}

func WithViews(directory string) Option {
	return func(app *App) { app.views = NewRenderer(directory) }
}

func (app *App) Use(middlewares ...Middleware) {
	app.middlewares = append(app.middlewares, middlewares...)
}

func (app *App) Get(pattern string, handler Handler) { app.Handle(http.MethodGet, pattern, handler) }
func (app *App) Post(pattern string, handler Handler) { app.Handle(http.MethodPost, pattern, handler) }

func (app *App) Handle(method, pattern string, handler Handler) {
	if !strings.HasPrefix(pattern, "/") {
		panic("pola rute harus diawali '/'")
	}
	app.routes = append(app.routes, route{method: method, pattern: normalizePath(pattern), handler: handler})
}

func (app *App) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, registeredRoute := range app.routes {
		params, matched := match(registeredRoute.pattern, request.URL.Path)
		if !matched || registeredRoute.method != request.Method {
			continue
		}

		handler := registeredRoute.handler
		for index := len(app.middlewares) - 1; index >= 0; index-- {
			handler = app.middlewares[index](handler)
		}
		context := &Context{Request: request, Response: writer, Params: params, renderer: app.views}
		if err := handler(context); err != nil {
			http.Error(writer, "Terjadi kesalahan pada aplikasi.", http.StatusInternalServerError)
			log.Printf("goku: %v", err)
		}
		return
	}

	http.NotFound(writer, request)
}

func normalizePath(path string) string {
	if path == "/" {
		return path
	}
	return strings.TrimSuffix(path, "/")
}

func match(pattern, path string) (map[string]string, bool) {
	patternParts := splitPath(normalizePath(pattern))
	pathParts := splitPath(normalizePath(path))
	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(map[string]string)
	for index, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[strings.TrimPrefix(part, ":")] = pathParts[index]
			continue
		}
		if part != pathParts[index] {
			return nil, false
		}
	}
	return params, true
}

func splitPath(path string) []string {
	if path == "/" {
		return nil
	}
	return strings.Split(strings.Trim(path, "/"), "/")
}

func Recover(next Handler) Handler {
	return func(context *Context) (err error) {
		defer func() {
			if recovered := recover(); recovered != nil {
				err = fmt.Errorf("panic dipulihkan: %v", recovered)
			}
		}()
		return next(context)
	}
}
