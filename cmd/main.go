package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	r := http.NewServeMux()

	frontEnd := http.FileServer(http.Dir("build"))
	r.Handle("/", frontEnd)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8085",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on PORT 8085...OK!")
	log.Fatal(srv.ListenAndServe())

}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "build/index.html")
}
