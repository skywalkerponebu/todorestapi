package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Fatal(c, http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		log.Fatal(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.services.ParseTonen(headerParts[1])
	if err != nil {
		log.Fatal(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}
