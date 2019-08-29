package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/axiomzen/cors/v2"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(c.HandlerFunc)

	m.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Run()
}
