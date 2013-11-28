GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_STATUS = $(shell test -n "`git status --porcelain`" && echo "+CHANGES")
BUILD_CMD = go build -ldflags "-X main.GITCOMMIT $(GIT_COMMIT)$(GIT_STATUS)"

default:
	go get github.com/dynport/metrix

all: clean test all
	./bin/metrix --cpu --memory --net --df --disk --processes --files

wip: ctags test

ctags:
	go get github.com/jstemmer/gotags
	gotags `find . -name "*.go" | grep -v '/test/' | xargs` 2> /dev/null > tags.tmp && mv tags.tmp tags

install_dependencies:
	go get -d -v ./... && go build -v ./...

clean:
	rm -f bin/*

release:
	GOOS=linux  GOARCH=amd64         bash ./scripts/release.sh
	GOOS=darwin GOARCH=amd64         bash ./scripts/release.sh

test:
	go test -v

jenkins: clean install_dependencies test all
	PROC_ROOT=./fixtures ./bin/metrix --loadavg --disk --memory --processes --cpu
	./bin/metrix --loadavg --disk --memory --processes --cpu --keys --files
