package handler

import (
	"github.com/gin-gonic/gin"
	project "github.com/sdf0106/os-project"
	"net/http"
	"strconv"
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

type getAllTransactionsResponse struct {
	Data []project.Transaction `json:"data"`
}

func (h *Handler) getAllTransactions(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	transactions, err := h.services.GetAllTransactions(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTransactionsResponse{
		Data: transactions,
	})

}

func (h *Handler) getTransactionById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	transactionId, err := strconv.Atoi(c.Param("id"))

	transaction, err := h.services.GetTransactionById(userId, transactionId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) deleteTransaction(c *gin.Context) {
	userId, err := getUserId(c)

	transactionId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	err = h.services.DeleteTransaction(userId, transactionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
