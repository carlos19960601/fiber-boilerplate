package http

type Option func(options *Server)

func WithHost(host string) Option {
	return func(options *Server) {
		options.host = host
	}
}

func WithPort(port int) Option {
	return func(options *Server) {
		options.port = port
	}
}
