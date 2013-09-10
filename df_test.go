package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDf(t *testing.T) {
	mh := new(MetricHandler)
	col := &Df{
		RawSpace: readFile("fixtures/df.txt"),
		RawInode: readFile("fixtures/df_inode.txt"),
	}
	stats, e := mh.Collect(col)
	if e != nil {
		t.Log("ERROR", e.Error())
		t.Failed()
		return
	}

	assert.Equal(t, len(stats), 8)

	assert.Equal(t, stats[0].Key, "df.space.Total")
	assert.Equal(t, stats[0].Value, int64(20511356))
	assert.Equal(t, stats[0].Tags["file_system"], "/dev/sda")
	assert.Equal(t, stats[0].Tags["mounted_on"], "/")

	assert.Equal(t, stats[4].Key, "df.inode.Total")
	assert.Equal(t, stats[4].Value, int64(1310720))
	assert.Equal(t, stats[4].Tags["file_system"], "/dev/sda")
	assert.Equal(t, stats[4].Tags["mounted_on"], "/")
}
