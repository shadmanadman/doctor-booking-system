package main

import (
    "log"
    "doctor-booking-system/internal/db"
    "github.com/joho/godotenv"
)

func main(){
	// Load .env (only for local development)
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using system environment variables")
    }

    db.Connect()
}