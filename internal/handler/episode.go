package handler

import (
	"github.com/gin-gonic/gin"
	"movies_online/internal/model/catalog"
)

type EpisodeItemResponse struct {
	Result struct {
		Item catalog.Episode
	}
}
type EpisodeItemsResponse struct {
	Result struct {
		Items []catalog.Episode
	}
}

// GetEpisode godoc
// @Summary Get episode by ID
// @Description Get detailed information about a TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Success 200 {object} EpisodeItemResponse "Successfully retrieved episode"
// @Failure 400 {object} ErrorResponse "Not found"
// @Router /otus.episode.get/{id} [get]
func (h *Handler[T]) GetEpisode(c *gin.Context) {
	h.getAction(c)
}

// GetListEpisode godoc
// @Summary Get episodes
// @Description Get list information about TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Success 200 {object} EpisodeItemsResponse "Successfully retrieved episode"
// @Router /otus.episode.list [get]
func (h *Handler[T]) GetListEpisode(c *gin.Context) {
	h.getListAction(c)
}

// AddEpisode godoc
// @Summary Create new TV episode
// @Description Add a new episode to the catalog
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param episode body catalog.Episode true "Episode data"
// @Success 200 {object} EpisodeItemResponse
// @Security ApiKeyAuth
// @Router /otus.episode.add [post]
func (h *Handler[T]) AddEpisode(c *gin.Context) {
	h.addAction(c)
}

// UpdateEpisode godoc
// @Summary Update episode
// @Description Update existing TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Param episode body catalog.Episode true "Update data"
// @Success 200 {object} EpisodeItemResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.episode.update/{id} [put]
func (h *Handler[T]) UpdateEpisode(c *gin.Context) {
	h.updateAction(c)
}

// DeleteEpisode godoc
// @Summary Delete episode
// @Description Delete a TV episode from catalog
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.episode.delete/{id} [delete]
func (h *Handler[T]) DeleteEpisode(c *gin.Context) {
	h.deleteAction(c)
}
