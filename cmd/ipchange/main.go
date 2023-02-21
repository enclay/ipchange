package main

import (
	"flag"
	"ipchange/internal"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func getIP() string {

	res, err := internal.DnsGet("+short", "myip.opendns.com", "@resolver1.opendns.com", "-4")
	if err == nil && internal.IsValidIP(res) {
		return res
	}

	res, err = internal.HttpGet("https://ifconfig.me")
	if err == nil && internal.IsValidIP(res) {
		return res
	}

	res, err = internal.HttpGet("https://api.ipify.org")
	if err == nil && internal.IsValidIP(res) {
		return res
	}

	return ""
}

func execStr(cmdstr string) {
	arr := strings.Fields(cmdstr)
	cmd := exec.Command(arr[0], arr[1:]...)

	stdout, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(string(stdout))
}

func updateIP() {
	current := os.Getenv("CURRENT_IPV4")
	newip := getIP()

	if current != newip {
		os.Setenv("CURRENT_IPV4", newip)
		if len(onchange) > 0 {
			execStr(onchange)
		}
	}
}

var each string
var onchange string

func parseFlags() {
	flag.StringVar(&each, "each", "5m", "How often should IP refreshed")
	flag.StringVar(&onchange, "on-change", "", "IP change hook")
	flag.Parse()
}

func main() {
	parseFlags()

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(each).Do(updateIP)
	if err != nil {
		log.Panic(err.Error())
	}

	s.StartBlocking()
}
