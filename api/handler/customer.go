package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Customer godoc
// @ID create_customer
// @Router /customer [POST]
// @Summary Create Customer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.CreateCustomer true "CreateCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCustomer(c *gin.Context) {

	var createCustomer models.CreateCustomer

	err := c.ShouldBindJSON(&createCustomer)
	if err != nil {
		h.handlerResponse(c, "create customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Customer().Create(context.Background(), &createCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create customer", http.StatusCreated, resp)
}

// Get By ID Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [GET]
// @Summary Get By ID Customer
// @Description Get By ID Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCustomer(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get customer by id", http.StatusOK, resp)
}

// Get List Customer godoc
// @ID get_list_customer
// @Router /customer [GET]
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCustomer(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Customer().GetList(context.Background(), &models.GetListCustomerRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list customer response", http.StatusOK, resp)
}

// Update Put Customer godoc
// @ID updat_patch_customer
// @Router /customer/{id} [PUT]
// @Summary Update Put Customer
// @Description Update Put Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.UpdateCustomer true "UpdateCustomer"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCustomer(c *gin.Context) {

	var updateCustomer models.UpdateCustomer

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		h.handlerResponse(c, "update customer", http.StatusBadRequest, err.Error())
		return
	}

	updateCustomer.Customer_id = id

	rowsAffected, err := h.storages.Customer().Update(context.Background(), &updateCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update customer", http.StatusAccepted, resp)
}

// Update Patch Customer godoc
// @ID updat_patch_customer
// @Router /customer/{id} [PATCH]
// @Summary Update Patch Customer
// @Description Update Patch Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.PatchRequest true "UpdatPatchCustomerRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCustomer(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch customer", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Customer().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.customer.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch customer", http.StatusAccepted, resp)
}

// Delete Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [DELETE]
// @Summary Delete Customer
// @Description Delete Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCustomer(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Customer().Delete(context.Background(), &models.CustomerPrimaryKey{Customer_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete customer", http.StatusAccepted, id)
}
