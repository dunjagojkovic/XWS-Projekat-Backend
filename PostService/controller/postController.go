package controller

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"postservice/model"
	"postservice/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseId struct {
	Id primitive.ObjectID
}

type User struct {
	Username string
}

type Following struct {
	Users []string
}

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{
		service: service,
	}

}

func renderJSON(w http.ResponseWriter, v interface{}) {

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func decodePostBody(r io.Reader) (*model.Post, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt model.Post
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func decodeCommentBody(r io.Reader) (*model.Comment, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt model.Comment
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func decodeFollowingBody(r io.Reader) (*Following, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt Following
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func decodeLikeBody(r io.Reader) (*User, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt User
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func (pc *PostController) GetAllHandler(w http.ResponseWriter, req *http.Request) {

	posts, _ := pc.service.GetAll()
	renderJSON(w, posts)
}

func (pc *PostController) GetPostCommentsHandler(w http.ResponseWriter, req *http.Request) {

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)
	comments, _ := pc.service.GetPostComments(_id)
	renderJSON(w, comments)
}

func (pc *PostController) GetUserPostsHandler(w http.ResponseWriter, req *http.Request) {

	username, _ := (mux.Vars(req)["username"])
	posts, _ := pc.service.GetUserPosts(username)
	renderJSON(w, posts)
}

func (pc *PostController) CreatePostHandler(w http.ResponseWriter, req *http.Request) {

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodePostBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := pc.service.Insert(rt)
	renderJSON(w, ResponseId{Id: id})
}

func (pc *PostController) CreatePostCommentHandler(w http.ResponseWriter, req *http.Request) {

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeCommentBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)

	idR, _ := pc.service.InsertComment(_id, rt)
	renderJSON(w, ResponseId{Id: idR})
}

func (pc *PostController) CreatePostLikeHandler(w http.ResponseWriter, req *http.Request) {

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeLikeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)

	idI, _ := pc.service.InsertPostLike(_id, rt.Username)
	renderJSON(w, ResponseId{Id: idI})
}

func (pc *PostController) CreatePostDislikeHandler(w http.ResponseWriter, req *http.Request) {

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeLikeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)

	idI, _ := pc.service.InsertPostDislike(_id, rt.Username)
	renderJSON(w, ResponseId{Id: idI})
}

func (pc *PostController) GetPostLikesHandler(w http.ResponseWriter, req *http.Request) {

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)
	likes, _ := pc.service.GetPostLikes(_id)
	renderJSON(w, likes)
}

func (pc *PostController) GetPostDislikesHandler(w http.ResponseWriter, req *http.Request) {

	id, _ := (mux.Vars(req)["id"])
	_id, _ := primitive.ObjectIDFromHex(id)
	dislikes, _ := pc.service.GetPostDislikes(_id)
	renderJSON(w, dislikes)
}

func (pc *PostController) GetFollowingPostsHandler(w http.ResponseWriter, req *http.Request) {

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeFollowingBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, _ := pc.service.GetFollowingPosts(rt.Users)
	renderJSON(w, posts)
}
