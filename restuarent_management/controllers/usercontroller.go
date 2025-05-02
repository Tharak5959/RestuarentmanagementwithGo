package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/yourusername/yourproject/database"
	"github.com/yourusername/yourproject/models"
)
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
func signUp() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func login() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
func HashPassword(password string) (string, error) {
	// Hashing logic here
	return password, nil // Replace with actual hashing logic
}
func VerifyPassword(userPassword string, providedpassword string) (bool, string) {
	// Password verification logic here
	return true, "" // Replace with actual verification logic
}