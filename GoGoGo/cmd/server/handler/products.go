package handler

import (
	"awesomeProject3/internal/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductsGetter interface {
	GetByID(id int) (products.Product, error)
}

type ProductCreator interface {
	ModifyByID(id int, product products.Product) (products.Product, error)
}

type ProductsHandler struct {
	productsGetter  ProductsGetter
	productsCreator ProductCreator
}

func NewProductsHandler(getter ProductsGetter, creator ProductCreator) *ProductsHandler {
	return &ProductsHandler{
		productsGetter:  getter,
		productsCreator: creator,
	}
}

// GetProductByID godoc
// @Summary      Gets a product by id
// @Description  Gets a product by id from the repository
// @Tags         products
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} products.Product
// @Router       /products/{id} [get]
func (ph *ProductsHandler) GetProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	product, err := ph.productsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *ProductsHandler) PutProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	productRequest := products.Product{}
	err = ctx.BindJSON(&productRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ph.productsCreator.ModifyByID(id, productRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, product)
}
