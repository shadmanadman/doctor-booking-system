package utils

import (
    "os"
    "strconv"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getEnv("JWT_SECRET", "supersecret"))

func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "user_id": strconv.FormatUint(uint64(userID), 10),
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
