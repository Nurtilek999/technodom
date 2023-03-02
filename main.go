package main

import (
	"fmt"
	"log"
	"merchant/api"
	"merchant/pkg/config"
	"merchant/pkg/database"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("wwws")

	defer db.Close()

	port := ":8080"
	app := api.SetupRouter(db)
	app.Run(port)

}
