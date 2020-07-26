package hello_go

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)
func Init() {
	fmt.Print("this is hello go v0.1.0")
	mtx := mux.NewRouter()
	mtx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})
	srv := &http.Server{
		Handler:      mtx,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}