package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"merchant/internal/entity"
	"merchant/internal/service"
	"merchant/internal/usecase"
	"merchant/pkg/response"
	"net/http"
)

type ProductHandler struct {
	productService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	productHandler := ProductHandler{}
	productHandler.productService = productService
	return &productHandler
}

func (h *ProductHandler) EditStore(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	err = c.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	//var id int
	//jsonData, _ := ioutil.ReadAll(c.Request.Body)
	//err = json.Unmarshal(jsonData, &id)
	//if err != nil {
	//	response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
	//}

	var form entity.FormData

	if err := c.ShouldBindWith(&form, binding.FormMultipart); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := form.ID
	productMap, err := usecase.Parse(file.Filename, id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	info, err := h.productService.EditStore(productMap)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if len(info) == 0 {
		info = "Изменений не было"
	}

	response.ResponseWithData(c, info)
}
