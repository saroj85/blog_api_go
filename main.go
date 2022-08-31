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

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello home")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestUrl := r.RequestURI
		isToeknVerifyRequired := false

		token := r.Header.Get("Authorization") // this is our token
		fmt.Println("token------------", token, requestUrl)

		if requestUrl == "/user/update" {
			isToeknVerifyRequired = true
		}
		if requestUrl == "/post" {
			isToeknVerifyRequired = true
		}

		if isToeknVerifyRequired {

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
				// newClm := &claim
				// userEmail := newClm.Email
				// userId := *&claim.UserId
				// userName := *&claim.Name
				// userEmail := claim["email"]
				// w.Header.Set("name", "value")
				fmt.Printf("Token is valid: %v", claim)
				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	fmt.Println("App Initialized")

	arrayOpr()
	r := mux.NewRouter()
	routes.RegisterPostRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterCommentRoutes(r)

	r.Use(loggingMiddleware) /// use middleware
	config.Connect()

	log.Fatal(http.ListenAndServe("localhost:8080", r)) // create a new server and listen on port
}

// func updateName(name *string) {
// 	*name = "test1"
// 	fmt.Println("--pointer==", *name)
// }

/**

@ pointer concept
=========started==========

var user string = "saroj"
&user => we will get the user pointer address
*user => we will get the actule value of the pointer

**/

func arrayOpr() {

	// String Array
	// users := []string{"saroj", "saroj85", "saroj777"}
	// fmt.Println("user1", users[1], len(users))

	// users := map

	// fmt.Printf("users: %v\n", users)

}
