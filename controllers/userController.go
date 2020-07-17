package controllers

import (
	"News_API/models"
	"encoding/json"
	"net/http"
)

// swagger:route GET /user users list_user
// Returns list of all users in the storage
// responses:
//	200: userResponse
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	users := models.GetUsers()

	res, err := json.Marshal(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}
