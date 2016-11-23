package main

import (
    "net/http"
    "log"
    "encoding/json"
)


type TestJSON struct {
  Name    string
  Types []string
}

func main() {

  var addr = "localhost:8090"

  log.Printf("Listening at http://%s", addr)

  http.HandleFunc("/", testHandler)
  http.ListenAndServe(addr, nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
  profile := TestJSON{"HelloWorld", []string{"hello", "world"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
