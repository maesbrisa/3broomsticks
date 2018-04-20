package routes

import (
  "github.com/gorilla/mux"
  "github.com/maesbrisa/3broomsticks/general"
)

func RouteMain(r *mux.Router) {
  r.Methods("GET").Path("/api/main").HandlerFunc(general.List)	
}
