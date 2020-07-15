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
		ID: "1st",
		Tag: Tag{
			ID:   "id1",
			Name: "Sport",
		}}

	art2 := Article{Title: "Sport is bad",
		ID: "2nd",
		Tag: Tag{
			ID:   "id2",
			Name: "Celebrity",
		}}

	art3 := Article{Title: "Lovi yadro x1",
		ID: "3rd",
		Tag: Tag{
			ID:   "id0",
			Name: "None",
		}}
	art4 := Article{Title: "Lovi yadro x2",
		ID: "4th",
		Tag: Tag{
			ID:   "id1",
			Name: "Sport",
		}}
	art5 := Article{Title: "Lovi yadro x3",
		ID: "5th",
		Tag: Tag{
			ID:   "id0",
			Name: "None",
		}}

	res := map[string]Article{}

	res[art1.ID] = art1
	res[art2.ID] = art2
	res[art3.ID] = art3
	res[art4.ID] = art4
	res[art5.ID] = art5

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
