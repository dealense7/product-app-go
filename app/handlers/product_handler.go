package handlers

import (
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service interfaces.ProductService
}

func NewProductHandler(service interfaces.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts()
	if err != nil {
		c.HTML(500, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "index.html", gin.H{ // Changed from "index.html" to "index"
		"Title":    "Products List",
		"products": products,
	})
}
