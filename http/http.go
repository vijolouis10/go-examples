package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "<html><h1>Hello %s!</h1></html>",req.URL.Path[1:])
}

// func headers(w http.ResponseWriter, req *http.Request) {

//     for name, headers := range req.Header {
//         for _, h := range headers {
//             fmt.Fprintf(w, "%v: %v\n", name, h)
//         }
//     }
// }

func main() {

    http.HandleFunc("/", hello)
    // http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}
