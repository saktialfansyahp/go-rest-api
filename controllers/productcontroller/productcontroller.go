package productcontroller

import (
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Halo Admin"}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}
func User(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Halo User"}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}