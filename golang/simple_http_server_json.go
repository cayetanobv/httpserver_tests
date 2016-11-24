
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
