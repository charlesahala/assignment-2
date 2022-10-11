package main

import (
	"assignment-2/database"
	"assignment-2/routers"
)

func main() {
	database.StartDB()
	r := router.StartService()
	r.Run(":8080")
}
