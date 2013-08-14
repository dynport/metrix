GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_STATUS = $(shell test -n "`git status --porcelain`" && echo "+CHANGES")
BASE_PACKAGES = constants.go util.go output.go metric.go metrix.go metrics_collection.go metric_handler.go logger.go optparse.go table.go
BASE_COLLECTORS = loadavg.go memory.go cpu.go disk.go processes.go net.go df.go 
EXTRA_COLLECTORS = riak.go elasticsearch.go redis.go postgres.go pgbouncer.go nginx.go
COLLECTORS = $(BASE_COLLECTORS) $(EXTRA_COLLECTORS)
BUILD_CMD = go build -a -ldflags "-X main.GITCOMMIT $(GIT_COMMIT)$(GIT_STATUS)"

default: all

wip: ctags test

ctags:
	go get github.com/jstemmer/gotags
	gotags `find . -name "*.go" | grep -v '/test/' | xargs` 2> /dev/null > tags.tmp && mv tags.tmp tags

install_dependencies:
	go get github.com/remogatto/prettytest
	go get github.com/lib/pq
	go get github.com/stretchr/testify/assert

clean:
	rm -f bin/*

test:
	go test -v $(BASE_PACKAGES) $(COLLECTORS) *_test.go

jenkins: clean install_dependencies test all
	PROC_ROOT=./fixtures ./bin/metrix --loadavg --disk --memory --processes --cpu
	./bin/metrix --loadavg --disk --memory --processes --cpu --keys

all:
	go build -o bin/metrix metrix-cli.go $(BASE_PACKAGES) $(COLLECTORS)
