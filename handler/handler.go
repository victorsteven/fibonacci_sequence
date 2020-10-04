package handler

import (
	"fibonacci_sequence/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)


type service struct {
	fib domain.FibonacciService
}

func NewHandlerService(fib domain.FibonacciService) *service {

	return &service{fib}

}

func (h *service) CurrentNumber(c *gin.Context) {

	current := h.fib.FibonacciGenerator("current")

	c.JSON(http.StatusOK, gin.H{
		"result": current,
	})

	return

}

func (h *service) NextNumber(c *gin.Context) {

	next := h.fib.FibonacciGenerator("next")

	c.JSON(http.StatusOK, gin.H{
		"result": next,
	})

	return

}

func (h *service) PreviousNumber(c *gin.Context) {

	previous := h.fib.FibonacciGenerator("previous")

	if previous == -1 {
		c.JSON(http.StatusOK, gin.H{
			"result": "This cannot go below 0",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": previous,
	})

	return

}


