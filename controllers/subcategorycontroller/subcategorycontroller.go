package subcategorycontroller

import (
	"encoding/json"
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
	"github.com/saktialfansyahp/go-rest-api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var subcategory []models.Subcategory
	if err := models.DB.Preload("Category").Find(&subcategory).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	response := map[string]interface{}{"message": "success", "data": subcategory}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var subcategory models.Subcategory
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&subcategory); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Preload("Category").Create(&subcategory).Find(&subcategory).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "success", "data": subcategory}
	helper.ResponseJSON(w, http.StatusOK, response)
}