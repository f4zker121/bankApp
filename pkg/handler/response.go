package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, fields logrus.Fields) {
	logrus.WithFields(fields).Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
