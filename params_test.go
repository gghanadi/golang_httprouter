/*
fungsi params disini adalah seperti dynamic website pada alamat website seperti product/1 prodcut/2
seperti itu
jadi hampir sama kaya laraverl /product/{id?}

*/

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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// untuk name parameter harus ada ":"
		// byname harus sama dengan judul title web nya (kalo :id ya maka harus id juga)
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})
	request := httptest.NewRequest("GET", "http://localhost:8081/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
