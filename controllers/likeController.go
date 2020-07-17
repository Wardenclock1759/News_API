package controllers

import (
	"News_API/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// swagger:route POST /like likes post_like
// Create like and add it to the storage
// responses:
//	200: likeByIDResponse
//	400: noContent
func AddLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var like models.Like
	err = json.Unmarshal(jsonBody, &like)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	article, articleExist := models.GetArticleByID(like.ArticleID)
	user, userExist := models.GetUserByID(like.UserID)

	if !articleExist || !userExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Wrong ID for Article or User")))
		return
	}

	var ok bool
	like, ok = models.AddLike(like)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Article '%s' is already liked by '%s'!", article.Title, user.Name)))
		return
	}

	json.NewEncoder(w).Encode(like)
}

// swagger:route GET /like likes list_likes
// Returns list of all likes in the storage
// responses:
//	200: likeResponse
func GetLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	like := models.GetLike()

	res, err := json.Marshal(like)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

// swagger:route GET /like/{id} likes get_like_by_id
// Returns like with matching id if any
// responses:
//	200: likeByIDResponse
//	404: noContent
func GetLikeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["ID"]
	like, ok := models.GetLikeByID(id)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonLike, err := json.Marshal(like)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jsonLike)
}

// swagger:route DELETE /like/{id} likes delete_like_by_id
// Deletes like with matching id if any
// responses:
//	200: noContent
func DeleteLike(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["ID"]
	ok := models.DeleteLikeByID(id)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
