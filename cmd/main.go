package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "server address port")
	flag.Parse()

	server := rest.NewServer(*listenAddr)

	msg := fmt.Sprintf("started server on [::]:%s, url: http://localhost:%s", *listenAddr, *listenAddr)
	log.Println(msg)

	log.Fatal(server.Start())
}
