package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
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

func UploadFile(r *http.Request) (string, error) {

	file, fileHeader, err := r.FormFile("my_image")
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", os.ModePerm)

	if err != nil {
		return "", err
	}

	// real_file_name := fileHeader.Filename;

	file_location := "./uploads/"
	file_name := GenerateUniqueID() + filepath.Ext(fileHeader.Filename)
	final_file_path := file_location + file_name

	fmt.Println("file_name", file_name)
	dst, err := os.Create(final_file_path)
	if err != nil {
		return "", err
	}

	fmt.Println(dst)
	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	return file_name, nil
}

type ClaimsStr struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

const JwtSignature = "hmacSampleSecret"

func VerifyToken(token string) (ClaimsStr, error) {

	mySignature := []byte(JwtSignature)
	claims := &ClaimsStr{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return mySignature, nil
	})

	clm := *claims
	return clm, err
}
