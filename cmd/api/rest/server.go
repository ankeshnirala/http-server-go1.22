package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
	"github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/routers"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {

	// create a new router
	router := httpsetup.NewRouter()

	routers.New(router).RegisterPublicRouter()
	routers.New(router).RegisterPrivateRouter()

	srv := &http.Server{
		Addr:    s.listenAddr,
		Handler: router.Mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("failed while starting the server", err)
		}
	}()

	gracefulShutdown(srv)

	return nil
}

func gracefulShutdown(s *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown...", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	s.Shutdown(tc)
}
