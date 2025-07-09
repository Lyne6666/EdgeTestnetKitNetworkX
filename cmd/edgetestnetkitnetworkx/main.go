// cmd/edgetestnetkitnetworkx/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"edgetestnetkitnetworkx/internal/edgetestnetkitnetworkx"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := edgetestnetkitnetworkx.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
