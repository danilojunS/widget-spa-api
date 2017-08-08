package main

import (
	web "github.com/danilojunS/widgets-spa-api/infra/web"
	scripts "github.com/danilojunS/widgets-spa-api/scripts"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"log"
)

func main() {
	// create dummy data
	scripts.UserPopulate()
	scripts.WidgetPopulate()

	log.Println("Server listening on http://localhost:4000")

	err := web.StartServer()
	utils.CheckError(err)
}
