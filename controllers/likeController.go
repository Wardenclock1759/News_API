package controllers

import (
	"News_API/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

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
