package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Store godoc
// @ID create_store
// @Router /store [POST]
// @Summary Create Store
// @Description Create Store
// @Tags Store
// @Accept json
// @Produce json
// @Param store body models.CreateStore true "CreateStoreRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStore(c *gin.Context) {

	var createStore models.CreateStore

	err := c.ShouldBindJSON(&createStore)
	if err != nil {
		h.handlerResponse(c, "create store", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Store().Create(context.Background(), &createStore)
	if err != nil {
		h.handlerResponse(c, "storage.store.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create store", http.StatusCreated, resp)
}

// Get By ID Store godoc
// @ID get_by_id_store
// @Router /store/{id} [GET]
// @Summary Get By ID Store
// @Description Get By ID Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStore(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get store by id", http.StatusOK, resp)
}

// Get List Store godoc
// @ID get_list_store
// @Router /store [GET]
// @Summary Get List Store
// @Description Get List Store
// @Tags Store
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStore(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list store", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list store", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Store().GetList(context.Background(), &models.GetListStoreRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.store.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list store response", http.StatusOK, resp)
}

// Update Put Store godoc
// @ID updat_patch_store
// @Router /store/{id} [PUT]
// @Summary Update Put Store
// @Description Update Put Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.UpdateStore true "UpdateStore"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStore(c *gin.Context) {

	var updateStore models.UpdateStore

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateStore)
	if err != nil {
		h.handlerResponse(c, "update store", http.StatusBadRequest, err.Error())
		return
	}

	updateStore.Store_id = id

	rowsAffected, err := h.storages.Store().Update(context.Background(), &updateStore)
	if err != nil {
		h.handlerResponse(c, "storage.store.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.store.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update store", http.StatusAccepted, resp)
}

// Update Patch Store godoc
// @ID updat_patch_store
// @Router /store/{id} [PATCH]
// @Summary Update Patch Store
// @Description Update Patch Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.PatchRequest true "UpdatPatchStoreRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchStore(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch store", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Store().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.store.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.store.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch store", http.StatusAccepted, resp)
}

// Delete Store godoc
// @ID get_by_id_store
// @Router /store/{id} [DELETE]
// @Summary Delete Store
// @Description Delete Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStore(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Store().Delete(context.Background(), &models.StorePrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.store.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete store", http.StatusAccepted, id)
}
