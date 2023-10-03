package main

import "GenSecureSign/api"

func main() {
	router := api.InitRoutes()
	router.Run(":8080")

}
