package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	s := martini.Classic()
	s.Use(martini.Static("./root"))
	s.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("%s/%s", "./root", "layout.tmpl"))
	})
	s.RunOnAddr(":8080")
}
