package main

import "gopkg.in/macaron.v1"

func main() {

	m := macaron.New()
	m.SetAutoHead(true)

	m.Get("/", func() {
		// show something
	})

	m.Patch("/", func() {
		// update something
	})

	m.Post("/", func() {
		// create something
	})

	m.Put("/", func() {
		// replace something
	})

	m.Delete("/", func() {
		// destroy something
	})

	m.Options("/", func() {
		// http options
	})

	m.Any("/", func() {
		// do anything
	})

	m.Route("/", "GET,POST", func() {
		// combine something
	})

	m.Combo("/").
		Get(func() string { return "GET" }).
		Patch(func() string { return "PATCH" }).
		Post(func() string { return "POST" }).
		Put(func() string { return "PUT" }).
		Delete(func() string { return "DELETE" }).
		Options(func() string { return "OPTIONS" }).
		Head(func() string { return "HEAD" })

	m.NotFound(func() {
		// Custom handle for 404
	})
}
