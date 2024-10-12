package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func OnShutdown(shutdown func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-c
	id := time.Now().UnixNano()
	log.Println("OnShutdown...", id)
	if shutdown != nil {
		shutdown()
	}
	log.Println("Shutdown done", id)
}
