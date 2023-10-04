.PHONY: build, run

build: 
		go build -buildvcs=false main.go
run:
		./main.exe

.DEFAULT_GOAL := build