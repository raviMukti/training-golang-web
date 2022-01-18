package traininggolangweb

import (
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080", // Bikin dulu servernya
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
