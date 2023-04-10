package main

import (
	"flag"
	"github.com/erry-az/test-go/internal/app"
	"log"
)

func main() {
	dsn := flag.String("dsn", "", "db postgresql dsn")
	flag.Parse()

	err := app.Rest(*dsn)
	if err != nil {
		log.Fatal(err)
	}
}
