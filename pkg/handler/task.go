package handler

import (
	"Proctor/models"
	"Proctor/models/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const INVALID_PARAMETER_ERROR_MSG = "Invalid parameter!"

// @Summary Create a new task
// @Security ApiKeyAuth
// @Tags tasks
// @Description create a new task
// @ID create-task
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Task "Task created successfully"
// @Failure 400 {object} Error "Invalid input data"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/create [post]
func (h *Handler) createTask(c *gin.Context) {
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

	if roleId == 1 {
		newErrorResponse(c, http.StatusForbidden, "Only Teacher or Admin can create tasks!")
		return
	}

	var input DTO.TaskDTO
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var task models.Task
	task.Title = input.Title
	task.Description = input.Description
	task.AccessFrom = input.AccessFrom
	task.AccessTo = input.AccessTo
	task.TeacherID = userId

	taskId, err := h.service.Task.CreateTask(task)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": taskId,
	})
}

// @Summary Get all tasks
// @Security ApiKeyAuth
// @Tags tasks
// @Description get a list of all tasks
// @ID get-all-tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task "List of all tasks"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks [get]
func (h *Handler) getAllTasks(c *gin.Context) {
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
		newErrorResponse(c, http.StatusForbidden, "Only admin can see all tasks!")
		return
	}

	tasks, err := h.service.Task.GetAllTasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary Get a task by ID
// @Security ApiKeyAuth
// @Tags tasks
// @Description get a specific task by its ID
// @ID get-task-by-id
// @Param id path int true "Task ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Task "Detailed information about a task"
// @Failure 404 {object} Error "Task not found"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/{id} [get]
func (h *Handler) getTaskByID(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, INVALID_PARAMETER_ERROR_MSG)
		return
	}

	task, err := h.service.Task.GetTaskByID(uint(taskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Summary Get all tasks for a specific teacher
// @Security ApiKeyAuth
// @Tags tasks
// @Description get all tasks created by a specific teacher
// @ID get-all-teacher-tasks
// @Param id path int true "Teacher ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task "List of tasks created by the teacher"
// @Failure 400 {object} Error "Invalid teacher ID"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/teacher/{id} [get]
func (h *Handler) getAllTeacherTasks(c *gin.Context) {
	teacherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tasks, err := h.service.Task.GetAllTeacherTasks(uint(teacherId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary Get all tasks for a specific student
// @Security ApiKeyAuth
// @Tags tasks
// @Description get all tasks assigned to a specific student
// @ID get-all-student-tasks
// @Param id path int true "Student ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task "List of tasks assigned to the student"
// @Failure 400 {object} Error "Invalid student ID"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/student/{id} [get]
func (h *Handler) getAllStudentTasks(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tasks, err := h.service.GetAllStudentTasks(uint(studentId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// @Summary Update a task
// @Security ApiKeyAuth
// @Tags tasks
// @Description update details of a specific task, accessible only by the task's creator or an admin
// @ID update-task
// @Param id path int true "Task ID"
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task details to update"
// @Success 200 {object} nil "Task updated successfully"
// @Failure 400 {object} Error "Invalid parameter or bad request"
// @Failure 403 {object} Error "Access denied"
// @Failure 404 {object} Error "Task not found"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/update/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
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

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, INVALID_PARAMETER_ERROR_MSG)
		return
	}

	task, err := h.service.Task.GetTaskByID(uint(taskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.TeacherID != userId && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only Teacher created task and Admin can update task!")
		return
	}

	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task.Title, task.Description, task.AccessFrom, task.AccessTo = input.Title, input.Description, input.AccessFrom, input.AccessTo

	if err := h.service.Task.UpdateTask(task); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a task
// @Security ApiKeyAuth
// @Tags tasks
// @Description delete a specific task, accessible only by the task's creator or an admin
// @ID delete-task
// @Param id path int true "Task ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} nil "Task deleted successfully"
// @Failure 400 {object} Error "Invalid task ID"
// @Failure 403 {object} Error "Access denied"
// @Failure 404 {object} Error "Task not found"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/tasks/delete/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
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

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, INVALID_PARAMETER_ERROR_MSG)
		return
	}

	task, err := h.service.Task.GetTaskByID(uint(taskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if task.TeacherID != userId && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only Teacher created task and Admin can delete task!")
		return
	}

	if err = h.service.Task.DeleteTask(uint(taskId)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
