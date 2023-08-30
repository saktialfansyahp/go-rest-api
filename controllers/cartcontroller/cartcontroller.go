package cartcontroller

import (
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "cart"}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}