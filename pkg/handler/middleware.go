package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authHeader = "Authorization"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Authorization header is empty!")
		return
	}

	headerParts := strings.Split(header, " ")
	fmt.Println(headerParts)
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header!")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("User id not found")
	}

	idInt, ok := id.(uint)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
