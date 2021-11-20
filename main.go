package main

import (
	"rest/routes"
	"rest/db"
)

func main(){
	db.Init()
	
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}