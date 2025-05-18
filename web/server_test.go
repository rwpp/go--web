package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var s Server
	http.ListenAndServe(":8081", s)

}
