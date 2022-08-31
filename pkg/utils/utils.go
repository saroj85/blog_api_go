package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic goes here...
		fmt.Println("ENTER MiddlewareOne")
		// next.ServeHTTP(w, r)
	})
}

func CustomMiddleWare(h http.Handler) http.Handler {
	// fmt.Println("middleware middle")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}
func GenerateUniqueID() string {
	uuidValue := uuid.New()
	return uuidValue.String()
}

func HashPassword(password string) (string, error) {

	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err

}

func ValidateHashPassword(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// truncated for brevity
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("../uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("../uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(dst)
	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")
}

type ClaimsStr struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

const JwtSignature = "hmacSampleSecret"

func VerifyToken(token string) (*ClaimsStr, error) {

	mySignature := []byte(JwtSignature)
	// Verify and extract claims from a token:
	// verifiedToken, err := jwt.Parse(token, jwt.SigningMethodHS256)
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match

	claims := &ClaimsStr{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return mySignature, nil
	})

	return claims, err
}
