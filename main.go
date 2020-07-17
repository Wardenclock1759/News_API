package main

import (
	"News_API/controllers"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/article", controllers.GetArticle).Methods("GET")
	router.HandleFunc("/article", controllers.AddArticle).Methods("POST")
	router.HandleFunc("/article/{ID}", controllers.GetArticleByID).Methods("GET")

	router.HandleFunc("/tag", controllers.AddTag).Methods("POST")
	router.HandleFunc("/tag", controllers.GetTag).Methods("GET")
	router.HandleFunc("/tag/{ID}", controllers.GetTagByID).Methods("GET")

	router.HandleFunc("/user", controllers.GetUser).Methods("GET")

	router.HandleFunc("/like", controllers.AddLike).Methods("POST")
	router.HandleFunc("/like", controllers.GetLike).Methods("GET")
	router.HandleFunc("/like/{ID}", controllers.GetLikeByID).Methods("GET")
	router.HandleFunc("/like/{ID}", controllers.DeleteLike).Methods("DELETE")

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Print(err)
	}
}
