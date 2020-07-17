package controllers

import (
	"News_API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// swagger:route GET /tag tags list_tag
// Returns list of all tags in the storage
// responses:
//	200: tagResponse
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

// swagger:route GET /tag/{id} tags get_tag_by_id
// Returns tag with matching id if any
// responses:
//	200: tagByIDResponse
//	404: noContent
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

// swagger:route POST /tag tags post_tag
// Create tag and add it to the storage
// responses:
//	200: tagByIDResponse
//	400: noContent
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
