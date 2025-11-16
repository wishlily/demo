package main

type Option struct {
	Port     int
	Protocol string
}

type Server struct {
	Option
}

type ServerBuilder struct {
	option Option
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.option.Port = port
	return b
}

func (b *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	b.option.Protocol = protocol
	return b
}

func (b *ServerBuilder) Build() *Server {
	return &Server{b.option}
}

func main() {
	server := NewServerBuilder().
		WithPort(8080).
		WithProtocol("http").
		Build()
	println(server.Option.Port)
	println(server.Option.Protocol)
}
