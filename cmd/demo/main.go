package main

import (
	"going-demo/api"

	"github.com/PirateDreamer/going"
)

func main() {
	router := going.InitService()
	api.InitApi()
	going.GranceRun(router)
}
