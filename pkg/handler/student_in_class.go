package handler

import (
	"net/http"
	"schedule/models"

	"github.com/gin-gonic/gin"
)

type getAllStudentInClassResponse struct {
	Data []models.StudentInClass `json:"data"`
}

func (h *Handler) getStudentInClass(c *gin.Context) {
	studentId := c.Param("Id")
	classId := c.Param("ClassId")

	studentInClass, err := h.service.StudentInClass.GetStudentInClass(studentId, classId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllStudentInClassResponse{
		Data: studentInClass,
	})
}

func (h *Handler) createStudentInClass(c *gin.Context) {
	var studentInClass models.StudentInClass

	if err := c.BindJSON(&studentInClass); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.StudentInClass.CreateStudentInClass(studentInClass)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Id": id,
	})
}