package paramid

import (
	"net/http"
	"router"
	"utils"
	"encoding/json"
	"github.com/gorilla/context"
)

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := router.Params(r)
		ID := params.ByName("id")
		id := utils.StringToUInt(ID)
		if (id == 0) {
			json.NewEncoder(w).Encode("invalid id")
			return
		}
		context.Set(r, "id", id)
		next.ServeHTTP(w, r)
	})
}
