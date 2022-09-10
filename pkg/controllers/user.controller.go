package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/saroj85/blog_api_go/pkg/config"
	"github.com/saroj85/blog_api_go/pkg/models"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

/**

	This function is used to create a new user and Email should be unique
	{
		Fullname string,
		Email string,
		Phone Int,
		Password string
	}
**/

type UserToken struct {
	Token string `json:"token"`
	User  models.User
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Register User")

	db := config.GetDb()

	// defer config.CloseDb(db)
	var response ResponseStruct

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	var dbUser models.User

	db.Model(models.User{}).Where("email=?", user.Email).First(&dbUser)

	if dbUser.Email != "" {

		// this user already exist
		w.Header().Set("Content-Type", "application/json")
		response.Code = 400
		response.Message = "User already exist.!!"
		response.SentAt = time.Now()
		json.NewEncoder(w).Encode(response)

		return
	}

	user.ID = utils.GenerateUniqueID()

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.HashPassword = hashedPassword

	db.Model(models.User{}).Create(&user)

	token, _ := GeneratJwtToken(user.Email, user.Fullname, user.ID)
	w.Header().Set("Content-Type", "application/json")

	var user_token UserToken
	user_token.Token = token
	user_token.User = models.User{Fullname: user.Fullname, Email: user.Email, Phone: user.Phone}

	w.Header().Set("Content-Type", "application/json")
	response.Code = 200
	response.Message = "Register user successfully"
	response.Data = user_token
	response.SentAt = time.Now()
	json.NewEncoder(w).Encode(&response)

}

/**
	This function is used for login USER and will return THE JWT token
	{
		Email string,
		Password string
	}
**/

func LoginUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login User")

	var response ResponseStruct
	db := config.GetDb()

	// defer config.CloseDb(db)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	var dbUser models.User

	db.Model(models.User{}).Where("email = ?", user.Email).Find(&dbUser)

	if dbUser.Email != "" {

		// user exist  let's check the password

		if utils.ValidateHashPassword(dbUser.HashPassword, user.Password) {

			// everything is correct let's return the jwt for the user
			token, _ := GeneratJwtToken(dbUser.Email, dbUser.Fullname, dbUser.ID)
			w.Header().Set("Content-Type", "application/json")

			var user_token UserToken
			user_token.Token = token
			user_token.User = models.User{Fullname: dbUser.Fullname, Email: dbUser.Email, Phone: dbUser.Phone, AvtarURL: dbUser.Phone}

			response.Message = "Login successful"
			response.Code = 200
			response.SentAt = time.Now().UTC()
			response.Data = user_token

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&response)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		response.Message = "Password incorrect"
		response.Code = 400
		response.SentAt = time.Now().UTC()

		json.NewEncoder(w).Encode(&response)

		return
	}

}

/**
	This function is used to Update the user profile. we will only be able to update these fields
	{
		Fullname: string,
		Mobile: Int,
		Password: string,
		AvtarURL: string,
	}


**/

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update user profile")

	db := config.GetDb()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	FullName := user.Fullname

	db.Model(models.User{}).Where("id = ?", user.ID).Update("fullname", FullName)

	fmt.Println("goo here")

	w.Header().Set("Content-Type", "application/json")

	// var response ResponseStruct
	// response.Code = 200
	// response.Message = "User Updated Successfully"
	// response.SentAt = time.Now()
	// json.NewEncoder(w).Encode(&response)

	// response.Code = 200
	// response.Message = "User Updated Successfully"
	// response.SentAt = time.Now()
	json.NewEncoder(w).Encode("sdd")
	json.NewEncoder(w).Encode("Somthing went wrong")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func GeneratJwtToken(email, name, id string) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	var JwtSignature = "hmacSampleSecret"
	mySignature := []byte(JwtSignature)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"name":    name,
		"user_id": id,
		"nbf":     time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySignature)
	fmt.Println("asd", string(tokenString), err)
	return tokenString, err
}
