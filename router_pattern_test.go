package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterNamedPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// untuk name parameter harus ada ":"
		// byname harus sama dengan judul title web nya (kalo :id ya maka harus id juga)
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(w, text)
	})
	request := httptest.NewRequest("GET", "http://localhost:8081/product/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

// fungsi pattern all parameter hampir mirip semisal kita input get pada form di laravel
// untuk deklarasikan nya harus menggunakan "*" dan diletakan di akhir
func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image " + image
		fmt.Fprint(w, text)
	})
	request := httptest.NewRequest("GET", "http://localhost:8081/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image /small/profile.png", string(body))
}
