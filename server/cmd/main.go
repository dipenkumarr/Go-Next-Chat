package main

import (
	"log"

	"github.com/dipenkumarr/go-next-chat/db"
	"github.com/dipenkumarr/go-next-chat/internal/user"
	"github.com/dipenkumarr/go-next-chat/internal/ws"
	"github.com/dipenkumarr/go-next-chat/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not initialize database: ", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
