package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func Sum(c *gin.Context) {
	var req SumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := SumResponse{Result: req.A + req.B}
	c.JSON(http.StatusOK, resp)
}
