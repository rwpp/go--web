//go:build e2e

package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := NewHTTPServer()
	handler1 := func(ctx *Context) {
		fmt.Println("1")
	}
	handler2 := func(ctx *Context) {
		fmt.Println("2")
	}
	h.addRoute(http.MethodGet, "/user", func(ctx *Context) {
		handler1(ctx)
		handler2(ctx)
	})
	h.addRoute(http.MethodGet, "/order/detail", func(ctx *Context) {
		ctx.Resp.Write([]byte("order"))
	})
	h.Start(":8081")
}
