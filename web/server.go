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
	addRoute(method, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
	router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	n, ok := h.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || n.handler == nil {
		ctx.Resp.WriteHeader(404)
		_, _ = ctx.Resp.Write([]byte("404 not found"))
		return
	}
	n.handler(ctx)
}

func (s *HTTPServer) Get(Path string, handler HandleFunc) {
	s.addRoute(http.MethodGet, Path, handler)
}

func (s *HTTPServer) Post(Path string, handler HandleFunc) {
	s.addRoute(http.MethodPost, Path, handler)
}

func (s *HTTPServer) Options(Path string, handler HandleFunc) {
	s.addRoute(http.MethodOptions, Path, handler)
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}
