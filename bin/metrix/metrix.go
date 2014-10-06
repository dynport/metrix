package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/dynport/metrix/metrix"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	if e := run(); e != nil {
		logger.Fatal(e)
	}
}

var mapping = map[string]func() error{
	"snapshot":   snapshot,
	"disks":      disks,
	"memory":     memory,
	"stat":       stat,
	"load":       load,
	"proc-stats": procStats,
}

func benchmark(message string) func() {
	started := time.Now()
	logger.Printf("started  %s", message)
	return func() {
		logger.Printf("finished %s in %.06f", message, time.Since(started).Seconds())
	}
}

func snapshot() error {
	defer benchmark("create snapshot")()
	s := &metrix.Snapshot{}
	err := s.Load()
	if err != nil {
		return err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	localPath := "/data/metrix/snapshots/" + time.Now().UTC().Format("2006/01/02/15/2006-01-02T150405") + "-" + hostname + ".json.gz"
	logger.Printf("writing snapshot to %s", localPath)
	tmpPath := localPath + ".tmp"
	err = os.MkdirAll(path.Dir(tmpPath), 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	defer f.Close()
	gz := gzip.NewWriter(f)
	defer gz.Close()
	enc := json.NewEncoder(gz)
	err = enc.Encode(s)
	if err != nil {
		return err
	}
	return os.Rename(tmpPath, localPath)
}

func procStats() error {
	stats, e := metrix.LoadProcStats()
	if e != nil {
		return e
	}
	for _, s := range stats {
		logger.Printf("%d: %s utime=%d stime=%d, rss=%d", s.Pid, s.Comm, s.Utime, s.Stime, s.RSS)
	}
	return nil
}

func load() error {
	l, e := metrix.LoadLoadAvg()
	if e != nil {
		return e
	}
	logger.Printf("1min: %.02f", l.Min1)
	logger.Printf("5min: %.02f", l.Min5)
	logger.Printf("15min: %.02f", l.Min15)
	return nil

}

func availableNames() []string {
	names := []string{}
	for n := range mapping {
		names = append(names, n)
	}
	return names
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("argument expected. available are %#v", availableNames())
	}
	selected := []func() error{}
	for n, f := range mapping {
		if strings.HasPrefix(n, os.Args[1]) {
			selected = append(selected, f)
		}
	}
	if len(selected) == 1 {
		return selected[0]()
	}
	return fmt.Errorf("no handler found, knonw are %s", strings.Join(availableNames(), ", "))
}

func memory() error {
	mem, e := metrix.LoadMeminfo()
	if e != nil {
		return e
	}
	logger.Printf("total: %d", mem.MemTotal)
	logger.Printf("free: %d", mem.MemFree)
	return nil
}

func disks() error {
	disks, e := metrix.LoadDisks()
	if e != nil {
		return e
	}
	for _, d := range disks {
		logger.Printf("%s: %d", d.Filesystem, d.Use)
	}
	return nil

}

func stat() error {
	stat, e := metrix.LoadStat()
	if e != nil {
		return e
	}
	logger.Printf("user=%d iowait=%d", stat.Cpu.User, stat.Cpu.IOWait)
	return nil
}
