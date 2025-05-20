package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		token, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := fmt.Sprintf("%v", claims["user_id"])
		c.Set("userID", userID)

		c.Next()
	}
}

func WSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Query("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token query parameter missing"})
			c.Abort()
			return
		}
		fmt.Println("Token from query:", tokenString)

		token, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userID := fmt.Sprintf("%v", claims["user_id"])
		c.Set("userID", userID)

		c.Next()
	}
}

func BackendWSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(105)

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
		fmt.Println(105)

		recievedKey := c.Query("key")
		if recievedKey == "" {
			fmt.Println(115)

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token query parameter missing"})
			c.Abort()
			return
		}
		fmt.Println("key from query:", recievedKey)

		key := os.Getenv("BACKEND_WS_KEY")

		if key != recievedKey {
			fmt.Println(128)

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid key"})
			c.Abort()
			return
		}

		c.Set("userID", "-1")

		c.Next()
	}
}
