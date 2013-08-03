package main

func (t *testSuite) TestParsePostgresUrl() {
	raw := "psql://ff:ffpwd@127.0.0.1:1234/ff_test"
	u := ParsePostgresUrl(raw)
	t.Equal(u.Host, "127.0.0.1")
	t.Equal(u.User, "ff")
	t.Equal(u.Password, "ffpwd")
	t.Equal(u.Database, "ff_test")
	t.Equal(u.Port, 1234)

	s, _ := u.ConnectString()
	t.Equal(s, "host=127.0.0.1 dbname=ff_test user=ff password=ffpwd port=1234 sslmode=disable")
}
