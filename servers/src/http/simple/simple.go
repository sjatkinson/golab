package main

import (
   "fmt"
   "net/http"
   "net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
   s,_ := httputil.DumpRequest(r, true)
   fmt.Fprintf(w, "Dump: %s\n", s)
}

func main() {
   http.HandleFunc("/", handler)
   http.ListenAndServe(":8080", nil)
}
