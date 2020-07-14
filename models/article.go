package models

import (
	"sort"
)

type Article struct {
	ID      string `json:"article_id"`
	Title   string `json:"article_title"`
	Tag     Tag    `json:"article_tag"`
	LikedBy []User `json:"liked_by"`
}

var articleStorage = NewArticleController()

func NewArticleController() *map[string]Article {
	art1 := Article{Title: "Sport is cool",
		ID: "First",
		Tag: Tag{
			ID:   "ID1",
			Name: "Sport",
		}}

	art2 := Article{Title: "Sport is bad",
		ID: "Second",
		Tag: Tag{
			ID:   "ID2",
			Name: "Health",
		}}

	res := map[string]Article{}

	res[art1.ID] = art1
	res[art2.ID] = art2

	return &res
}

func AddArticle(article Article) (Article, bool) {
	s := *articleStorage
	for _, _article := range s {
		if _article.Title == article.Title {
			return article, false
		}
	}

	article.ID = GenerateID()
	article.Tag = *GetOrCreateTagByName(article.Tag.Name)
	s[article.ID] = article

	return article, true
}

func GetArticles(filter []string) []Article {
	articles := []Article{}
	storage := *articleStorage
	defer clearData()

	likes := GetLike()
	for _, like := range likes {
		a, articleIsValid := storage[like.ArticleID]
		u, userIsValid := GetUserByID(like.UserID)

		if articleIsValid && userIsValid {
			a.LikedBy = append(a.LikedBy, *u)
			storage[like.ArticleID] = a
		}
	}

	var filterIsEmpty = len(filter) == 0
	for _, article := range storage {
		if !filterIsEmpty {
			if stringInArray(article.Tag.Name, filter) {
				articles = append(articles, article)
			}
		} else {
			articles = append(articles, article)
		}
	}
	return articles
}

func clearData() {
	storage := *articleStorage
	for _, article := range storage {
		article.LikedBy = []User{}
		storage[article.ID] = article
	}
}

func GetArticleByID(id string) (*Article, bool) {
	storage := *articleStorage
	users := GetArticleLikedUsers(id)

	article, ok := storage[id]
	article.LikedBy = *users
	return &article, ok
}

func GetArticlesSortedByLikes() []Article {
	storage := *articleStorage

	articles := make([]Article, len(storage))
	i := 0
	for _, article := range storage {
		articles[i] = article
		i++
	}

	sort.Sort(ByLikesDesc(articles))

	return articles
}

func SortByLikesDesc(articles []Article) []Article {
	sort.Sort(ByLikesDesc(articles))
	return articles
}

func SortByLikesAsc(articles []Article) []Article {
	sort.Sort(ByLikesAsc(articles))
	return articles
}

type ByLikesDesc []Article

func (a ByLikesDesc) Len() int           { return len(a) }
func (a ByLikesDesc) Less(i, j int) bool { return len(a[i].LikedBy) > len(a[j].LikedBy) }
func (a ByLikesDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLikesAsc []Article

func (a ByLikesAsc) Len() int           { return len(a) }
func (a ByLikesAsc) Less(i, j int) bool { return len(a[i].LikedBy) < len(a[j].LikedBy) }
func (a ByLikesAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
