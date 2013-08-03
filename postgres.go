package main

import (
	_ "github.com/lib/pq"
	"errors"
	"strings"
	"database/sql"
	"time"
	"strconv"
	"fmt"
	"net/url"
)

type PostgresURL struct {
	Host, User, Password, Database string
	Port int
}

func (self *PostgresURL) ConnectString() (s string, e error) {
	if self.Host == "" {
		return "", errors.New("Host must be set")
	}
	if self.Database == "" {
		return "", errors.New("Database must be set")
	}
	if self.User == "" {
		return "", errors.New("User must be set")
	}
	s = fmt.Sprintf("host=%s dbname=%s user=%s", self.Host, self.Database, self.User)
	if self.Password != "" {
		s += " password=" + self.Password
	}
	if self.Port > 0 {
		s += " port=" + strconv.Itoa(self.Port)
	}
	s += " sslmode=disable"
	return
}

func ParsePostgresUrl(raw string) (out *PostgresURL) {
	parsed, e := url.Parse(raw)
	if e != nil {
		return
	}
	out = &PostgresURL{}
	host := parsed.Host
	hostParts := strings.Split(host, ":")
	out.Host = hostParts[0]
	out.Port = 5432
	if len(hostParts) == 2 {
		out.Port, e = strconv.Atoi(hostParts[1])
		if e != nil {
			return
		}
	}
	out.User = parsed.User.Username()
	if pwd, ok := parsed.User.Password(); ok {
		out.Password = pwd
	}
	out.Database = parsed.Path[1:]
	return
}

type PostgreSQLStats struct {
	Uri string
	DB* sql.DB
	*PostgresURL
}

func (self *PostgreSQLStats) Keys() []string {
	return []string {
		"StatActivity", "tables.SeqScan", "tables.SeqTupRead", "tables.IdxScan", "tables.IdxTupFetch", "tables.NTupIns", "tables.NTupUpd", "tables.NTupDel",
		"tables.NTupHotUpd", "tables.NLiveTup", "tables.NDeadTup", "tables.VacuumCount", "tables.AutoVacuumAcount", "tables.AnalyzeCount",
		"tables.AutoAnalyzeCount", "databases.NumBackends", "databases.XactCommit", "databases.XactRollback", "databases.BlksRead", "databases.BlksHit",
		"databases.TupReturned", "databases.TupFetched", "databases.TupInserted", "databases.TupUpdated", "databases.TupDeleted", "databases.Conflicts",
		"databases.TempFiles", "databases.TempBytes", "databases.Deadlocks", "databases.BlkReadTime", "databases.BlkWriteTime", "indexes.IdxScan",
		"indexes.IdxTupRead", "indexes.IdxTupFetch",
	}
}

func (self *PostgreSQLStats) Prefix() string {
	return "postgres"
}

func (self *PostgreSQLStats) Database() string {
	return self.ParsedUrl().Database
}

func (self *PostgreSQLStats) Connect() (e error) {
	s, e := self.ParsedUrl().ConnectString()
	if e != nil {
		return
	}
	self.DB, e = sql.Open("postgres", s)
	return
}

func (self *PostgreSQLStats) Collect(c* MetricsCollection) (e error) {
	if e = self.Connect(); e != nil {
		return
	}

	stats, e := self.StatActivity()
	if e != nil {
		return
	}
	for _, stat := range stats {
		tags := map[string]string {
			"database": stat.Datname,
			"pid": strconv.Itoa(stat.Pid),
			"state": stat.State,
			"waiting": strconv.FormatBool(stat.Waiting),
		}
		c.AddWithTags("StatActivity", 1, tags)
	}

	tableStats, e := self.StatAllTables()
	if e != nil {
		return
	}
	for _, stat := range tableStats {
		tags := map[string]string {
			"table": stat.Relname,
			"database": self.Database(),
		}
		c.AddWithTags("tables.SeqScan", stat.SeqScan, tags)
		c.AddWithTags("tables.SeqTupRead", stat.SeqTupRead, tags)
		c.AddWithTags("tables.IdxScan", stat.IdxScan, tags)
		c.AddWithTags("tables.IdxTupFetch", stat.IdxTupFetch, tags)
		c.AddWithTags("tables.NTupIns", stat.NTupIns, tags)
		c.AddWithTags("tables.NTupUpd", stat.NTupUpd, tags)
		c.AddWithTags("tables.NTupDel", stat.NTupDel, tags)
		c.AddWithTags("tables.NTupHotUpd", stat.NTupHotUpd, tags)
		c.AddWithTags("tables.NLiveTup", stat.NLiveTup, tags)
		c.AddWithTags("tables.NDeadTup", stat.NDeadTup, tags)
		c.AddWithTags("tables.VacuumCount", stat.VacuumCount, tags)
		c.AddWithTags("tables.AutoVacuumAcount", stat.AutoVacuumAcount, tags)
		c.AddWithTags("tables.AnalyzeCount", stat.AnalyzeCount, tags)
		c.AddWithTags("tables.AutoAnalyzeCount", stat.AutoAnalyzeCount, tags)
	}
	dbStats, e := self.StatDatabases()
	if e != nil {
		return
	}
	for _, stat := range dbStats {
		tags := map[string]string {
			"database": stat.Datname,
		}
		c.AddWithTags("databases.NumBackends", stat.NumBackends, tags)
		c.AddWithTags("databases.XactCommit", stat.XactCommit, tags)
		c.AddWithTags("databases.XactRollback", stat.XactRollback, tags)
		c.AddWithTags("databases.BlksRead", stat.BlksRead, tags)
		c.AddWithTags("databases.BlksHit", stat.BlksHit, tags)
		c.AddWithTags("databases.TupReturned", stat.TupReturned, tags)
		c.AddWithTags("databases.TupFetched", stat.TupFetched, tags)
		c.AddWithTags("databases.TupInserted", stat.TupInserted, tags)
		c.AddWithTags("databases.TupUpdated", stat.TupUpdated, tags)
		c.AddWithTags("databases.TupDeleted", stat.TupDeleted, tags)
		c.AddWithTags("databases.Conflicts", stat.Conflicts, tags)
		c.AddWithTags("databases.TempFiles", stat.TempFiles, tags)
		c.AddWithTags("databases.TempBytes", stat.TempBytes, tags)
		c.AddWithTags("databases.Deadlocks", stat.Deadlocks, tags)
		c.AddWithTags("databases.BlkReadTime", stat.BlkReadTime, tags)
		c.AddWithTags("databases.BlkWriteTime", stat.BlkWriteTime, tags)
	}

	indexStats, e := self.IndexStats()
	if e != nil {
		return
	}
	for _, stat := range indexStats {
		tags := map[string]string {
			"database": self.Database(),
			"table": stat.Relname,
			"index": stat.IndexRelname,
		}
		c.AddWithTags("indexes.IdxScan", stat.IdxScan, tags)
		c.AddWithTags("indexes.IdxScan", stat.IdxScan, tags)
		c.AddWithTags("indexes.IdxTupRead", stat.IdxTupRead, tags)
		c.AddWithTags("indexes.IdxTupFetch", stat.IdxTupFetch, tags)
	}
	return
}

func (self *PostgreSQLStats) ParsedUrl() (u *PostgresURL) {
	if self.PostgresURL == nil {
		self.PostgresURL = ParsePostgresUrl(self.Uri)
	}
	return self.PostgresURL
}

type PgTableStat struct {
	Relname string
	SeqScan, SeqTupRead, IdxScan, IdxTupFetch, NTupIns, NTupUpd, NTupDel, NTupHotUpd, NLiveTup, NDeadTup, VacuumCount int64
	AutoVacuumAcount, AnalyzeCount, AutoAnalyzeCount int64
	LastVacuum, LastAutoVacuum, LastAnalyze, LastAutoAnalyze *time.Time
}

type PgStatActivity struct {
	Datname, State string
	Pid int
	BackendStart, XactStart, QueryStart, StateChange time.Time
	Waiting bool
}

func (p* PostgreSQLStats) StatActivity() (out []*PgStatActivity, e error) {
	rows, e := p.DB.Query("SELECT datname, pid, backend_start, xact_start, query_start, state_change, waiting, state FROM pg_stat_activity")
	if e != nil {
		return
	}
	for rows.Next() {
		act := &PgStatActivity{}
		rows.Scan(&act.Datname, &act.Pid, &act.BackendStart, &act.XactStart, &act.QueryStart, &act.StateChange, &act.Waiting, &act.State)
		out = append(out, act)
	}
	return
}

const allTablesSQL = `
	SELECT relname, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch, n_tup_ins, n_tup_upd, n_tup_del, n_tup_hot_upd, n_live_tup, n_dead_tup,
	vacuum_count, autovacuum_count, analyze_count, autoanalyze_count, last_vacuum, last_autovacuum, last_analyze, last_autoanalyze
	FROM pg_stat_all_tables
	WHERE schemaname = 'public'
`


func (p* PostgreSQLStats) StatAllTables() (out []*PgTableStat, e error) {
	rows, e := p.DB.Query(allTablesSQL)
	if e != nil {
		return
	}
	for rows.Next() {
		act := &PgTableStat{}
		e = rows.Scan(
			&act.Relname,
			&act.SeqScan,
			&act.SeqTupRead,
			&act.IdxScan,
			&act.IdxTupFetch,
			&act.NTupIns,
			&act.NTupUpd,
			&act.NTupDel,
			&act.NTupHotUpd,
			&act.NLiveTup,
			&act.NDeadTup,
			&act.VacuumCount,
			&act.AutoVacuumAcount,
			&act.AnalyzeCount,
			&act.AutoAnalyzeCount,
			&act.LastVacuum,
			&act.LastAutoVacuum,
			&act.LastAnalyze,
			&act.LastAutoAnalyze,
		)
		if e != nil {
			return
		}
		out = append(out, act)
	}
	return
}

type PgDatabaseStat struct {
	Datname string
	NumBackends, XactCommit, XactRollback, BlksRead, BlksHit int64
	TupReturned, TupFetched, TupInserted, TupUpdated, TupDeleted int64
	Conflicts, TempFiles, TempBytes, Deadlocks int64
	BlkReadTime, BlkWriteTime  int64
	StatsReset *time.Time
}

const statSQL = `
	SELECT datname, numbackends, xact_commit, xact_rollback, blks_read, blks_hit, tup_returned, tup_fetched, tup_inserted, tup_updated,
		tup_deleted, conflicts, temp_files, temp_bytes, deadlocks, blk_read_time, blk_write_time, stats_reset
	FROM pg_stat_database
	WHERE datname = '%s'
`

func (p* PostgreSQLStats) StatDatabases() (ret []*PgDatabaseStat, e error) {
	rows, e := p.DB.Query(fmt.Sprintf(statSQL, p.Database))
	if e != nil {
		return
	}
	for rows.Next() {
		act := &PgDatabaseStat{}
		e = rows.Scan(
			&act.Datname,
			&act.NumBackends,
			&act.XactCommit,
			&act.XactRollback,
			&act.BlksRead,
			&act.BlksHit,
			&act.TupReturned,
			&act.TupFetched,
			&act.TupInserted,
			&act.TupUpdated,
			&act.TupDeleted,
			&act.Conflicts,
			&act.TempFiles,
			&act.TempBytes,
			&act.Deadlocks,
			&act.BlkReadTime,
			&act.BlkWriteTime,
			&act.StatsReset,
		)
		if e != nil {
			return
		}
		ret = append(ret, act)
	}
	return
}

type PgIndexStat struct {
	Relname, IndexRelname string
	IdxScan, IdxTupRead, IdxTupFetch int64
}

func (p* PostgreSQLStats) IndexStats() (ret []*PgIndexStat, e error) {
	query := "SELECT relname, indexrelname, idx_scan, idx_tup_read, idx_tup_fetch FROM  pg_stat_all_indexes WHERE schemaname = 'public'"
	rows, e := p.DB.Query(query)
	if e != nil {
		return
	}
	for rows.Next() {
		act := &PgIndexStat{}
		e = rows.Scan(
			&act.Relname,
			&act.IndexRelname,
			&act.IdxScan,
			&act.IdxTupRead,
			&act.IdxTupFetch,
		)
		if e != nil {
			return
		}
		ret = append(ret, act)
	}
	return
}

type PgBgWriterStat struct {
	CheckpointsTimes, CheckpointsReq int64
	CheckpointWriteTime, CheckpointSyncTime float64
	BuffersCheckpoint, BuffersClean, MaxwrittenClean int64
	BuffersBackend, BuffersBackendFsync, BuffersAlloc int64
	StatsReset *time.Time
}

const bgWriterSQL = `
	SELECT checkpoints_timed, checkpoints_req, checkpoint_write_time, checkpoint_sync_time, buffers_checkpoint, buffers_clean,
	maxwritten_clean, buffers_backend, buffers_backend_fsync, buffers_alloc, stats_reset
	FROM pg_stat_bgwriter
`

func (p* PostgreSQLStats) BgWriterStats() (ret []*PgBgWriterStat, e error) {
	rows, e := p.DB.Query(bgWriterSQL)
	if e != nil {
		return
	}
	for rows.Next() {
		act := &PgBgWriterStat{}
		e = rows.Scan(
			&act.CheckpointsTimes,
			&act.CheckpointsReq,
			&act.CheckpointWriteTime,
			&act.CheckpointSyncTime,
			&act.BuffersCheckpoint,
			&act.BuffersClean,
			&act.MaxwrittenClean,
			&act.BuffersBackend,
			&act.BuffersBackendFsync,
			&act.BuffersAlloc,
			&act.StatsReset,
		)
		if e != nil {
			return
		}
		ret = append(ret, act)
	}
	return
}
