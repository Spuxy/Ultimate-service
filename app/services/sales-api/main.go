package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build string = "develop"

func main() {
	_, err := maxprocs.Set()
	if err != nil {
		fmt.Errorf("maxprocs: %w", err)
		os.Exit(1)
	}
	cpus := runtime.GOMAXPROCS(-1)
	log.Printf("Today starting service build[%s] CPU[%d]", build, cpus)
	defer postSetup()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("shutting down service version { ", build, " }")

}

func postSetup() {
	log.Println("ended service version { ", build, " }")
}
