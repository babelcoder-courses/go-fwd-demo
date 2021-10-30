package util

import (
	"net/http"
	"net/http/httptest"
)

func StartServer(content []byte) *httptest.Server {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)
			rw.Write(content)
		}),
	)

	return server
}
