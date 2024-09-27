package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "product-api/database"
)

func main() {
    router := gin.Default()
    database.InitDB()

    // Routes for CRUD operations
    router.GET("/products", getProducts)
    router.GET("/products/:id", getProductByID)
    router.POST("/products", createProduct)
    router.PUT("/products/:id", updateProduct)
    router.DELETE("/products/:id", deleteProduct)

    router.Run(":8080")
}
