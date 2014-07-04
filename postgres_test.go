package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsePostgresUrl(t *testing.T) {
	Convey("PostgresUrl", t, func() {
		raw := "psql://ff:ffpwd@127.0.0.1:1234/ff_test"
		u := ParsePostgresUrl(raw)
		So(u.Host, ShouldEqual, "127.0.0.1")
		So(u.User, ShouldEqual, "ff")
		So(u.Password, ShouldEqual, "ffpwd")
		So(u.Database, ShouldEqual, "ff_test")
		So(u.Port, ShouldEqual, 1234)

		s, _ := u.ConnectString()
		So(s, ShouldEqual, "host=127.0.0.1 dbname=ff_test user=ff password=ffpwd port=1234 sslmode=disable")

	})
}
