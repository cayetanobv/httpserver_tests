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
  http.HandleFunc("/", testHandler)
  log.Print("Listening at http://localhost:8090")
  http.ListenAndServe(":8090", nil)
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
