package productcontroller

import (
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Halo"}
	helper.ResponseJSON(w, http.StatusBadRequest, response)
	return
}