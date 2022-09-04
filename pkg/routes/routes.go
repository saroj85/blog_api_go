package routes

import (
	"github.com/gorilla/mux"
	"github.com/saroj85/blog_api_go/pkg/controllers"
)

// USER ROUTE
var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user/update", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/delete", controllers.DeleteUser).Methods("DELETE")
}

// POST ROUTES
var RegisterPostRoutes = func(router *mux.Router) {

	router.HandleFunc("/posts", controllers.GetAllPost).Methods("GET")
	router.HandleFunc("/post/{PostId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{PostId}", controllers.DeletePost).Methods("DELETE")

}

// COMMENT  ROUTES
var RegisterCommentRoutes = func(router *mux.Router) {
	router.HandleFunc("/post/comments/{PostID}", controllers.GetCommentsByPostID).Methods("GET")
	router.HandleFunc("/post/comment", controllers.AddComment).Methods("POST")
}

var RegisterUploadRoutes = func(router *mux.Router) {
	router.HandleFunc("/upload/image", controllers.UploadImage).Methods("POST")
}
