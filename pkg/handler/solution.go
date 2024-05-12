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
	"sync"
	"time"
)

type WebsocketMessage struct {
	StudentTaskID uint
	Message       string
}

var upgrader = websocket.Upgrader{}
var Solutions sync.Map

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

			solutionCtx, loaded := Solutions.LoadOrStore(message.StudentTaskID, solution)
			if loaded {
				solutionContext := solutionCtx.(models.StudentSolution)

				solutionContext.TimeStart = &startTime
				solutionContext.StudentTaskID = message.StudentTaskID

				Solutions.Store(message.StudentTaskID, solutionContext)
			}

			break

		case "finish":
			finishTime = time.Now()

			solutionCtx, ok := Solutions.Load(message.StudentTaskID)
			if !ok {
				newErrorResponse(c, http.StatusInternalServerError, "There is no such StudentTaskID in context!")
				continue
			}

			solutionContext := solutionCtx.(models.StudentSolution)
			solutionContext.TimeEnd = &finishTime
			Solutions.Store(message.StudentTaskID, solutionContext)
			break
		}

		if err := conn.WriteMessage(messageType, append(msg, []byte("Nice")...)); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

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

func (h *Handler) createSolution(c *gin.Context) {
	studentTaskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var solutionText DTO.SolutionDTO
	if err := c.BindJSON(&solutionText); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	solutionContext, ok := Solutions.Load(uint(studentTaskId))
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Fail to get solution from context!")
		return
	}

	solution := solutionContext.(models.StudentSolution)
	solution.Solution = solutionText.Solution

	solutionId, err := h.service.Solution.CreateSolution(solution)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Solutions.Delete(strconv.Itoa(studentTaskId))

	c.JSON(http.StatusOK, gin.H{
		"solutionId": solutionId,
	})
}

func (h *Handler) updateSolutionCheatingRate(c *gin.Context) {
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

	var cheatingRate DTO.SolutionCheatingRateDTO
	if err := c.BindJSON(&cheatingRate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	solutionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Solution.UpdateCheatingRate(uint(solutionId), cheatingRate); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

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
