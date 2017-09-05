package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/secure"
	"net/http"
	"fmt"
)

const (
	sslCRT = "./server.crt"
	sslKEY = "./server.key"
)

func main() {
	go serverOne(":8080")
	serverTwo(":8081")
}

func serverOne(addr string) {
	m := martini.Classic()
	m.Use(secure.Secure(secure.Options{
		SSLRedirect: true,
		SSLHost:     addr,
	}))
	m.Get("/one", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("one"))
	})
	fmt.Println("one start ...")
	if err := http.ListenAndServeTLS(addr, sslCRT, sslKEY, m); err != nil {
		panic(err)
	}
}

func serverTwo(addr string) {
	m := martini.Classic()
	m.Use(secure.Secure(secure.Options{
		SSLRedirect: true,
		SSLHost:     addr,
	}))
	m.Get("/two", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("two"))
	})
	fmt.Println("two start ...")
	if err := http.ListenAndServeTLS(addr, sslCRT, sslKEY, m); err != nil {
		panic(err)
	}
}
