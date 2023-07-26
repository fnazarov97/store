package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "CreateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder)
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: createOrder.Customer_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.create.GetCustomerByIDForCreateOrder", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: createOrder.Store_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.create.GetStoreByIDForCreateOrder", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: createOrder.Staff_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.create.GetStaffByIDForCreateOrder", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Order().Create(context.Background(), &createOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// Get By ID Order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{Order_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get order by id", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Get List Order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListOrder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Order().GetList(context.Background(), &models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.order.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list order response", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update Put Order godoc
// @ID updat_patch_order
// @Router /order/{id} [PUT]
// @Summary Update Put Order
// @Description Update Put Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.UpdateOrder true "UpdateOrder"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var updateOrder models.UpdateOrder

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "update order", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{Customer_id: updateOrder.Customer_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.update.GetCustomerByIDForUpdateOrder", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{Store_id: updateOrder.Store_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.update.GetStoreByIDForUpdateOrder", http.StatusNotFound, err.Error())
		return
	}

	_, err = h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{Staff_id: updateOrder.Staff_id})
	if err != nil {
		h.handlerResponse(c, "handler.order.update.GetStaffByIDForUpdateOrder", http.StatusInternalServerError, err.Error())
		return
	}

	updateOrder.Order_id = id

	rowsAffected, err := h.storages.Order().Update(context.Background(), &updateOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{Order_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update order", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Update Patch Order godoc
// @ID updat_patch_order
// @Router /order/{id} [PATCH]
// @Summary Update Patch Order
// @Description Update Patch Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.PatchRequest true "UpdatPatchOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var object models.PatchRequest

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch order", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Order().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.order.patchupdate", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{Order_id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch order", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete Order godoc
// @ID get_by_id_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.storages.Order().Delete(context.Background(), &models.OrderPrimaryKey{Order_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete order", http.StatusAccepted, id)
}

// -------------------------------------------------------------------------------------------
// TASK 5

// @Security ApiKeyAuth
// Create Order Item godoc
// @ID create_order_item
// @Router /order_item [POST]
// @Summary Create Order Item
// @Description Create Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param order_item body models.CreateOrder_item true "CreateOrder_itemRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrderItem(c *gin.Context) {

	var createOrderItem models.CreateOrder_item

	err := c.ShouldBindJSON(&createOrderItem) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create order_item", http.StatusBadRequest, err.Error())
		return
	}
	// get order for getting store_id
	order, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{Order_id: createOrderItem.Order_id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	// check count of products in store
	stockData, err := h.storages.Stock().GetByIdProductStock(context.Background(), order.Store_id, createOrderItem.Product_id)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}
	QuantInt, _ := stockData.Quantity.(int)
	QuantF, _ := stockData.Quantity.(float64)
	if QuantInt <= 0 || createOrderItem.Quantity > QuantF {
		h.handlerResponse(c, "create order_item", http.StatusBadRequest, "Товарь не найден")
		return
	}
	// ----------CREATE ORDER ITEM------------------------------------------------------------------------------------------
	// WHEN create order item in postgres will execute trigger for getting products from store
	_, err = h.storages.Order().AddOrderItem(context.Background(), &models.OrderItem{})
	if err != nil {
		h.handlerResponse(c, "storage.order_item.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{Order_id: createOrderItem.Order_id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Order Item Added", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// DELETE Order Item godoc
// @ID delete_order_item
// @Router /order_item/{id} [DELETE]
// @Summary Delete Order Item
// @Description Delete Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param item_id query string true "item_id"
// @Param orderItem body models.OrderItemPrimaryKey true "DeleteOrderItemRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrderItem(c *gin.Context) {

	id := c.Param("id")
	itemId := c.Query("item_id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order_item.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	idItemInt, err := strconv.Atoi(itemId)
	if err != nil {
		h.handlerResponse(c, "storage.order_item.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	_, err = h.storages.Order().RemoveOrderItem(context.Background(), &models.OrderItemPrimaryKey{Order_id: idInt, Item_id: idItemInt})
	if err != nil {
		h.handlerResponse(c, "storage.order_item.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete order_item", http.StatusNoContent, "Deleted succesfully")
}
