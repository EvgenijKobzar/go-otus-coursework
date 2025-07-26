package handler

import (
	"github.com/gin-gonic/gin"
	"movies_online/internal/core"
	"movies_online/internal/middleware"
	"movies_online/internal/model/catalog"
	"strconv"
)

type DeleteResponse struct {
	Result struct {
		Deleted bool
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"entity not found"`
}

const Item = "item"
const Items = "items"

type Handler[T catalog.HasId] struct {
	service *core.Service[T]
}

func New[T catalog.HasId](service *core.Service[T]) *Handler[T] {
	return &Handler[T]{service: service}
}

// Action region
func (h *Handler[T]) getAction(c *gin.Context) {
	var entity T
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		entity, err = h.service.GetInner(id)
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func (h *Handler[T]) addAction(c *gin.Context) {
	var entity *T
	bindings := new(T)

	err := c.ShouldBind(bindings)

	if err == nil {
		entity, err = h.service.AddInner(bindings)
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func (h *Handler[T]) updateAction(c *gin.Context) {
	var entity T
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		var inputFields map[string]any
		if err = c.ShouldBindJSON(&inputFields); err == nil {
			entity, err = h.service.UpdateInner(id, inputFields)
		}
	}
	setResponse(gin.H{Item: entity}, err, c)
}

func (h *Handler[T]) deleteAction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		err = h.service.DeleteInner(id)
	}
	setResponse(gin.H{"deleted": true}, err, c)
}

func (h *Handler[T]) getListAction(c *gin.Context) {
	var err error
	var result gin.H

	items, _ := h.service.GetListInner(c.QueryMap("filter"), c.QueryMap("order"))
	result = gin.H{Items: items}

	setResponse(result, err, c)
}

// end region

func setResponse(result gin.H, err error, c *gin.Context) {
	if err == nil {
		c.Set(middleware.KeyResponse, result)
	} else {
		c.Set(middleware.KeyError, err)
	}
}
