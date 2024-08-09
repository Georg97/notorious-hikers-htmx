package main

import (
    "fmt"
    "net/http"

    "georg97/notorious-hikers-htmx/pkg/env"

    "github.com/a-h/templ"
)

func main() {
    env.InitiateEnv()
    component := hikersHello("Georg")

    http.Handle("/", templ.Handler(component))
    http.HandleFunc("/test", func (w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "This is the test")
    })

    fmt.Printf("Listening on :%v\n", env.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", env.Port), nil)
}
