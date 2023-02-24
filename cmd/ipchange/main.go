package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var each string
var onchange string

func main() {
	flag.StringVar(&each, "each", "5m", "How often should IP refreshed")
	flag.StringVar(&onchange, "on-change", "", "IP change handler cmd")
	flag.Parse()

	log.Printf("Current IP: %s\n", GetIP())

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(each).Do(Update)
	if err != nil {
		log.Panic(err.Error())
	}

	s.StartBlocking()
}
