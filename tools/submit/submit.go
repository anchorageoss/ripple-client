package main

import (
	"flag"
	"log"
	"os"

	"github.com/anchorageoss/ripple-client/config"
)

var (
	host = flag.String("host", "wss://s2.ripple.com:443", "websockets host")
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	flag.Parse()
	actions, err := config.Parse(os.Stdin)
	checkErr(err)
	checkErr(actions.Prepare())
	checkErr(actions.Submit(*host))
	log.Printf("Submitted %d transactions", actions.Count())
}
