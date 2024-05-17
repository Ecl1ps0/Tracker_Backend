package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get all users
// @Security ApiKeyAuth
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "List of all users"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
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
		newErrorResponse(c, http.StatusForbidden, "Only admins can observe all users!")
		return
	}

	users, err := h.service.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get user
// @Security ApiKeyAuth
// @Tags users
// @Description get user
// @ID get-user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User "The profile of the user"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/users/profile [get]
func (h *Handler) getProfile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.User.GetProfile(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Add student to task
// @Security ApiKeyAuth
// @Tags users
// @Description add student to task
// @ID add-student-to-task
// @Accept  json
// @Produce  json
// @Param  taskID path int true "Task ID"
// @Param  studentID path int true "Student ID"
// @Success 200 {object} nil "Successfully added student to task"
// @Failure 400 {object} Error "Invalid input data"
// @Failure 403 {object} Error "Access denied"
// @Failure 404 {object} Error "Task or user not found"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/{taskID}/students/{studentID} [post]
func (h *Handler) addStudentToTask(c *gin.Context) {
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

	taskId, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.Task.GetTaskByID(uint(taskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.TeacherID != userId && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only Teacher created task and Admin can add task to student task!")
		return
	}

	studentId, err := strconv.Atoi(c.Param("studentID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.User.AddStudentToTask(uint(studentId), uint(taskId)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
