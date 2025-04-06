package main

import (
	"log"

	"github.com/dipenkumarr/go-next-chat/db"
	"github.com/dipenkumarr/go-next-chat/internal/user"
	"github.com/dipenkumarr/go-next-chat/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not initialize database: ", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")

}
