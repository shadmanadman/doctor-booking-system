package middleware

import(
	"context"
	"doctor-booking-system/internal/db"
	"doctor-booking-system/internal/utils"
	"doctor-booking-system/internal/db/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getEnv("JWT_SECRET","supersecret"))

func JWTAuthMiddleware(next http.Handler) http.Handler{
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            next.ServeHTTP(w, r) // Let it pass as unauthenticated (optional)
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        userIDStr, ok := claims["user_id"].(string)
        if !ok {
            http.Error(w, "Missing user_id in token", http.StatusUnauthorized)
            return
        }

        userID, err := strconv.ParseUint(userIDStr, 10, 32)
        if err != nil {
            http.Error(w, "Invalid user ID", http.StatusUnauthorized)
            return
        }

        var user models.User
        if err := db.DB.First(&user, uint(userID)).Error; err != nil {
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }

        ctx := utils.WithUser(r.Context(), &user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}