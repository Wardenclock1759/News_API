package models

import "time"

// swagger:model
type Like struct {
	// uuid for like
	//
	// required: false
	ID string `json:"id"`

	// uuid for article in the system
	//
	// required: true
	ArticleID string `json:"article_id"`

	// uuid for user in the system
	//
	// required: true
	UserID string `json:"user_id"`

	// time when like was created
	//
	// required: false
	LikedAt time.Time `json:"liked_at"`

	// time when like was deleted
	//
	// required: false
	UnlikedAt time.Time `json:"unliked_at"`
}

var likeStorage = NewLikeController()

func NewLikeController() *map[string]Like {
	l1 := Like{
		ID:        "id1",
		ArticleID: "1st",
		UserID:    "id0",
		LikedAt:   time.Now(),
	}

	l2 := Like{
		ID:        "id0",
		ArticleID: "1st",
		UserID:    "id1",
		LikedAt:   time.Now(),
		UnlikedAt: time.Now(),
	}

	res := map[string]Like{}

	res[l1.ID] = l1
	res[l2.ID] = l2

	return &res
}

func AddLike(like Like) (Like, bool) {
	s := *likeStorage

	for _, _like := range s {
		if _like.ArticleID == like.ArticleID && _like.UserID == like.UserID {
			if !_like.UnlikedAt.IsZero() {
				_like.LikedAt = time.Now()
				_like.UnlikedAt = time.Time{}
				s[_like.ID] = _like
				return _like, true
			} else {
				return like, false
			}
		}
	}

	like.ID = GenerateID()
	like.LikedAt = time.Now()
	like.UnlikedAt = time.Time{}
	s[like.ID] = like
	return like, true
}

func GetLike() []Like {
	likes := []Like{}

	for _, like := range *likeStorage {
		if like.UnlikedAt.IsZero() {
			likes = append(likes, like)
		}
	}

	return likes
}

func GetLikeByID(id string) (*Like, bool) {
	storage := *likeStorage

	like, ok := storage[id]
	return &like, ok
}

func DeleteLikeByID(id string) bool {
	storage := *likeStorage

	like, ok := GetLikeByID(id)
	if ok && like.UnlikedAt.IsZero() {
		like.UnlikedAt = time.Now()
		storage[like.ID] = *like
		return true
	}
	return false
}

func GetArticleLikedUsers(id string) *[]User {
	storage := *likeStorage
	users := []User{}

	for _, like := range storage {
		if like.UnlikedAt.IsZero() && like.ArticleID == id {
			user, ok := GetUserByID(like.UserID)
			if ok {
				users = append(users, *user)
			}
		}
	}
	return &users
}
