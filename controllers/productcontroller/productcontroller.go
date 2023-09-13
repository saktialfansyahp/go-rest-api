package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/helper"
	"github.com/saktialfansyahp/go-rest-api/models"
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

func Index(w http.ResponseWriter, r *http.Request){
	var product []models.Product_Color
	if err := models.DB.Preload("Product").Preload("Product.Subcategory").Preload("Product.Subcategory.Category").
	Preload("Color").Find(&product).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	// response := map[string]interface{}{"message":"success", "data": product}
	helper.ResponseJSON(w, http.StatusOK, product)
}
func ById(w http.ResponseWriter, r *http.Request, id string){
	var product models.Product_Color
	if err := models.DB.Preload("Product").Preload("Product.Subcategory").Preload("Product.Subcategory.Category").
	Preload("Color").Where("product_id = ?", id).First(&product).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	// response := map[string]interface{}{"message":"success", "data": product, "id": id}
	helper.ResponseJSON(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
    var productInput models.ProductInput
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&productInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	product := models.Product{
		Image: productInput.Image,
		ProductName: productInput.ProductName,
		Description: productInput.Description,
		Price: productInput.Price,
		SubcategoryID: productInput.SubcategoryID,
	}

	if  err := models.DB.Create(&product).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	productcolor := models.Product_Color{
		ProductID: product.Id,
		ColorID: productInput.ColorID,
	}

	if err := models.DB.Preload("Product").Preload("Product.Subcategory").
	Preload("Product.Subcategory.Category").Preload("Color").Create(&productcolor).Find(&productcolor).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{"message": "success", "data": productcolor}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Edit(w http.ResponseWriter, r *http.Request, id string) {
    var productInput models.ProductInput
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&productInput); err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusBadRequest, response)
        return
    }
    defer r.Body.Close()

    // Cek apakah produk dengan product_id tersebut ada dalam database
    var existingProduct models.Product
    if err := models.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
        response := map[string]string{"message": "Produk tidak ditemukan"}
        helper.ResponseJSON(w, http.StatusNotFound, response)
        return
    }

    // Update data produk berdasarkan product_id
    existingProduct.Image = productInput.Image
    existingProduct.ProductName = productInput.ProductName
    existingProduct.Description = productInput.Description
    existingProduct.Price = productInput.Price
    existingProduct.SubcategoryID = productInput.SubcategoryID

    if err := models.DB.Save(&existingProduct).Error; err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // Update data product_color berdasarkan product_id
    var existingProductColor models.Product_Color
    if err := models.DB.Where("product_id = ?", id).First(&existingProductColor).Error; err != nil {
        response := map[string]string{"message": "Product Color tidak ditemukan"}
        helper.ResponseJSON(w, http.StatusNotFound, response)
        return
    }

    existingProductColor.ColorID = productInput.ColorID
    if err := models.DB.Save(&existingProductColor).Error; err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    response := map[string]interface{}{"message": "Perubahan sukses", "data": existingProductColor}
    helper.ResponseJSON(w, http.StatusOK, response)
}

func Delete(w http.ResponseWriter, r *http.Request, id string) {

    var existingProduct models.Product
    if err := models.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
        response := map[string]string{"message": "Produk tidak ditemukan"}
        helper.ResponseJSON(w, http.StatusNotFound, response)
        return
    }

    if err := models.DB.Delete(&existingProduct).Error; err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    if err := models.DB.Where("product_id = ?", id).Delete(&models.Product_Color{}).Error; err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    response := map[string]string{"message": "Produk berhasil dihapus"}
    helper.ResponseJSON(w, http.StatusOK, response)
}
