package main

import "net/http"

func main() {
   http.HandleFunc("/", Hello)
   http.ListenAndServe(":9000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("<h1> Hello Kubernetes</h1>"))
}