package main

import (
	"fmt"

	"blog/templates"

	"github.com/a-h/templ"
	"net/http"
)

func main() {
	component := templates.Home()
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
