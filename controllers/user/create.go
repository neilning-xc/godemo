package user

import (
	"time"
	. "godemo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"godemo/model"
	"godemo/service"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Birthday *time.Time `json:"birthday"`
}

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendJSONResponse(c, http.StatusUnprocessableEntity, "Parameter error")
		return
	}

	user := model.User{Username: r.Username, Password: r.Password, Email: r.Email, Birthday: r.Birthday}

	if service.CreateUser(&user) != nil {
		SendJSONResponse(c, http.StatusOK, "Create fail")
		return
	}

	SendJSONResponse(c, http.StatusOK, user)
}
