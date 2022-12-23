package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	project "github.com/sdf0106/os-project"
	"net/http"
	"strconv"
)

func (h *Handler) createWallet(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var wallet project.Wallet
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&wallet); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "body of the request is invalid")
		return
	}
	walletId, err := h.services.CreateWallet(userId, wallet)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "unable to create wallet")
		return
	}

	wallet.Id = walletId
	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) getAllWallets(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	wallets, err := h.services.GetAllWallets(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "unable to find wallet")
		return
	}

	c.JSON(http.StatusOK, wallets)
}

func (h *Handler) getWalletById(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "wallet id not found")
		return
	}

	wallet, err := h.services.GetWalletById(userID, walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "wallet not found")
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) updateWalletBalance(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var wallet project.UpdateWallet
	walletId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "wallet id not found")
		return
	}

	if err := h.services.UpdateWalletBalance(userID, walletId, wallet.Amount); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "wallet update failed")
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
