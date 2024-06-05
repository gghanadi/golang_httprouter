package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// kalo menggunakan library httprouter cara hampir mirip denghan server mux tetapi ini menggunakan httprouter.New() sama seperti contoh di bawah
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8081",
	}

	server.ListenAndServe()
}
