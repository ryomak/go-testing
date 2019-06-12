package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/user",userHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}


