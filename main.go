package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saroj85/blog_api_go/pkg/config"
	"github.com/saroj85/blog_api_go/pkg/controllers"
	"github.com/saroj85/blog_api_go/pkg/routes"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello home")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welocme to the blog Api")

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUrl := r.RequestURI
		isPrivateRoute := false

		token := r.Header.Get("Authorization") // this is our token
		fmt.Println("URL------------", requestUrl)

		if requestUrl == "/user/update" {
			isPrivateRoute = true
		}
		if requestUrl == "/post" {
			isPrivateRoute = true
		}
		if requestUrl == "/upload/image" {
			isPrivateRoute = true
		}

		if isPrivateRoute {

			claim, err := utils.VerifyToken(token)

			if err != nil {

				fmt.Printf("Token is invalid")
				var response controllers.ResponseStruct
				w.Header().Set("Content-Type", "application/json")
				response.Code = 400
				response.Message = "Token is invalid"
				response.SentAt = time.Now()
				json.NewEncoder(w).Encode(&response)
				return

			} else {
				r.Header.Set("user_id", claim.UserId)
				r.Header.Set("user_email", claim.Email)
				r.Header.Set("name", "value")
				fmt.Println("Token is valid:", claim.UserId)
				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	fmt.Println("App Initialized")

	r := mux.NewRouter()

	routes.RegisterPostRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterCommentRoutes(r)
	routes.RegisterUploadRoutes(r)

	r.Use(loggingMiddleware)

	config.Connect()

	r.HandleFunc("/", HomeHandler).Methods("GET")

	port := "8080" //utils.GetDotEnvVariable("PORT")

	serverUrl := ":" + port
	fmt.Println("PORT IS", port)
	log.Fatal(http.ListenAndServe(serverUrl, r)) // create a new server and listen on port
}
