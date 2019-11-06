package user

import (
	. "godemo/controllers"
	"godemo/pkg/auth"
	"net/http"

	"godemo/model"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if c.Bind(&req) != nil {
		SendJSONResponse(c, http.StatusBadRequest, nil)
		return
	}

	user := model.User{}

	if model.DB.Self.Where(&model.User{Username: req.Username}).First(&user).Error != nil {
		SendJSONResponse(c, http.StatusInternalServerError, nil)
		return
	}

	if auth.ComparePassword(user.Password, req.Password) != nil {
		SendJSONResponse(c, http.StatusNoContent, nil)
	}
	SendJSONResponse(c, http.StatusOK, model.UserInfo{Username: user.Username, Email: user.Email})
}
