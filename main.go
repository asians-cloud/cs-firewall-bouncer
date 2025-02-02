package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/asians-cloud/firewall-bouncer/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
