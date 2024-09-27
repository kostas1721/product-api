package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/kostas1721/product-api/database"
    "github.com/kostas1721/product-api/models"
)

// GET /v1/products
func GetProducts(c *gin.Context) {
    limitStr := c.DefaultQuery("limit", "10")
    pageStr := c.DefaultQuery("page", "1")

    limit, _ := strconv.Atoi(limitStr)
    page, _ := strconv.Atoi(pageStr)
    offset := (page - 1) * limit

    rows, _ := database.DB.Query("SELECT id, name, description, price FROM products LIMIT ? OFFSET ?", limit, offset)
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}

// GET /v1/products/:id
func GetProductByID(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    row := database.DB.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?", id)
    err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// POST /v1/products
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Ensure that name, description, and price are provided
    if product.Name == "" || product.Description == "" || product.Price <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
        return
    }

    result, err := database.DB.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)",
        product.Name, product.Description, product.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, _ := result.LastInsertId()
    product.ID = int(id)

    c.JSON(http.StatusCreated, product)
}

// PUT /v1/products/:id
func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Ensure that name, description, and price are provided
    if product.Name == "" || product.Description == "" || product.Price <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
        return
    }

    _, err := database.DB.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?",
        product.Name, product.Description, product.Price, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}

// DELETE /v1/products/:id
func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    _, err := database.DB.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
