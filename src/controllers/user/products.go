package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	services "github.com/fathimasithara01/ecommerce/src/services/user"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductController(productUsecase services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productUsecase}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	products, err := h.productService.SearchProducts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"

// 	"github.com/fathimasithara01/ecommerce/src/services"
// 	"github.com/fathimasithara01/ecommerce/utils/models"
// 	"github.com/fathimasithara01/ecommerce/utils/response"
// )

// func CreateProduct(c *gin.Context) {
// 	var product models.Product
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invallid data", nil, err.Error()))
// 		return
// 	}

// 	if err := validator.New().Struct(product); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
// 		return
// 	}

// 	created, err := services.CreateProduct(product)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to add product", nil, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "product created", created, nil))
// }

// func GetAllProduct(c *gin.Context) {
// 	products, err := services.GetAllProduct()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "couldnot retrieve products", nil, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "product created", products, err))
// }

// func UpdateProduct(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid ID", nil, err.Error()))
// 		return
// 	}

// 	var update models.ProductUpdate
// 	if err := c.ShouldBindJSON(&update); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid data", nil, err.Error()))
// 		return
// 	}
// 	if err := validator.New().Struct(update); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
// 		return
// 	}

// 	updated, err := services.UpdateProduct(uint(id), update)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "update failed", nil, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "product updated", updated, nil))
// }

// func DeleteProduct(c *gin.Context) {
// 	idParam := c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid ID", nil, err.Error()))
// 		return
// 	}

// 	if err := services.DeleteProduct(uint(id)); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "delete failed", nil, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "product deleted", nil, nil))
// }

// func SearchProduct(c *gin.Context) {
// 	query := c.Query("query")
// 	if query == "" {
// 		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "search query is missiing", nil, nil))
// 		return
// 	}

// 	results, err := services.SearchProduct(query)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "search failed", nil, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "products found", results, nil))
// }
