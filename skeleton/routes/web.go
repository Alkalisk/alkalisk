package routes

import "github.com/Alkalisk/alkalisk"

func RegisterRoutes() interface{} {
	r := Alkalisk.NewRouter()

	//Set static path
	r.SetStaticPath("/static/")

	return r
}
