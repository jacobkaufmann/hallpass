package main

import (
	"flag"
	"log"

	"github.com/fatih/color"
	"github.com/jacobkaufmann/hallpass/pkg/api"
	"github.com/jacobkaufmann/hallpass/pkg/datastore"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "port for server")
}

func main() {
	flag.Parse()

	datastore.Connect()
	color.Green("Connected to database")

	m := api.Handler()
	log.Fatal(m.Run(":" + port))
}
