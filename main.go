package main

import (
  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "github.com/urfave/negroni"
  "net/http"
  "time"
  "fmt"
  "github.com/maesbrisa/3broomsticks/routes"
)

const (
  Port = "8080"
)

func main() {
  c := cors.AllowAll()
  n := negroni.New()
  n.Use(c)
  r := mux.NewRouter().StrictSlash(false)
  r.KeepContext = true
 
  routes.RouteMain(r)
  
  n.UseHandler(r)
  s := &http.Server {
    Addr: ":" + Port,
    Handler: n,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  fmt.Sprintf("Initializing...")
}
