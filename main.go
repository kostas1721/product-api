package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/kostas1721/product-api/database"
	"github.com/kostas1721/product-api/handlers"
	"github.com/kostas1721/product-api/models"
)

func main() {
    router := gin.Default()
    database.InitDB()

    // Routes for CRUD operations
    router.GET("/products", handlers.getProducts)
    router.GET("/products/:id", handlers.getProductByID)
    router.POST("/products", handlers.createProduct)
    router.PUT("/products/:id", handlers.updateProduct)
    router.DELETE("/products/:id", handlers.deleteProduct)

    router.Run(":8080")
}
