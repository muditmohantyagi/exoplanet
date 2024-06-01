package main

import (
	"planet.com/model"
	"planet.com/route"
)

func main() {

	db := model.GoConnect()
	db.AutoMigrate(&model.Expplanet{})
	r := route.SetupRouter()
	r.Run()
}
