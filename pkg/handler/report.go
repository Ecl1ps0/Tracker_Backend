package handler

import (
	"Proctor/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FileData struct {
	Name string
	Data string
}

type FileUpload struct {
	StudentTaskID uint
	Files         []FileData
}

// @Summary Get all reports
// @Security ApiKeyAuth
// @Tags reports
// @Description get a list of all reports, accessible only by admins
// @ID get-all-reports
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Report "List of all reports"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/reports [get]
func (h *Handler) getAllReports(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roleId, err := h.service.User.GetRoleByUserID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only admins can observe all reports!")
		return
	}

	reports, err := h.service.Report.GetAllReports()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reports)
}

// @Summary Create a report
// @Security ApiKeyAuth
// @Tags reports
// @Description create a new report with uploaded files
// @ID create-report
// @Accept  json
// @Produce  json
// @Param files body FileUpload true "Files for the report"
// @Success 200 {object} map[string]interface{} "Report creation successful with data ID returned"
// @Failure 400 {object} Error "Bad request due to invalid input"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/reports/createReport [post]
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

	var solution models.StudentSolution
	if err = h.service.Redis.Get(c, strconv.Itoa(int(files.StudentTaskID)), &solution); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	solution.ReportID = id
	if err = h.service.Redis.Set(c, strconv.Itoa(int(files.StudentTaskID)), solution, 0); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dataId": id,
	})
}
