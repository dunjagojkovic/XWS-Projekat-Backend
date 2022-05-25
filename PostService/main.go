package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"postservice/controller"
	"postservice/repository"
	"postservice/service"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	store := repository.NewPostStore()
	service := service.NewPostService(store)

	postController := controller.NewPostController(service)
	router.HandleFunc("/post/", postController.CreatePostHandler).Methods("POST")
	router.HandleFunc("/post/", postController.GetAllHandler).Methods("GET")
	router.HandleFunc("/post/{id}/comment/", postController.CreatePostCommentHandler).Methods("POST")
	router.HandleFunc("/post/{id}/comments", postController.GetPostCommentsHandler).Methods("GET")
	router.HandleFunc("/post/{id}/like/", postController.CreatePostLikeHandler).Methods("POST")
	router.HandleFunc("/post/{id}/likes", postController.GetPostLikesHandler).Methods("GET")
	router.HandleFunc("/post/{id}/dislike/", postController.CreatePostDislikeHandler).Methods("POST")
	router.HandleFunc("/post/{id}/dislikes", postController.GetPostDislikesHandler).Methods("GET")
	router.HandleFunc("/profile/posts/{username}", postController.GetUserPostsHandler).Methods("GET")
	router.HandleFunc("/following/posts/", postController.GetFollowingPostsHandler).Methods("POST")

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")

}
