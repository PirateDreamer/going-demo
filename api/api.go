package api

import "going-demo/internal/rest"

func InitApi() {
	InitIoc()

	rest.InitUserRest()
}

func InitIoc() {
	rest.InitUserIoc()
}
