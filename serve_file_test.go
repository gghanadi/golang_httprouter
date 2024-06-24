package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServerFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources") // fungsi disini agar nama link nya tidak ada resources nya langsung ke nama file nya
	router.ServeFiles("/files/*filepath", http.FS(directory))
	request := httptest.NewRequest("GET", "http://localhost:8081/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Good Bye Http Router", string(body))
}
