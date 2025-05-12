package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	uc *usecase.InventoryUsecase
}

func NewHandler(uc *usecase.InventoryUsecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProduct)
	r.PATCH("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/products", h.ListProducts)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var p entity.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(c, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	p, err := h.uc.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p entity.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Update(c, id, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) ListProducts(c *gin.Context) {
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	skip, _ := strconv.ParseInt(c.DefaultQuery("skip", "0"), 10, 64)
	category := c.Query("category")
	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	products, err := h.uc.List(c, filter, limit, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list"})
		return
	}
	c.JSON(http.StatusOK, products)
}
