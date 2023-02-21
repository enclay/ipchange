package main

import (
	"fmt"
	"ipbash/internal"
)

func getIP() string {

	res, err := internal.DnsGet("dig", "+short", "myip.opendns.com", "@resolver1.opendns.com", "-4")
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

func main() {
	fmt.Println("Your public IPv4 address:")
	fmt.Println(getIP())
}
