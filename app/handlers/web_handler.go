package handlers

import (
	"encoding/json"
	"github.com/dealense7/product-app/app/interfaces"
	"github.com/gin-gonic/gin"
)

type WebHandler struct {
	productService  interfaces.ProductService
	currencyService interfaces.CurrencyService
	gasService      interfaces.GasService
}

func NewWebHandler(productService interfaces.ProductService, currencyService interfaces.CurrencyService, gasService interfaces.GasService) *WebHandler {
	return &WebHandler{productService: productService, currencyService: currencyService, gasService: gasService}
}

func (h *WebHandler) GetProducts(c *gin.Context) {
	// Random Products
	products, err := h.productService.GetProducts(map[string]interface{}{})
	if err != nil {
		paintError(c, err)
		return
	}

	// Products grouped by category
	categoryGroups, err := h.productService.GroupByCategory()
	if err != nil {
		paintError(c, err)
		return
	}

	// Currency
	currencies, err := h.currencyService.GetAll()
	if err != nil {
		paintError(c, err)
		return
	}
	currencyJson, _ := json.Marshal(currencies)

	// Gas
	gasRates, err := h.gasService.GetAll()
	if err != nil {
		paintError(c, err)
		return
	}
	gasRatesJson, _ := json.Marshal(gasRates)

	c.HTML(200, "index.html", gin.H{
		"Title":          "Products List",
		"products":       products,
		"categoryGroups": categoryGroups,
		"currencies":     string(currencyJson),
		"gasRates":       string(gasRatesJson),
	})
}

func paintError(c *gin.Context, err error) {
	c.HTML(500, "error.html", gin.H{"error": err.Error()})
}
