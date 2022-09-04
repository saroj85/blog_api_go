package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saroj85/blog_api_go/pkg/config"
	"github.com/saroj85/blog_api_go/pkg/models"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

type ResponseStruct struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	SentAt  time.Time   `json:"sent_at"`
	Data    interface{} `json:"data"`
}

func GetAllPost(w http.ResponseWriter, r *http.Request) {
	db := config.GetDb()
	// defer config.CloseDb(db)

	var post []models.Post
	db.Find(&post)
	w.Header().Set("Content-Type", "application/json")

	var response ResponseStruct
	response.Message = "Get all posts"
	response.Code = 200
	response.SentAt = time.Now()
	response.Data = &post
	json.NewEncoder(w).Encode(&response)

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	db := config.GetDb()

	// defer config.CloseDb(db)

	vars := mux.Vars(r)

	postId := vars["PostId"]

	var post models.Post
	db.Model(&post).Where("id=?", postId).Find(&post)

	// fmt.Println("post", post)
	// if post.ID == "" {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode("ds")
	// }

	var response ResponseStruct
	response.Code = 200
	response.Message = "Get Post By Id Success"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	db := config.GetDb()
	user_id := r.Header.Get("user_id")

	fmt.Println("POST user_id", user_id)
	var post models.Post // Post Object;

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Somthing went wrong")

	}

	post.AuthorId = user_id
	post.ID = utils.GenerateUniqueID()

	db.Create(&post)
	w.Header().Set("Content-Type", "application/json")
	var response ResponseStruct
	response.Code = 200
	response.Message = "done"
	response.Data = post
	response.SentAt = time.Now()
	json.NewEncoder(w).Encode(&response)

	fmt.Println("CREATE POST", post)

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

	fmt.Println("UPDATE POST")

	db := config.GetDb()

	var post models.Post // Post Object;

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	db.Model(models.Post{}).Where("id = ?", post.ID).Updates(&post)

	var response ResponseStruct
	response.Code = 200
	response.Message = "Post Updated successfully"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	db := config.GetDb()

	// defer config.CloseDb(db)

	var post models.Post // Post Object;

	vars := mux.Vars(r)
	postId := vars["PostId"]
	post.ID = postId

	db.Model(models.Post{}).Where("id=?", postId).Delete(&post)

	var response ResponseStruct
	response.Code = 200
	response.Message = "Post Deleted successfully"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}
