package handler

import (
	"Proctor/models"
	"Proctor/models/DTO"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type WebsocketMessage struct {
	StudentTaskID uint
	Message       string
}

var upgrader = websocket.Upgrader{}

func (h *Handler) webSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Fail to update connection!")
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "Fail to close websocket!")
			return
		}
	}(conn)

	var solution models.StudentSolution
	var startTime, finishTime time.Time

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				return
			}
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var message WebsocketMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		logrus.Infof("Recived a message: %v", message.Message)

		switch message.Message {
		case "start":
			startTime = time.Now()

			solution.TimeStart = &startTime
			solution.StudentTaskID = message.StudentTaskID

			if err = h.service.Redis.Set(c, strconv.Itoa(int(message.StudentTaskID)), solution, 0); err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			break

		case "finish":
			finishTime = time.Now()

			var tempSolution models.StudentSolution
			if err = h.service.Redis.Get(c, strconv.Itoa(int(message.StudentTaskID)), &tempSolution); err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			tempSolution.TimeEnd = &finishTime
			if err = h.service.Redis.Set(c, strconv.Itoa(int(message.StudentTaskID)), tempSolution, 0); err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			break
		}

		if err := conn.WriteMessage(messageType, append(msg, []byte("Nice")...)); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

// @Summary Get all solutions
// @Security ApiKeyAuth
// @Tags solutions
// @Description get a list of all solutions, accessible only by admins
// @ID get-all-solutions
// @Accept  json
// @Produce  json
// @Success 200 {array} models.StudentSolution "List of all solutions"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/solutions [get]
func (h *Handler) getAllSolutions(c *gin.Context) {
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
		newErrorResponse(c, http.StatusForbidden, "Only admin can update solution!")
		return
	}

	solutions, err := h.service.Solution.GetAllSolutions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, solutions)
}

// @Summary Get solutions by student ID
// @Security ApiKeyAuth
// @Tags solutions
// @Description retrieve solutions submitted by a specific student
// @ID get-solutions-by-student-id
// @Param id path int true "Student ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.StudentSolution "List of solutions submitted by the student"
// @Failure 400 {object} Error "Invalid student ID"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/solutions/by-student/{id} [get]
func (h *Handler) getSolutionsByStudentID(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	solutions, err := h.service.Solution.GetSolutionsByStudentID(uint(studentId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, solutions)
}

// @Summary Get student solutions on solved task
// @Security ApiKeyAuth
// @Tags solutions
// @Description retrieve solutions submitted by students on a solved task, accessible only by teachers and admins
// @ID get-student-solutions-on-solved-task
// @Param id path int true "Task ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.StudentSolution "List of solutions submitted by students on the solved task"
// @Failure 400 {object} Error "Invalid task ID"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/solutions/solved-task/{id} [get]
func (h *Handler) getStudentSolutionsOnTask(c *gin.Context) {
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

	if roleId != 2 && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only teachers and admins can see list of students!")
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	solutions, err := h.service.Solution.GetUserSolutionsOnSolvedTask(uint(taskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, solutions)
}

// @Summary Get student solution on task
// @Security ApiKeyAuth
// @Tags solutions
// @Description Get a specific student solution on a given task by student task ID
// @ID get-student-solution-on-task
// @Accept  json
// @Produce  json
// @Param id path int true "Student Task ID"
// @Success 200 {object} models.StudentSolution "Student solution for the specified task"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /get-solution-on-student-task/{id} [get]
func (h *Handler) getStudentSolutionOnTask(c *gin.Context) {
	studentTaskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	solution, err := h.service.Solution.GetStudentSolutionOnTask(uint(studentTaskId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, solution)
}

// @Summary Create a solution
// @Security ApiKeyAuth
// @Tags solutions
// @Description create a new solution for a student task
// @ID create-solution
// @Param id path int true "Student Task ID"
// @Param solutionText body DTO.SolutionDTO true "Solution details"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Solution created successfully with solution ID returned"
// @Failure 400 {object} Error "Bad request due to invalid input"
// @Failure 500 {object} Error "Internal server error or fail to get solution context"
// @Router /api/solutions/on-student-task/{id} [post]
func (h *Handler) createSolution(c *gin.Context) {
	studentTaskId := c.Param("id")

	var solutionText DTO.SolutionDTO
	if err := c.BindJSON(&solutionText); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var solution models.StudentSolution
	if err := h.service.Redis.Get(c, studentTaskId, &solution); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	solutionId, err := h.service.Solution.CreateSolution(solution)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.service.Redis.Delete(c, studentTaskId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"solutionId": solutionId,
	})
}

// @Summary Update the cheating rate of a solution
// @Security ApiKeyAuth
// @Tags solutions
// @Description update the cheating rate for a specific solution, accessible only by admins and teachers
// @ID update-solution-cheating-rate
// @Param id path int true "Solution ID"
// @Param cheatingRate body DTO.SolutionCheatingRateDTO true "Cheating rate data"
// @Accept  json
// @Produce  json
// @Success 200 {object} nil "Cheating rate generated and updated successfully"
// @Failure 400 {object} Error "Invalid parameter or bad request"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/solutions/generate-cheating-rate/{id} [put]
func (h *Handler) generateSolutionCheatingRate(c *gin.Context) {
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

	if roleId != 2 && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only admin can generate solution!")
		return
	}

	solutionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	solution, err := h.service.Solution.GetSolutionByID(uint(solutionId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	solutionMarshaled, err := json.Marshal(solution)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// info := "[{\"total_time\": 6894752,\"compiled_successfully\": 1116,\"add_lines\": 12365,\"cps\": 0.003435512,\"paste_lines\": 5202,\"max_pastes\": 120,\"avg_paste\": 4.553440984}]"
	rate, err := h.service.Solution.GenerateCheatingRate(string(solutionMarshaled))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Solution.UpdateCheatingRate(uint(solutionId), rate); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"CheatingRate:": rate})
}

// @Summary Update the final grade of a solution
// @Security ApiKeyAuth
// @Tags solutions
// @Description update the final grade for a specific solution, accessible only by the task's teacher or admins
// @ID update-solution-final-grade
// @Param id path int true "Solution ID"
// @Param grade body DTO.SolutionFinalGradeDTO true "Final grade data"
// @Accept  json
// @Produce  json
// @Success 200 {object} nil "Final grade updated successfully"
// @Failure 400 {object} Error "Invalid parameter or bad request"
// @Failure 403 {object} Error "Access denied"
// @Failure 500 {object} Error "Internal server error"
// @Router /api/solutions/update-final-grade/{id} [put]
func (h *Handler) updateSolutionFinalGrade(c *gin.Context) {
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

	solutionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	teacherId, err := h.service.Solution.GetTeacherBySolutionID(uint(solutionId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if teacherId != userId && roleId != 3 {
		newErrorResponse(c, http.StatusForbidden, "Only teacher created task and admins can update grade!")
		return
	}

	var grade DTO.SolutionFinalGradeDTO
	if err := c.BindJSON(&grade); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Solution.UpdateFinalGrade(uint(solutionId), grade.FinalGrade); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
