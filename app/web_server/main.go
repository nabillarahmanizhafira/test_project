package main

import (
	// internal
	"github.com/julienschmidt/httprouter"
	"github.com/nabillarahmanizhafira/test_project/common/configs"
	"github.com/nabillarahmanizhafira/test_project/common/log"
	"github.com/nabillarahmanizhafira/test_project/handler/rest"
	grace "gopkg.in/paytm/grace.v1"
)

func init() {
	// init log
	log.Init()

	// init config variables
	configs.InitConfigVar()
}

func main() {
	// init redis conn
	configs.InitRedisConn()

	// init http router
	router := httprouter.New()
	rest.InitRoutes(router)

	// serve it
	err := grace.Serve(configs.VarsConfig.Server.Port, router)
	if err != nil {
		log.Error(err, "There's an error during starting the server!")
		return
	}
}
