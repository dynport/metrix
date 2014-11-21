package main

import "os"

type Config struct {
	OpenTSDBUrl string `json:"opentsdb_url"`
}

var OUTPUTS = []string{"graphite", "opentsdb", "amqp","influxdb"}

func processCollector(mh *MetricHandler, output *OutputHandler, c MetricCollector) (e error) {
	all, e := mh.Collect(c)
	if e != nil {
		return
	}
	e = output.WriteMetrics(all)
	return
}

const (
	KEYS = "keys"
	HELP = "help"

	HOSTNAME = "hostname"

	OPENTSDB = "opentsdb"
	GRAPHITE = "graphite"
	AMQP     = "amqp"
	INFLUXDB = "influxdb"
)

func init() {
	parser.Add(HELP, "true", "Print this usage page")
	parser.Add(KEYS, "true", "Only list all known keys")
	parser.AddKey(OPENTSDB, "Report metrics to OpenTSDB host.\nEXAMPLE: opentsdb.host:4242")
	parser.AddKey(AMQP, "Report metrics to AMQP host.\nEXAMPLE: amqp.host:5672")
	parser.AddKey(INFLUXDB, "Report metrics to InfluxDB host.\nEXAMPLE: user:password@influxdb.host:8086/database")
	parser.AddKey(GRAPHITE, "Report metrics to Graphite host.\nEXAMPLE: graphite.host:2003")
	parser.AddKey(HOSTNAME, "Hostname to be used for tagging. If blank the local hostname is used")
}

func main() {
	if e := parser.ProcessAll(os.Args[1:]); e != nil {
		logError(e)
		os.Exit(1)
	}

	output := &OutputHandler{
		OpenTSDBAddress: parser.Get(OPENTSDB),
		GraphiteAddress: parser.Get(GRAPHITE),
		AmqpAddress:     parser.Get(AMQP),
		InfluxDBAddress:     parser.Get(INFLUXDB),
	}
	collectors := map[string]MetricCollector{
		CPU:           &Cpu{},     // migrated
		LOADAVG:       &LoadAvg{}, // migrated
		MEMORY:        &Memory{},  // migrated
		NET:           &Net{},     // migrated
		FREE:          &Free{},    // migrated
		DF:            &Df{},      // migrated
		PROCESSES:     &Processes{},
		FILES:         &Files{},
		ELASTICSEARCH: &ElasticSearch{Url: parser.Get(ELASTICSEARCH)},
		REDIS:         &Redis{Address: parser.Get(REDIS)},
		DISK:          &Disk{},
		RIAK:          &Riak{Address: parser.Get(RIAK)},
		PGBOUNCER:     &PgBouncer{Address: parser.Get(PGBOUNCER)},
		POSTGRES:      &PostgreSQLStats{Uri: parser.Get(POSTGRES)},
		NGINX:         &Nginx{Address: parser.Get(NGINX)},
	}

	mh := &MetricHandler{}

	if parser.Get(HELP) == "true" || len(parser.Values) == 0 {
		parser.PrintDefaults()
		os.Exit(0)
	}

	for key, _ := range parser.Values {
		if collector, ok := collectors[key]; ok {
			if e := processCollector(mh, output, collector); e != nil {
				logger.Print("ERROR: processong collector", key, e.Error())
			}
		} else {
			isOutput := false
			for _, out := range OUTPUTS {
				if out == key {
					isOutput = true
					break
				}
			}
			if !isOutput {
				logger.Printf("ERROR: no collector found for %q", key)
			}
		}
	}
}