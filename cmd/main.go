package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kyonko-moe/kagami/model"
	"github.com/kyonko-moe/kagami/server/v1/http"
)

var (
	servers = make([]model.Server, 0)
)

func init() {
	var (
		http = http.New()
	)
	servers = append(servers, http)
}

func main() {
	log.SetPrefix("[=3=] ")
	log.Println("Starting kagami ...")

	for _, s := range servers {
		if err := s.Start(); err != nil {
			log.Fatalf("%+v", err)
		}
		log.Printf("server started : %s", s)
	}

	log.Println("Started kagami !")
	sig()
}

func sig() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("kagami got signal %s", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Quitting kagami ...")
			for _, s := range servers {
				if err := s.Stop(); err != nil {
					log.Fatalf("%+v", err)
				}
				log.Printf("server stopped : %s", s)
			}
			log.Println("Quitted kagami !")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
