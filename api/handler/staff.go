package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Staff godoc
// @ID create_staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param staff body models.CreateStaff true "CreateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStaff(c *gin.Context) {

	var createStaff models.CreateStaff

	err := c.ShouldBindJSON(&createStaff)
	if err != nil {
		h.handlerResponse(c, "create staff", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: createStaff.Store_id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.create.GetStoreByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: createStaff.Manager_id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.create.GetManagerByID", http.StatusNotFound, err.Error())
		return
	}

	id, err := h.storages.Staff().Create(context.Background(), &createStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff", http.StatusCreated, resp)
}

// Get By ID Staff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [GET]
// @Summary Get By ID Staff
// @Description Get By ID Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStaff(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get staff by id", http.StatusOK, resp)
}

// Get List Staff godoc
// @ID get_list_staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStaff(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Staff().GetList(context.Background(), &models.GetListStaffRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff response", http.StatusOK, resp)
}

// Update Put Staff godoc
// @ID updat_patch_staff
// @Router /staff/{id} [PUT]
// @Summary Update Put Staff
// @Description Update Put Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param staff body models.UpdateStaff true "UpdateStaff"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStaff(c *gin.Context) {

	var updateStaff models.UpdateStaff

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateStaff)
	if err != nil {
		h.handlerResponse(c, "update staff", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: updateStaff.Store_id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.update.GetStoreByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: updateStaff.Manager_id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.update.GetManagerByID", http.StatusNotFound, err.Error())
		return
	}

	updateStaff.Staff_id = id

	rowsAffected, err := h.storages.Staff().Update(context.Background(), &updateStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update staff", http.StatusAccepted, resp)
}

// Update Patch Staff godoc
// @ID updat_patch_staff
// @Router /staff/{id} [PATCH]
// @Summary Update Patch Staff
// @Description Update Patch Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param staff body models.PatchRequest true "UpdatPatchStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchStaff(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch staff", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Staff().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.staff.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch staff", http.StatusAccepted, resp)
}

// Delete Staff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStaff(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Staff().Delete(context.Background(), &models.StaffPrimaryKey{Staff_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete staff", http.StatusAccepted, id)
}
