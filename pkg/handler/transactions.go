package handler

import (
	"github.com/gin-gonic/gin"
	project "github.com/sdf0106/os-project"
	"net/http"
)

func (h *Handler) createTransaction(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input project.Transaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transaction, err := h.services.CreateTransaction(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          transaction.Id,
		"sender_id":   userId,
		"receiver_id": transaction.ReceiverId,
		"wallet_id":   transaction.WalletId,
		"amount":      transaction.Amount,
		"created_at":  transaction.CreatedAt,
	})
}

func (h *Handler) getAllTransactions(c *gin.Context) {

}

func (h *Handler) getTransactionById(c *gin.Context) {}

func (h *Handler) deleteTransaction(c *gin.Context) {}
