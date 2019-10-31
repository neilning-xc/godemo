package user

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world!11111!")
}
