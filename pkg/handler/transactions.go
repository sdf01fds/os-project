package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTransaction(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllTransactions(c *gin.Context) {}

func (h *Handler) getTransactionById(c *gin.Context) {}

func (h *Handler) deleteTransaction(c *gin.Context) {}
