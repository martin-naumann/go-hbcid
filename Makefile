CURRENT_GITHASH=`git rev-parse HEAD`
APP_NAME=go-hbcid
VPATH=.:src:$(GOPATH)/src:

setup: deps
	@echo Checking availability of Go (make sure it's at least Go 1.2)
	go version
	@echo Checking availability of Git
	git --version

update:
	@echo Updating dependencies
	go get github.com/golang/glog
	go get github.com/stretchr/testify
	go get github.com/dchest/uniuri

build: deps
	go build -o build/server src/main.go

package: build
	rm -f build/$(APP_NAME).deb
	fpm -s dir -t deb -n $(APP_NAME) -p build/$(APP_NAME).deb -v $(CURRENT_GITHASH) --description "$(APP_NAME) $(CURRENT_GITHASH)" -x "*.deb" build/

run: deps
	go run src/main.go -stderrthreshold=INFO

test: deps
	go test ./src/...

coverage:
	@echo Coverage not implemented

lint:
	gofmt src/...

loc:
	cloc src/

doc:
	@echo Documentation not implemented

clean:
	rm build/*

deps: github.com/stretchr/testify github.com/golang/glog github.com/dchest/uniuri

github.com/stretchr/testify:
	go get github.com/stretchr/testify

github.com/golang/glog:
	go get github.com/golang/glog

github.com/dchest/uniuri:
	go get github.com/dchest/uniuri
