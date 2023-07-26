package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Stock godoc
// @ID create_stock
// @Router /stock [POST]
// @Summary Create Stock
// @Description Create Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body models.CreateStock true "CreateStockRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStock(c *gin.Context) {

	var createStock models.CreateStock

	err := c.ShouldBindJSON(&createStock)
	if err != nil {
		h.handlerResponse(c, "create stock", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: createStock.Store_id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.create.GetStoreByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: createStock.Product_id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.create.GetProductByID", http.StatusNotFound, err.Error())
		return
	}

	id, err := h.storages.Stock().Create(context.Background(), &createStock)
	if err != nil {
		h.handlerResponse(c, "storage.stock.create", http.StatusInternalServerError, err.Error())
		return
	}
	ID, _ := strconv.Atoi(id)
	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{Store_id: ID})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create stock", http.StatusCreated, resp)
}

// Get By ID Stock godoc
// @ID get_by_id_stock
// @Router /stock/{id} [GET]
// @Summary Get By ID Stock
// @Description Get By ID Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStock(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get stock by id", http.StatusOK, resp)
}

// Get List Stock godoc
// @ID get_list_stock
// @Router /stock [GET]
// @Summary Get List Stock
// @Description Get List Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStock(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list stock", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list stock", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Stock().GetList(context.Background(), &models.GetListStockRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list stock response", http.StatusOK, resp)
}

// Update Put Stock godoc
// @ID updat_patch_stock
// @Router /stock/{id} [PUT]
// @Summary Update Put Stock
// @Description Update Put Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param stock body models.UpdateStock true "UpdateStock"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStock(c *gin.Context) {

	var updateStock models.UpdateStock

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateStock)
	if err != nil {
		h.handlerResponse(c, "update stock", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: updateStock.Store_id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.update.GetStoreByID", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Product_id: updateStock.Product_id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.update.GetProductByID", http.StatusNotFound, err.Error())
		return
	}

	updateStock.Store_id = id

	rowsAffected, err := h.storages.Stock().Update(context.Background(), &updateStock)
	if err != nil {
		h.handlerResponse(c, "storage.stock.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.stock.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update stock", http.StatusAccepted, resp)
}

// Update Patch Stock godoc
// @ID updat_patch_stock
// @Router /stock/{id} [PATCH]
// @Summary Update Patch Stock
// @Description Update Patch Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param stock body models.PatchRequest true "UpdatPatchStockRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchStock(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch stock", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Stock().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.stock.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.stock.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{Store_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch stock", http.StatusAccepted, resp)
}

// Delete Stock godoc
// @ID get_by_id_stock
// @Router /stock/{id} [DELETE]
// @Summary Delete Stock
// @Description Delete Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStock(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Stock().Delete(context.Background(), &models.StockPrimaryKey{Store_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.stock.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete stock", http.StatusAccepted, id)
}
