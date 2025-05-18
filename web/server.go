package web

import (
	"net"
	"net/http"
)

var _Server = &HTTPServer{}

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	AddRoute(method, path string, handleFunc HandleFunc)
}

func (h *HTTPServer) AddRoute(method, path string, handleFunc HandleFunc) {

}

type HTTPServer struct{}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {

}
func (s *HTTPServer) Get(Path string, handler HandleFunc) {
	s.AddRoute(http.MethodGet, Path, handler)
}
func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}
