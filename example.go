package main

import (
	"log"

	"github.com/syndicatepro/go.pingdom/pingdom"
)

func main() {
	client := pingdom.NewClient(
		"my appKey from https://my.pingdom.com/account/appkeys",
		"the email address I use to log into Pingdom with",
		"my random-looking password from Pingdom")

	checks, err := client.Checks()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, check := range checks {
		log.Printf("%s is %s\n", check.Hostname, check.Status)
	}
}
