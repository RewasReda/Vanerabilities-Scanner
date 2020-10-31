package main

import (
	"fmt"
	l "log"
	"os"

	"github.com/aquasecurity/trivy/internal"

	"github.com/aquasecurity/trivy/pkg/log"
)

var (
	version = "dev"
)

func main() {
	scan("mongo")
}

func scan(Contain string) {
	app := internal.NewApp(version)
	//err := app.Run(os.Args)
	fmt.Println(os.Args)
	theArray := []string{"", "image", "-f", "json", Contain}
	// theArray[0] = "/tmp/go-build538054758/b001/exe/main"
	// theArray[1] = "mongo"
	// theArray[2] = "mongo"

	err := app.Run(theArray)
	if err != nil {
		if log.Logger != nil {
			log.Fatal(err)
		}
		l.Fatal(err)
	}
}
