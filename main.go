package main

import (
    "github.com/gin-gonic/gin"
    "github.com/kostas1721/product-api/database"
	"github.com/kostas1721/product-api/handlers"
)

func main() {
    router := gin.Default()
    database.InitDB()

    // Routes for CRUD operations
    router.GET("/products", handlers.GetProducts)
    router.GET("/products/:id", handlers.GetProductByID)
    router.POST("/products", handlers.CreateProduct)
    router.PUT("/products/:id", handlers.UpdateProduct)
    router.DELETE("/products/:id", handlers.DeleteProduct)

    router.Run(":8080")
}
