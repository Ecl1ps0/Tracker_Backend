package handler

import (
	"Proctor/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileData struct {
	Name string
	Data string
}

type FileUpload struct {
	Files []FileData
}

func (h *Handler) fileUpload(c *gin.Context) {
	var files FileUpload
	if err := c.ShouldBindJSON(&files); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var data models.Data
	for _, file := range files.Files {
		switch file.Name {
		case "logs":
			data.Logs = []byte(file.Data)
		case "report":
			data.Report = []byte(file.Data)
		case "clipboard":
			data.Clipboard = []byte(file.Data)
		}
	}

	id, err := h.service.FileHandler.SaveFile(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"dataId": id,
	})
}
