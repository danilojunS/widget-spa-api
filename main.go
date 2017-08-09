package main

import (
	"fmt"
	config "github.com/danilojunS/widgets-spa-api/config"
	web "github.com/danilojunS/widgets-spa-api/infra/web"
	utils "github.com/danilojunS/widgets-spa-api/utils"
	"log"
)

func main() {
	config.Init()

	port := config.Get().Port

	log.Println(fmt.Sprint("Server listening on http://localhost:", port))
	err := web.StartServer(port)
	utils.CheckError(err)
}
