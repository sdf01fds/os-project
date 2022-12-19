package handler

import (
	"github.com/gin-gonic/gin"
	project "github.com/sdf0106/os-project"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input project.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         user.Id,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Email, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
