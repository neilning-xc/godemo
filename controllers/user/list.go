package user

import (
	. "godemo/controllers"
	"godemo/model"
	"godemo/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
	users, err := service.ListUser()
	if err != nil {
		log.Fatalf(err, "List user error")
	}

	results := make([]*model.UserInfo, 0)
	for _, user := range users {
		results = append(results, &model.UserInfo{
			Id:        uint64(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			Gender:    int8(user.Gender),
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.CreatedAt.String(),
		})
	}

	SendJSONResponse(c, http.StatusOK, results)
}
