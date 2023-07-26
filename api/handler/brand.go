package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Brand godoc
// @ID create_brand
// @Router /brand [POST]
// @Summary Create Brand
// @Description Create Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param brand body models.CreateBrand true "CreateBrandRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateBrand(c *gin.Context) {

	var createBrand models.CreateBrand

	err := c.ShouldBindJSON(&createBrand)
	if err != nil {
		h.handlerResponse(c, "create brand", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Brand().Create(context.Background(), &createBrand)
	if err != nil {
		h.handlerResponse(c, "storage.brand.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create brand", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// Get By ID Brand godoc
// @ID get_by_id_brand
// @Router /brand/{id} [GET]
// @Summary Get By ID Brand
// @Description Get By ID Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Authorization path string false "Authorization"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdBrand(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get brand by id", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Get List Brand godoc
// @ID get_list_brand
// @Router /brand [GET]
// @Summary Get List Brand
// @Description Get List Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param Authorization header string false "Authorization"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListBrand(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list brand", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list brand", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Brand().GetList(context.Background(), &models.GetListBrandRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list brand response", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Put Brand godoc
// @ID updat_patch_brand
// @Router /brand/{id} [PUT]
// @Summary Update Put Brand
// @Description Update Put Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param brand body models.UpdateBrand true "UpdateBrand"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateBrand(c *gin.Context) {

	var updateBrand models.UpdateBrand

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateBrand)
	if err != nil {
		h.handlerResponse(c, "update brand", http.StatusBadRequest, err.Error())
		return
	}

	updateBrand.Brand_id = id

	rowsAffected, err := h.storages.Brand().Update(context.Background(), &updateBrand)
	if err != nil {
		h.handlerResponse(c, "storage.brand.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.brand.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update brand", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Update Patch Brand godoc
// @ID updat_patch_brand
// @Router /brand/{id} [PATCH]
// @Summary Update Patch Brand
// @Description Update Patch Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param brand body models.PatchRequest true "UpdatPatchBrandRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchBrand(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch brand", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Brand().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.brand.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.brand.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{Brand_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch brand", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Brand godoc
// @ID get_by_id_brand
// @Router /brand/{id} [DELETE]
// @Summary Delete Brand
// @Description Delete Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteBrand(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Brand().Delete(context.Background(), &models.BrandPrimaryKey{Brand_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.brand.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete brand", http.StatusAccepted, id)
}
