package main

import (
	"fmt"

	"github.com/neeerp/templ-blog/templates"

	"github.com/a-h/templ"
	"net/http"
)

func main() {
	http.Handle("/", templ.Handler(templates.Home()))
	http.Handle("/todos", templ.Handler(templates.TodoList()))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
