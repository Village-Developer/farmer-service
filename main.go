package main

import (
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/routers"
)

var db = make(map[string]string)

func main() {
	db := new(configs.Configs)
	db.Connect()
	router := new(routers.IndexRouter)
	r := router.SetupRouter()
	r.Run(":9100")
}
