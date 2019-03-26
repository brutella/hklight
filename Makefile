GOCMD=go
GORUN=$(GOCMD) run

export GO111MODULE=on

run:
	$(GORUN) hklightd.go