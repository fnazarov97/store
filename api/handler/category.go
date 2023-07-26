package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Category godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "CreateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		h.handlerResponse(c, "create category", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Category().Create(context.Background(), &createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create category", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// Get By ID Category godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCategory(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get category by id", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Get List Category godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Category().GetList(context.Background(), &models.GetListCategoryRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.category.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list category response", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Put Category godoc
// @ID updat_patch_category
// @Router /category/{id} [PUT]
// @Summary Update Put Category
// @Description Update Put Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.UpdateCategory true "UpdateCategory"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCategory(c *gin.Context) {

	var updateCategory models.UpdateCategory

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "update category", http.StatusBadRequest, err.Error())
		return
	}

	updateCategory.Category_id = id

	rowsAffected, err := h.storages.Category().Update(context.Background(), &updateCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update category", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Update Patch Category godoc
// @ID updat_patch_category
// @Router /category/{id} [PATCH]
// @Summary Update Patch Category
// @Description Update Patch Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.PatchRequest true "UpdatPatchCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCategory(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch category", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Category().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.category.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Category_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch category", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Category godoc
// @ID get_by_id_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCategory(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Category().Delete(context.Background(), &models.CategoryPrimaryKey{Category_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete category", http.StatusAccepted, id)
}
