package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.CreateProduct true "CreateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		h.handlerResponse(c, "create product", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: createProduct.Brand_id})
	if err != nil {
		h.handlerResponse(c, "storage.product.create.GetBrandByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: createProduct.Category_id})
	if err != nil {
		h.handlerResponse(c, "storage.product.create.GetCategoryByID", http.StatusNotFound, err.Error())
		return
	}

	id, err := h.storages.Product().Create(context.Background(), &createProduct)
	if err != nil {
		h.handlerResponse(c, "storage.product.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}
	err = h.caches.ProductCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache product:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.handlerResponse(c, "create product", http.StatusCreated, resp)
}

// Get By ID Product godoc
// @ID get_by_id_product
// @Router /product/{id} [GET]
// @Summary Get By ID Product
// @Description Get By ID Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get product by id", http.StatusOK, resp)
}

// Get List Product godoc
// @ID get_list_product
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid limit")
		return
	}

	ok, er := h.caches.ProductCache().Exists()
	fmt.Println("-------------->", ok, er)
	if er != nil {
		log.Println("error whiling check redis:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var resp *models.GetListProductResponse

	if !ok {
		resp, err = h.storages.Product().GetList(context.Background(), &models.GetListProductRequest{
			Offset: offset,
			Limit:  limit,
			Search: c.Query("search"),
		})
		if err != nil {
			h.handlerResponse(c, "storage.product.getlist", http.StatusInternalServerError, err.Error())
			return
		}

		err = h.caches.ProductCache().Create(resp)

		if err != nil {
			log.Println("error whiling set product_list to redis:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("Postgress------------------------")
	} else {
		resp, err = h.caches.ProductCache().GetList()
		if err != nil {
			log.Println("error whiling get list product from redis:", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("Redis------------------------")

	}

	h.handlerResponse(c, "get list product response", http.StatusOK, resp)
}

// Update Put Product godoc
// @ID updat_patch_product
// @Router /product/{id} [PUT]
// @Summary Update Put Product
// @Description Update Put Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param product body models.UpdateProduct true "UpdateProduct"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var updateProduct models.UpdateProduct

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "update product", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: updateProduct.Brand_id})
	if err != nil {
		h.handlerResponse(c, "storage.product.update.GetBrandByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: updateProduct.Category_id})
	if err != nil {
		h.handlerResponse(c, "storage.product.update.GetCategoryByID", http.StatusNotFound, err.Error())
		return
	}

	updateProduct.Product_id = id

	rowsAffected, err := h.storages.Product().Update(context.Background(), &updateProduct)
	if err != nil {
		h.handlerResponse(c, "storage.product.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.product.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.caches.ProductCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache product:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(c, "update product", http.StatusAccepted, resp)
}

// Update Patch Product godoc
// @ID updat_patch_product
// @Router /product/{id} [PATCH]
// @Summary Update Patch Product
// @Description Update Patch Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param product body models.PatchRequest true "UpdatPatchProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchProduct(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch product", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Product().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.product.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.product.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.caches.ProductCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache product:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(c, "update patch product", http.StatusAccepted, resp)
}

// Delete Product godoc
// @ID get_by_id_product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Product().Delete(context.Background(), &models.ProductPrimaryKey{Product_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.product.delete", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.caches.ProductCache().Delete()
	if err != nil {
		log.Println("error whiling delete cache product:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(c, "delete product", http.StatusAccepted, id)
}
