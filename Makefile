.PHONY: ipchange

ipchange:
	go build -o build/ipchange cmd/ipchange/*.go

clean:
	rm -rf build/
