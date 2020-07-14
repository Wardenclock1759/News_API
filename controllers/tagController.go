package controllers

import (
	"News_API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	tags := models.GetTags()

	res, err := json.Marshal(tags)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

func GetTagByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["ID"]
	tag, ok := models.GetTagByID(id)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonTag, err := json.Marshal(tag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jsonTag)
}

func AddTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var tag models.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	models.AddTag(tag)
	json.NewEncoder(w).Encode(tag)
}
