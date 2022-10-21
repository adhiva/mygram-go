package main

import (
	"log"

	"golang-mygram/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Hacktiv8 Final Project - MyGram API
// @version 1.0
// @description This is API created, requirements for completing training Hacktiv8
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
