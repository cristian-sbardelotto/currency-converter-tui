package services

import (
	"log"

	"github.com/joho/godotenv"
)

func CalculateCurrency() {
	panic("not implemented")
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}
