go version >= 1.8

from https://github.com/vladimirvivien/go-plugin-example

	go build -buildmode=plugin -o eng/eng.so eng/greeter.go
	go build -buildmode=plugin -o chi/chi.so chi/greeter.go

	go run greeter.go english
	go run greeter.go chinese