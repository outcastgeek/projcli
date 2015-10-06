APP=projcli
VENDOR=$(CURDIR)/vendor
PKG=$(CURDIR)/pkg
BIN=$(CURDIR)/bin
GOPATH := $(CURDIR):$(VENDOR):$(GOPATH)
#go get -u github.com/constabulary/gb/...
# BUILD=go build
BUILD=gb build
DEPS=gb vendor restore
GOFMT=gofmt -w $(CURDIR)
VET=go tool vet -all $(CURDIR)/src
EXE=$(BIN)/$(APP)
#EXE=$(BIN)/$(shell basename `pwd`)
DEPS_LIST=go list -f '{{join .Deps "\n"}}' ogsols

all: fmt vet
	GOOS=linux GOARCH=amd64 $(BUILD)
	GOOS=linux GOARCH=386 $(BUILD)
	GOOS=linux GOARCH=arm $(BUILD)
	GOOS=linux GOARCH=arm64 $(BUILD)
	GOOS=darwin GOARCH=amd64 $(BUILD)
	GOOS=windows GOARCH=amd64 $(BUILD)
	GOOS=windows GOARCH=386 $(BUILD)

# all: restore
# 	$(BUILD)

fmt:
	$(GOFMT)

vet:
	$(VET)

deps_list:
	$(DEPS_LIST)

restore:
	$(DEPS)

run:
	$(EXE)

.PHONY: clean
clean:
	rm -r $(BIN) $(VENDOR)/src $(PKG)

# Some Target Architechtures:
# Example ENV Variables: GOOS=darwin GOARCH=386
# -->      darwin/386: github.com/mitchellh/gox
# -->    darwin/amd64: github.com/mitchellh/gox
# -->       linux/386: github.com/mitchellh/gox
# -->     linux/amd64: github.com/mitchellh/gox
# -->       linux/arm: github.com/mitchellh/gox
# -->     freebsd/386: github.com/mitchellh/gox
# -->   freebsd/amd64: github.com/mitchellh/gox
# -->     openbsd/386: github.com/mitchellh/gox
# -->   openbsd/amd64: github.com/mitchellh/gox
# -->     windows/386: github.com/mitchellh/gox
# -->   windows/amd64: github.com/mitchellh/gox
# -->     freebsd/arm: github.com/mitchellh/gox
# -->      netbsd/386: github.com/mitchellh/gox
# -->    netbsd/amd64: github.com/mitchellh/gox
# -->      netbsd/arm: github.com/mitchellh/gox
# -->       plan9/386: github.com/mitchellh/gox
