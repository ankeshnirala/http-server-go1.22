package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/models"
	authservice "github.com/ankeshnirala/http-server-go1.22/internals/infrastructure/AuthorizationService"
)

func main() {

	authService := authservice.NewAuthorizationService()
	principal := models.SecurityPrincipal{ID: "122", Role: authservice.User}
	resource := authservice.Accounts
	action := authservice.Create
	context := map[string]interface{}{"entity": map[string]interface{}{"id": "124", "name": "Account123"}}
	policy, err := authService.AuthorizeOrThrow(principal, resource, action, context)
	if err != nil {
		fmt.Println("Authorization error:", err)
		return
	}
	fmt.Println("Authorization granted for resource:", policy.Resource)

	listenAddr := flag.String("listenAddr", ":4000", "server address port")
	flag.Parse()

	server := rest.NewServer(*listenAddr)

	msg := fmt.Sprintf("started server on [::]:%s, url: http://localhost:%s", *listenAddr, *listenAddr)
	log.Println(msg)

	log.Fatal(server.Start())
}
