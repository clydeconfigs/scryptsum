build:
	go build scryptsum.go

install:
	go build scryptsum.go
	sudo mv scryptsum /usr/local/bin

uninstall:
	sudo rm /usr/local/bin/scryptsum
