//go:generate go run ../statik.go -src=./public

package main

import (
	"log"
	"net/http"

	"github.com/FZambia/statik/example/statik"
	"github.com/FZambia/statik/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New(statik.Data)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	log.Println("visit http://localhost:8080/public/hello.txt")
	http.ListenAndServe(":8080", nil)
}
