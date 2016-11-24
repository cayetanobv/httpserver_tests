// Author: Cayetano Benavent, 2016.
// cayetano.benavent@geographica.gs
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
// MA 02110-1301, USA.


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
