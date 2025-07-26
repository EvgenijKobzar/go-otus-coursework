package handler

import (
	"github.com/gin-gonic/gin"
	"movies_online/internal/model/catalog"
)

type SerialItemResponse struct {
	Result struct {
		Item catalog.Serial
	}
}
type SerialItemsResponse struct {
	Result struct {
		Items []catalog.Serial
	}
}

// GetSerial godoc
// @Summary Get serial by ID
// @Description Get detailed information about a TV serial
// @Tags serials
// @Accept  json
// @Produce  json
// @Param id path int true "Serial ID"
// @Success 200 {object} SerialItemResponse "Successfully retrieved serial"
// @Failure 400 {object} ErrorResponse "Not found"
// @Router /otus.serial.get/{id} [get]
func (h *Handler[T]) GetSerial(c *gin.Context) {
	h.getAction(c)
}

// GetListSerial godoc
// @Summary Get serials
// @Description Get list information about TV series
// @Tags serials
// @Accept  json
// @Produce  json
// @Success 200 {object} SerialItemsResponse "Successfully retrieved serial"
// @Router /otus.serial.list [get]
func (h *Handler[T]) GetListSerial(c *gin.Context) {
	h.getListAction(c)
}

// AddSerial godoc
// @Summary Create new TV serial
// @Description Add a new serial to the catalog
// @Tags serials
// @Accept  json
// @Produce  json
// @Param serial body catalog.Serial true "Serial data"
// @Success 200 {object} SerialItemResponse
// @Security ApiKeyAuth
// @Router /otus.serial.add [post]
func (h *Handler[T]) AddSerial(c *gin.Context) {
	h.addAction(c)
}

// UpdateSerial godoc
// @Summary Update serial
// @Description Update existing TV serial
// @Tags serials
// @Accept  json
// @Produce  json
// @Param id path int true "Serial ID"
// @Param serial body catalog.Serial true "Update data"
// @Success 200 {object} SerialItemResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.serial.update/{id} [put]
func (h *Handler[T]) UpdateSerial(c *gin.Context) {
	h.updateAction(c)
}

// DeleteSerial godoc
// @Summary Delete serial
// @Description Delete a TV serial from catalog
// @Tags serials
// @Accept  json
// @Produce  json
// @Param id path int true "Serial ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.serial.delete/{id} [delete]
func (h *Handler[T]) DeleteSerial(c *gin.Context) {
	h.deleteAction(c)
}
