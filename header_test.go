package traininggolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	fmt.Fprint(rw, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json") // Add header
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("X-Powered-By", "Afterwork")

	fmt.Fprint(rw, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("x-powered-by")

	fmt.Println(poweredBy)
}
