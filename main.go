package main

import (
	"github/Claudye/autoserver/servers"
	"log"
	"os"
)

func main() {
	if os.Geteuid() != 0 {
		log.Fatal("Please run this programme with sudo access")
	}
	servers.Start()
}
