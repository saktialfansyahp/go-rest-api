package colorcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
	"github.com/saktialfansyahp/go-rest-api/models"
)

func Index(w http.ResponseWriter, r *http.Request){
	var color []models.Color
	if err := models.DB.Find(&color).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	response := map[string]interface{}{"message":"success", "data": color}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var color models.Color
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&color); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&color).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "success", "data": color}
	helper.ResponseJSON(w, http.StatusOK, response)
}