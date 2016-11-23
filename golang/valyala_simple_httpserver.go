package main

import (
  "log"
  "encoding/json"
  "github.com/valyala/fasthttp"
)


type TestJSON struct {
  Name    string
  Types []string
}

func main() {

  var addr = "localhost:8090"

  log.Printf("Listening at http://%s", addr)

  h := requestHandler

  if err := fasthttp.ListenAndServe(addr, h); err != nil {
    log.Fatalf("Error in ListenAndServe: %s", err)
  }
}

func requestHandler(ctx *fasthttp.RequestCtx) {
  profile := TestJSON{"HelloWorld", []string{"hello", "world"}}
  js, err := json.Marshal(profile)
  if err != nil {
    log.Fatalf("Error: %s", err)
    return
  }

  ctx.SetContentType("application/json; charset=utf8")
  ctx.Write(js)
}
