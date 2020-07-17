// Package classification of News API
//
// Documentation for News API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package controllers

import (
	"News_API/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

// swagger:route GET /article articles list_articles
// Returns list of all articles in the storage
// responses:
//	200: articleResponse
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := r.URL.Query()
	tags := params.Get("tags")
	filter := []string{}
	if tags != "" {
		filter = strings.Split(tags, ",")
	}
	order := params.Get("order")

	articles := models.GetArticles(filter)
	if order != "" {
		if order == "DESC" {
			articles = models.SortByLikesDesc(articles)
		} else if order == "ASC" {
			articles = models.SortByLikesAsc(articles)
		}
	}

	res, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

// swagger:route GET /article/{id} articles get_article_by_id
// Returns article with matching id if any
// responses:
//	200: articleByIDResponse
//	404: noContent
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := mux.Vars(r)["ID"]
	article, ok := models.GetArticleByID(id)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonArticle, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jsonArticle)
}

// swagger:route POST /article articles post_article
// Create article and add it to the storage
// responses:
//	200: articleByIDResponse
//	400: noContent
func AddArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var article models.Article
	err = json.Unmarshal(jsonBody, &article)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if article.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Insuficient data for Article")))
		return
	}

	var ok bool
	article, ok = models.AddArticle(article)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("News article with title '%s' already exists!", article.Title)))
		return
	}

	json.NewEncoder(w).Encode(article)
}
