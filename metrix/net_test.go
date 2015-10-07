package metrix

import (
	"testing"

	"github.com/dynport/dgtk/tskip/assert"
)

// sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
// sl: slot
// st: socket state
// tx_queue: outgoing data queue in terms of kernel memory usage
// rx_queue: incoming data queue in terms of kernel memory usage
// uid: process id
func TestParseNetLine(t *testing.T) {
	// tcp        0      0 192.168.178.23:38963    193.219.128.49:6697     ESTABLISHED 2932/irssi
	raw := "3: 17B2A8C0:9833 3180DBC1:1A29 01 00000000:00000000 02:0007756C 00000000  1000        0 32802 2 0000000000000000 26 4 30 10 -1"
	nc, err := ParseNetCon(raw)
	if err != nil {
		t.Fatal(err)
	}
	a := assert.New(t)
	a.Equal(nc.Slot, 3)
	a.Equal(nc.LocalIP.String(), "192.168.178.23")
	a.Equal(nc.LocalPort, 38963)
	a.Equal(nc.RemoteIP.String(), "193.219.128.49")
	a.Equal(nc.RemotePort, 6697)
	a.Equal(nc.Status, "ESTABLISHED")
	a.Equal(nc.UUID, "32802")
	a.Equal(nc.Timeout, 2)
}
