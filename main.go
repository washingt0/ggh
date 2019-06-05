package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/washingt0/ggh/settings"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	gh, err := settings.NewFromFile(usr.HomeDir + "/.ggh")
	if err != nil {
		log.Fatal(err)
	}

	if u, err := gh.Source.Client.GetViewer(gh); err == nil {
		fmt.Printf("%+v\n", u)
	} else {
		log.Fatal(err)
	}
}
