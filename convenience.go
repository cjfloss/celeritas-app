package main

import "net/http"

func (a *application) get(route string, h http.HandlerFunc) {
	a.App.Routes.Get(route, h)
}

func (a *application) post(route string, h http.HandlerFunc) {
	a.App.Routes.Post(route, h)
}

func (a *application) use(m ...func(http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
