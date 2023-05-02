package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "time"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%v", time.Now())
    })

    http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "ok")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}