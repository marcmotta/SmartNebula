// cmd/smartnebula/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartnebula/internal/smartnebula"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartnebula.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
