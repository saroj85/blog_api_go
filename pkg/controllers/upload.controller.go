package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/saroj85/blog_api_go/pkg/config"
	"github.com/saroj85/blog_api_go/pkg/models"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

func UpdateUserAvtar(avtar_url string, user_id string) {
	db := config.GetDb()
	db.Model(&models.User{}).Where("id = ?", user_id).Update("avtar_url", avtar_url)

}

func UpdatePostThumbnail(thumbnail_url string, post_id string) {
	db := config.GetDb()
	db.Model(&models.Post{}).Where("id = ?", post_id).Update("thumbnail", thumbnail_url)

}

func UploadImage(w http.ResponseWriter, r *http.Request) {

	update_type := r.PostFormValue("update_type")
	update_id := r.PostFormValue("update_id")

	fmt.Println("update_type", update_type)
	user_id := r.Header.Get("user_id")

	_, _, err := r.FormFile("my_image") //if we don't get the error it means we have file to upload;

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Image Not Found")

	}

	if update_type == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Upload Type not found")
	}

	upload_path, err := utils.UploadFile(r)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if update_type == "user_avtar" {
		// update user avtar here
		UpdateUserAvtar(upload_path, user_id)
	}
	if update_type == "post_thumbnail" {
		// update user avtar here
		UpdatePostThumbnail(upload_path, update_id)
	}

	w.Header().Set("Content-Type", "application/json")
	var response ResponseStruct
	response.Code = 200
	response.Message = "done"
	response.Data = upload_path
	response.SentAt = time.Now()
	json.NewEncoder(w).Encode(&response)

}
