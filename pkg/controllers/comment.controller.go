package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saroj85/blog_api_go/pkg/config"
	"github.com/saroj85/blog_api_go/pkg/models"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {

	connection := config.GetDb()

	postId := mux.Vars(r)["PostID"]

	var comments []models.Comment

	connection.Model(&models.Comment{}).Where("post_id = ?", postId).Find(&comments)

	w.Header().Set("Content-Type", "application/json")

	var response ResponseStruct
	response.Message = "Get All Comments By Post ID"
	response.Code = 200
	response.SentAt = time.Now()
	response.Data = &comments
	json.NewEncoder(w).Encode(&response)

}

func AddComment(w http.ResponseWriter, r *http.Request) {

	connection := config.GetDb()

	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	comment.ID = utils.GenerateUniqueID()
	connection.Create(&comment)
	w.Header().Set("Content-Type", "application/json")
	var response ResponseStruct
	response.Code = 200
	response.Message = "Comment Created successfully"
	response.Data = comment
	response.SentAt = time.Now()
	json.NewEncoder(w).Encode(&response)

}
