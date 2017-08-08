package main

import (
	web "github.com/danilojunS/widgets-spa-api/infra/web"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"log"
)

func main() {
	log.Println("Server listening on http://localhost:4000")

	err := web.StartServer()
	utils.CheckError(err)
}
