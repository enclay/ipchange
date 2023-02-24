.PHONY: ipchange

ipchange:
	go build -o build/ipchange cmd/ipchange/main.go

clean:
	rm -rf build/
