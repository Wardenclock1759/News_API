package main

import (
	"News_API/controllers"
	"News_API/models"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArticle(t *testing.T) {

	tt := []struct {
		name           string
		tagsValue      string
		orderValue     string
		expectedLength int
	}{
		{name: "get all articles without sorting", tagsValue: "", orderValue: "", expectedLength: 5},
		{name: "get all articles sorted ascending", tagsValue: "", orderValue: "ASC", expectedLength: 5},
		{name: "get all articles sorted descending", tagsValue: "", orderValue: "DESC", expectedLength: 5},
		{name: "get all articles with 'None' tags without sorting", tagsValue: "None", orderValue: "", expectedLength: 2},
		{name: "get all articles with 'None' tags sorted ascending", tagsValue: "None", orderValue: "ASC", expectedLength: 2},
		{name: "get all articles with 'None' tags sorted descending", tagsValue: "None", orderValue: "DESC", expectedLength: 2},
		{name: "get all articles with 'None,Sport' tags without sorting", tagsValue: "None,Sport", orderValue: "", expectedLength: 4},
		{name: "get all articles with 'None,Sport' tags sorted ascending", tagsValue: "None,Sport", orderValue: "ASC", expectedLength: 4},
		{name: "get all articles with 'None,Sport' tags sorted descending", tagsValue: "None,Sport", orderValue: "DESC", expectedLength: 4},
		{name: "get all articles with 'None,UnlistedTag' tags without sorting", tagsValue: "None,UnlistedTag", orderValue: "", expectedLength: 2},
		{name: "get all articles with 'None,UnlistedTag' tags sorted ascending", tagsValue: "None,UnlistedTag", orderValue: "ASC", expectedLength: 2},
		{name: "get all articles with 'None,UnlistedTag' tags sorted descending", tagsValue: "None,UnlistedTag", orderValue: "DESC", expectedLength: 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:5000/article?tags="+tc.tagsValue+"&order="+tc.orderValue, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			rec := httptest.NewRecorder()

			controllers.GetArticle(rec, req)

			res := rec.Result()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("coud not read response: %v", err)
			}

			var article []models.Article
			err = json.Unmarshal(body, &article)

			if len(article) != tc.expectedLength {
				t.Fatalf("expected length of %v, but got %v", tc.expectedLength, len(article))
			}

			if tc.orderValue != "" {
				if tc.orderValue == "ASC" {
					if len(article[0].LikedBy) > len(article[1].LikedBy) {
						t.Fatalf("articles are not sorted correctly; expected %v, but got ascending", tc.orderValue)
					}
				} else {
					if len(article[0].LikedBy) < len(article[1].LikedBy) {
						t.Fatalf("articles are not sorted correctly; expected %v, but got descending", tc.orderValue)
					}
				}
			}
		})
	}
}

func TestGetArticleByID(t *testing.T) {
	tt := []struct {
		name               string
		id                 string
		expectedTitle      string
		expectedStatusCode int
	}{
		{name: "get article with empty id", id: "", expectedTitle: "", expectedStatusCode: http.StatusBadRequest},
		{name: "get article with valid id", id: "1st", expectedTitle: "Sport is cool", expectedStatusCode: http.StatusOK},
		{name: "get article with invalid id", id: "oisdfnkjsdnf", expectedTitle: "", expectedStatusCode: http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			str := "localhost:5000/article/" + tc.id
			req, err := http.NewRequest("GET", str, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req = mux.SetURLVars(req, map[string]string{
				"ID": tc.id,
			})
			rec := httptest.NewRecorder()

			controllers.GetArticleByID(rec, req)

			res := rec.Result()

			if res.StatusCode != tc.expectedStatusCode {
				t.Fatalf("expected status code %v; got %v", tc.expectedStatusCode, res.StatusCode)
			}

			if res.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("coud not read response: %v", err)
				}

				var article models.Article
				err = json.Unmarshal(body, &article)

				if article.Title != tc.expectedTitle {
					t.Fatalf("expected title '%s', but got '%s'", tc.expectedTitle, article.Title)
				}
			}
		})
	}
}

func TestAddArticle(t *testing.T) {
	tt := []struct {
		name               string
		article            models.Article
		expectedStatusCode int
	}{
		{name: "add empty article", article: models.Article{}, expectedStatusCode: http.StatusBadRequest},
		{name: "add article with not enough information", article: models.Article{Tag: models.Tag{ID: "id0", Name: "None"}}, expectedStatusCode: http.StatusBadRequest},
		{name: "add empty article", article: models.Article{Title: "testArticle", Tag: models.Tag{ID: "id0", Name: "None"}}, expectedStatusCode: http.StatusOK},
		{name: "add article with existing title", article: models.Article{Title: "Sport is cool", Tag: models.Tag{ID: "id0", Name: "None"}}, expectedStatusCode: http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.article)
			if err != nil {
				t.Fatalf("failed to marshal article '%s' : %v", tc.article.Title, err)
			}
			req, err := http.NewRequest("POST", "localhost:5000/article", bytes.NewReader(b))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			rec := httptest.NewRecorder()

			controllers.AddArticle(rec, req)

			res := rec.Result()

			if res.StatusCode != tc.expectedStatusCode {
				t.Fatalf("expected status code %v; got %v", tc.expectedStatusCode, res.StatusCode)
			}

			if res.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("coud not read response: %v", err)
				}

				var article models.Article
				err = json.Unmarshal(body, &article)

				if article.Title != tc.article.Title {
					t.Fatalf("expected title '%s', but got '%s'", tc.article.Title, article.Title)
				}

				if article.ID == "" {
					t.Fatalf("failed to create article '%s' ID", tc.article.Title)
				}
			}
		})
	}
}

func TestAddLike(t *testing.T) {
	tt := []struct {
		name               string
		articleID          string
		userID             string
		expectedStatusCode int
	}{
		{name: "add empty like", expectedStatusCode: http.StatusBadRequest},
		{name: "add like that exists in the system", articleID: "1st", userID: "id0", expectedStatusCode: http.StatusBadRequest},
		{name: "add like that is 'deleted' in the system", articleID: "1st", userID: "id1", expectedStatusCode: http.StatusOK},
		{name: "add like that does not exist in the system", articleID: "1st", userID: "id2", expectedStatusCode: http.StatusOK},
		{name: "add like with invalid-invalid pair", articleID: "jk;dfngjk", userID: "jasdnfk;", expectedStatusCode: http.StatusBadRequest},
		{name: "add like with invalid-valid pair", articleID: "jasfd", userID: "id2", expectedStatusCode: http.StatusBadRequest},
		{name: "add like with valid-invalid pair", articleID: "2nd", userID: "jsdjnf", expectedStatusCode: http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			like := models.Like{ArticleID: tc.articleID, UserID: tc.userID}
			b, err := json.Marshal(like)
			if err != nil {
				t.Fatalf("failed to marshal like '%s' : %v", like, err)
			}
			req, err := http.NewRequest("POST", "localhost:5000/like", bytes.NewReader(b))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			rec := httptest.NewRecorder()

			controllers.AddLike(rec, req)

			res := rec.Result()

			if res.StatusCode != tc.expectedStatusCode {
				t.Fatalf("expected status code %v; got %v", tc.expectedStatusCode, res.StatusCode)
			}

			if res.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("coud not read response: %v", err)
				}

				var _like models.Like
				err = json.Unmarshal(body, &_like)

				if _like.ID == "" {
					t.Fatalf("failed to generate like_id with article_id: '%s'; user_id: '%s'", _like.ArticleID, _like.UserID)
				}

				if !_like.UnlikedAt.IsZero() {
					t.Fatalf("failed to create like with article_id: '%s'; user_id: '%s'", _like.ArticleID, _like.UserID)
				}
			}
		})
	}
}

func TestDeleteLike(t *testing.T) {
	tt := []struct {
		name               string
		id                 string
		expectedStatusCode int
	}{
		{name: "delete like with empty id", id: "", expectedStatusCode: http.StatusBadRequest},
		{name: "delete like with invalid id", id: "adwqsdfxcv", expectedStatusCode: http.StatusBadRequest},
		{name: "delete like with valid id", id: "id1", expectedStatusCode: http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			str := "localhost:5000/like/" + tc.id
			req, err := http.NewRequest("DELETE", str, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req = mux.SetURLVars(req, map[string]string{
				"ID": tc.id,
			})
			rec := httptest.NewRecorder()

			controllers.DeleteLike(rec, req)

			res := rec.Result()

			if res.StatusCode != tc.expectedStatusCode {
				t.Fatalf("expected status code %v; got %v", tc.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
