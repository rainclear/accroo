package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"

	"github.com/rainclear/accroo/pkg/config"
	"github.com/rainclear/accroo/pkg/dbm"
	"github.com/rainclear/accroo/pkg/handlers"
	"github.com/rainclear/accroo/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	// change this to true when in production
	app.InProduction = false
	app.DBPath = "db/accroo.db"
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	dbm.NewDbm(&app)

	err := dbm.OpenDb()
	defer dbm.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting application on port: ", portNumber)
	fmt.Println("Account Types: ", len(app.AccountTypes))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
