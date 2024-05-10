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
	StudentTaskID uint
	Files         []FileData
}

func (h *Handler) createReport(c *gin.Context) {
	var files FileUpload
	if err := c.ShouldBindJSON(&files); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var data models.Report
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

	id, err := h.service.Report.CreateReport(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	solutionCtx, ok := Solutions.Load(files.StudentTaskID)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Fuck")
		return
	}

	solutionContext := solutionCtx.(models.StudentSolution)
	solutionContext.ReportID = id
	Solutions.Store(files.StudentTaskID, solutionContext)

	c.JSON(http.StatusOK, gin.H{
		"dataId": id,
	})
}