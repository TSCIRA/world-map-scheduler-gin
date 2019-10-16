package main

import (
	"world-map-scheduler-gin/db"
	"world-map-scheduler-gin/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)
}