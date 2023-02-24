package main

import (
	"flag"
	"fmt"
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

	res, err = internal.DnsGet("+short", "@ns1-1.akamaitech.net", "ANY", "whoami.akamai.net")
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

const ENV_NAME = "CURRENT_IP"

func updateIP() {

	fmt.Println("checking...")

	current := os.Getenv(ENV_NAME)
	newip := getIP()

	if current != newip {
		fmt.Println("updating to " + newip)
		os.Setenv(ENV_NAME, newip)
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

	fmt.Println(getIP())

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(each).Do(updateIP)
	if err != nil {
		log.Panic(err.Error())
	}

	s.StartBlocking()
}
