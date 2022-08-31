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
	connection := config.GetDb()
	// defer config.CloseDb(connection)

	var post []models.Post
	connection.Find(&post)
	w.Header().Set("Content-Type", "application/json")

	var response ResponseStruct
	response.Message = "Get all posts"
	response.Code = 200
	response.SentAt = time.Now()
	response.Data = &post
	json.NewEncoder(w).Encode(&response)

	// fmt.Println("GET ALL POST", post)

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	connection := config.GetDb()

	// defer config.CloseDb(connection)

	vars := mux.Vars(r)
	postId := vars["PostId"]

	var post models.Post
	connection.Model(&post).Where("id=?", postId).Find(&post)

	var response ResponseStruct
	response.Code = 200
	response.Message = "Get Post By Id Success"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	connection := config.GetDb()

	// defer config.CloseDb(connection)

	var post models.Post // Post Object;

	// var result map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&post)

	// fmt.Println(result["post"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	post.ID = utils.GenerateUniqueID()
	connection.Create(&post)
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
	connection := config.GetDb()

	// defer config.CloseDb(connection)

	var post models.Post // Post Object;

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	connection.Model(models.Post{}).Where("id = ?", post.ID).Updates(&post)

	var response ResponseStruct
	response.Code = 200
	response.Message = "Post Updated successfully"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	connection := config.GetDb()

	// defer config.CloseDb(connection)

	var post models.Post // Post Object;

	vars := mux.Vars(r)
	postId := vars["PostId"]
	post.ID = postId

	fmt.Println("DELETE POST", postId)

	connection.Model(models.Post{}).Where("id=?", postId).Delete(&post)

	var response ResponseStruct
	response.Code = 200
	response.Message = "Post Deleted successfully"
	response.SentAt = time.Now()
	response.Data = &post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)

}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("DDD")
	utils.UploadFile(w, r)
}
