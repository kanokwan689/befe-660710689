package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"inStock"`
}

var products = []Product{
	{ID: "p001", Name: "Laptop Pro X", Category: "Electronics", Price: 45000.00, InStock: true},
	{ID: "p002", Name: "Comfort T-Shirt", Category: "Apparel", Price: 790.00, InStock: true},
	{ID: "p003", Name: "Wireless Mouse", Category: "Electronics", Price: 1200.50, InStock: false},
	{ID: "p004", Name: "Running Shoes", Category: "Apparel", Price: 3200.00, InStock: true},
}

func getProducts(c *gin.Context) {

	categoryQuery := c.Query("category")

	if categoryQuery != "" {
		filteredProducts := []Product{}
		for _, product := range products {

			if strings.EqualFold(product.Category, categoryQuery) {
				filteredProducts = append(filteredProducts, product)
			}
		}
		c.JSON(http.StatusOK, filteredProducts)
		return
	}

	c.JSON(http.StatusOK, products)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Healthy"})
	})

	api := r.Group("/api/v1")
	{

		api.GET("/products", getProducts)
	}

	fmt.Println("Server is running on port 8080")
	r.Run(":8080")
}
