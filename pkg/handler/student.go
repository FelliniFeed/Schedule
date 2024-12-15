package handler

import (
	"net/http"
	"schedule/models"

	"github.com/gin-gonic/gin"
)

type getAllStudentsResponse struct {
	Data []models.Student `json:"data"`
}

func (h *Handler) getStudents(c *gin.Context) {

	students, err := h.service.Student.GetStudents()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllStudentsResponse{
		Data: students,
	})
}

func (h *Handler) createStudent(c *gin.Context) {
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Student.CreateStudent(student)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}