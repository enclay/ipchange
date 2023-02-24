package main

import (
	"ipchange/internal"
	"log"
	"os"
)

func GetIP() string {

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

const ENV_NAME = "CURRENT_IP"

func Update() {
	log.Println("checking")

	current := os.Getenv(ENV_NAME)
	newip := GetIP()

	if current != newip {
		log.Println("updating to " + newip)
		os.Setenv(ENV_NAME, newip)
		if len(onchange) > 0 {
			res, err := internal.ExecStr(onchange)
			if err == nil {
				log.Println(res)
			}
		}
	}
}
