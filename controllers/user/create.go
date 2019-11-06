package user

import (
	. "godemo/controllers"
	"godemo/pkg/auth"
	"net/http"
	"time"

	"godemo/model"
	"godemo/service"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   int8   `json:"gender"`
}

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendJSONResponse(c, http.StatusBadRequest, nil)
		return
	}

	birthday, _ := time.Parse("2006-01-02 15:04:05", r.Birthday)
	password, _ := auth.Encrypt(r.Password)
	user := model.User{Username: r.Username, Password: password, Email: r.Email, Birthday: &birthday, Gender: r.Gender}

	if service.CreateUser(&user) != nil {
		SendJSONResponse(c, http.StatusInternalServerError, nil)
		return
	}

	SendJSONResponse(c, http.StatusCreated, user)
}
