package main

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Server struct {
  secretKey   string
}

func (s *Server) Init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  s.secretKey = os.Getenv("SECRET_KEY")

  log.Println(s.secretKey)
}
