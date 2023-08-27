package handler

import (
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/router"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  router.DefineRoutes()
}