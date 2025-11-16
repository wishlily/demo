package main

type Options struct {
	Port     int
	Protocol string
}

type Server struct {
	Options
}

type Option func(o *Options)

func WithPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func WithProtocol(protocol string) Option {
	return func(o *Options) {
		o.Protocol = protocol
	}
}

// changeable args
func NewServer(opts ...Option) *Server {
	o := &Options{
		Port:     8080,   // default
		Protocol: "http", // default
	}
	for _, opt := range opts {
		opt(o)
	}
	return &Server{*o}
}

func main() {
	s := NewServer(
		WithPort(8081),
		WithProtocol("https"),
	)
	println(s.Port)
	println(s.Protocol)
}
