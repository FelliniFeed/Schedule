package handler

import (
	"net/http"
	"schedule/models"

	"github.com/gin-gonic/gin"
)

type getAllListsResponse struct {
	Data []models.Class `json:"data"`
}

func (h *Handler) getClasses(c *gin.Context) {

	class, err := h.service.Class.GetClasses()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: class,
	})
}

func (h *Handler) createClass(c *gin.Context) {
	var class models.Class

	if err := c.BindJSON(&class); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Class.CreateClass(class )
	if err !=nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}