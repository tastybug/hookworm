build:
	go build cmd/cmdline/main.go && mv main hookworm
install:
	sudo mv ./hookworm /usr/local/bin/
