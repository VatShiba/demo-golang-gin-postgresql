// package auth

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authHeader := ctx.GetHeader("Authorization")
// 		if authHeader == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
// 			ctx.Abort()
// 			return
// 		}

// 		tokenString := authHeader[len("Bearer "):]
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Validate the signing method
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}

// 			// Return the secret key used for signing the tokens
// 			return []byte("your-secret-key"), nil
// 		})

// 		if err != nil {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			ctx.Abort()
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			// You can access the claims and perform further actions here
// 			ctx.Set("userID", claims["userID"])
// 			ctx.Next()
// 		} else {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			ctx.Abort()
// 			return
// 		}
// 	}
// }
