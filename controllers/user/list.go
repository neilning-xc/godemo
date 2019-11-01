package user

import (
	. "godemo/controllers"
	"godemo/model"
	"godemo/service"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := service.ListUser()
	if err != nil {
		log.Fatal("List user error")
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

	SendJSONResponse(w, results)
}
