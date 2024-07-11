package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
)

func main() {
 err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  stripe.Key = os.Getenv("SECRET_KEY")
  log.Println(stripe.Key)
  e := echo.New()
  e.GET("/", CreateIntent)
  e.GET("/list", ListIntents)
  e.Logger.Fatal(e.Start(":8080"))
}
