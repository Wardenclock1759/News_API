package main

import "News_API/models"

// swagger:parameters get_article_by_id
type articleIDWrapper struct {
	// uuid of article to get or delete article from storage
	// in: path
	// required: true
	// example: "22e07f21-85cf-4ee6-ba35-330066b046cf"
	ID string `json:"id"`
}

// swagger:parameters list_articles
type articleTagsWrapper struct {
	// string of tag names separated by ',' to filter articles in storage
	// in: query
	// example: "None,Sport,Astronomy"
	Tags string `json:"tags"`
}

// swagger:parameters list_articles
type sortOrderWrapper struct {
	// order "'ASC' or 'DESC' to sort articles by number of likes"
	// in: query
	// example: DESC
	Order string `json:"order"`
}

// swagger:parameters post_article
type postArticleWrapper struct {
	// POST required and optional data about article
	// in: body
	Body models.Article
}

// swagger:parameters get_like_by_id delete_like_by_id
type likeIDWrapper struct {
	// uuid of like to get or delete article-user pair from storage
	// in: path
	// required: true
	// example: "22e07f21-85cf-4ee6-ba35-330066b046cf"
	ID string `json:"id"`
}

// A list of articles in the response
// swagger:response articleResponse
type articleResponseWrapper struct {
	// Articles in the storage
	// in: body
	Body []models.Article
}

// Single article in the response
// swagger:response articleByIDResponse
type articleByIDResponseWrapper struct {
	// Article with matching id
	// in: body
	Body models.Article
}

// A list of likes in the response
// swagger:response likeResponse
type likeResponseWrapper struct {
	// Articles in the storage
	// in: body
	Body []models.Like
}

// A list of tags in the response
// swagger:response tagResponse
type tagResponseWrapper struct {
	// Articles in the storage
	// in: body
	Body []models.Tag
}

// Single like in the response
// swagger:response likeByIDResponse
type likeByIDResponseWrapper struct {
	// Like with matching id
	// in: body
	Body models.Like
}

// A list of users in the response
// swagger:response userResponse
type userResponseWrapper struct {
	// Articles in the storage
	// in: body
	Body []models.User
}

// Single tag in the response
// swagger:response tagByIDResponse
type tagByIDResponseWrapper struct {
	// Tag with matching id
	// in: body
	Body models.Tag
}

// swagger:response noContent
type noContentWrapper struct {
}
