package znet

type Server struct {
	// server name
	Name string
	// tcp4 or other
	IPVersion string
	// bind listen ip addr
	IP string
	// bind listen port
	Port int
}

// start server
func (s *Server) Start() {

}

// stop server
func (s *Server) Stop() {

}

// run server
func (s *Server) Serve() {

}
