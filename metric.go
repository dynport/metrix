package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Metric struct {
	Key   string
	Value int64
	Tags  map[string]string `json:"Tags,omitempty"`
}

func parseInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func parseInt64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 64)
	return
}

var pgbouncerMetricMapping = map[string]string{
	"sockets.RecvPos":    "sockets.RecvPos",
	"sockets.PktPos":     "sockets.PktPos",
	"sockets.PktRemain":  "sockets.PktRemain",
	"sockets.SendPos":    "sockets.SendPos",
	"sockets.SendRemain": "sockets.SendRemain",
	"sockets.PktAvail":   "sockets.PktAvail",
	"sockets.SendAvail":  "sockets.SendAvail",

	// stats
	"stats.TotalRequests":  "stats.TotalRequests",
	"stats.TotalReceived":  "stats.TotalReceived",
	"stats.TotalSent":      "stats.TotalSent",
	"stats.TotalQueryTime": "stats.TotalQueryTime",
	"stats.AvgReq":         "stats.AvgReq",
	"stats.AvgRecv":        "stats.AvgRecv",
	"stats.AvgSent":        "stats.AvgSent",
	"stats.AvgQuery":       "stats.AvgQuery",

	// memory
	"connections.Total": "connections.Total",
	"memory.Size":       "memory.Size",
	"memory.Used":       "memory.Used",
	"memory.Free":       "memory.Free",
	"memory.MemTotal":   "memory.MemTotal",

	// pools
	"pools.ClientsActive":  "pools.ClientsActive",
	"pools.ClientsWaiting": "pools.ClientsWaiting",
	"pools.ServersActive":  "pools.ServersActive",
	"pools.ServersIdle":    "pools.ServersIdle",
	"pools.ServersUsed":    "pools.ServersUsed",
	"pools.ServersTested":  "pools.ServersTested",
	"pools.ServersLogin":   "pools.ServersLogin",
	"pools.MaxWait":        "pools.MaxWait",
}

var postgresMetricMapping = map[string]string{
	"StatActivity":            "StatActivity",
	"tables.SeqScan":          "tables.SeqScan",
	"tables.SeqTupRead":       "tables.SeqTupRead",
	"tables.IdxScan":          "tables.IdxScan",
	"tables.IdxTupFetch":      "tables.IdxTupFetch",
	"tables.NTupIns":          "tables.NTupIns",
	"tables.NTupUpd":          "tables.NTupUpd",
	"tables.NTupDel":          "tables.NTupDel",
	"tables.NTupHotUpd":       "tables.NTupHotUpd",
	"tables.NLiveTup":         "tables.NLiveTup",
	"tables.NDeadTup":         "tables.NDeadTup",
	"tables.VacuumCount":      "tables.VacuumCount",
	"tables.AutoVacuumAcount": "tables.AutoVacuumAcount",
	"tables.AnalyzeCount":     "tables.AnalyzeCount",
	"tables.AutoAnalyzeCount": "tables.AutoAnalyzeCount",
	"databases.NumBackends":   "databases.NumBackends",
	"databases.XactCommit":    "databases.XactCommit",
	"databases.XactRollback":  "databases.XactRollback",
	"databases.BlksRead":      "databases.BlksRead",
	"databases.BlksHit":       "databases.BlksHit",
	"databases.TupReturned":   "databases.TupReturned",
	"databases.TupFetched":    "databases.TupFetched",
	"databases.TupInserted":   "databases.TupInserted",
	"databases.TupUpdated":    "databases.TupUpdated",
	"databases.TupDeleted":    "databases.TupDeleted",
	"databases.Conflicts":     "databases.Conflicts",
	"databases.TempFiles":     "databases.TempFiles",
	"databases.TempBytes":     "databases.TempBytes",
	"databases.Deadlocks":     "databases.Deadlocks",
	"databases.BlkReadTime":   "databases.BlkReadTime",
	"databases.BlkWriteTime":  "databases.BlkWriteTime",
	"indexes.IdxScan":         "indexes.IdxScan",
	"indexes.IdxTupRead":      "indexes.IdxTupRead",
	"indexes.IdxTupFetch":     "indexes.IdxTupFetch",
}

var metricMapping = map[string]map[string]string{
	"pgbouncer": pgbouncerMetricMapping,
	"postgres":  postgresMetricMapping,
	"riak": map[string]string{
		"VNodeGets":                         "VNodeGets",
		"VNodeGetsTotal":                    "VNodeGetsTotal",
		"VNodePuts":                         "VNodePuts",
		"VNodePutsTotal":                    "VNodePutsTotal",
		"VNodeIndexReads":                   "VNodeIndexReads",
		"VNodeIndexReadsTotal":              "VNodeIndexReadsTotal",
		"VNodeIndexWrites":                  "VNodeIndexWrites",
		"VNodeIndexWritesTotal":             "VNodeIndexWritesTotal",
		"VNodeIndexWritesPostings":          "VNodeIndexWritesPostings",
		"VNodeIndexWritesPostings_total":    "VNodeIndexWritesPostings_total",
		"VNodeIndexDeletes":                 "VNodeIndexDeletes",
		"VNodeIndexDeletes_total":           "VNodeIndexDeletes_total",
		"VNodeIndexDeletes_postings":        "VNodeIndexDeletes_postings",
		"VNodeIndexDeletes_postings_total":  "VNodeIndexDeletes_postings_total",
		"NodeGets":                          "NodeGets",
		"NodeGetsTotal":                     "NodeGetsTotal",
		"NodeGet_fsm_siblings_mean":         "NodeGet_fsm_siblings_mean",
		"NodegetFsmSiblingsMedian":          "NodegetFsmSiblingsMedian",
		"NodegetFsmSiblings95":              "NodegetFsmSiblings95",
		"NodegetFsmSiblings99":              "NodegetFsmSiblings99",
		"NodegetFsmSiblings100":             "NodegetFsmSiblings100",
		"NodegetFsmObjsizeMean":             "NodegetFsmObjsizeMean",
		"NodegetFsmObjsizeMedian":           "NodegetFsmObjsizeMedian",
		"NodegetFsmObjsize95":               "NodegetFsmObjsize95",
		"NodegetFsmObjsize99":               "NodegetFsmObjsize99",
		"NodegetFsmObjsize100":              "NodegetFsmObjsize100",
		"NodegetFsmTimeMean":                "NodegetFsmTimeMean",
		"NodegetFsmTimeMedian":              "NodegetFsmTimeMedian",
		"NodegetFsmTime95":                  "NodegetFsmTime95",
		"NodegetFsmTime99":                  "NodegetFsmTime99",
		"NodegetFsmTime100":                 "NodegetFsmTime100",
		"NodePuts":                          "NodePuts",
		"NodePutsTotal":                     "NodePutsTotal",
		"NodePutFsmTimeMean":                "NodePutFsmTimeMean",
		"NodePutFsmTimeMedian":              "NodePutFsmTimeMedian",
		"Node_put_fsm_time_95":              "Node_put_fsm_time_95",
		"Node_put_fsm_time_99":              "Node_put_fsm_time_99",
		"Node_put_fsm_time_100":             "Node_put_fsm_time_100",
		"ReadRepairs":                       "ReadRepairs",
		"ReadRepairsTotal":                  "ReadRepairsTotal",
		"CoordRedirsTotal":                  "CoordRedirsTotal",
		"ExecutingMappers":                  "ExecutingMappers",
		"PrecommitFail":                     "PrecommitFail",
		"PostcommitFail":                    "PostcommitFail",
		"PbcActive":                         "PbcActive",
		"PbcConnects":                       "PbcConnects",
		"PbcConnectsTotal":                  "PbcConnectsTotal",
		"NodeGetFsmActive":                  "NodeGetFsmActive",
		"NodeGetFsmActive60s":               "NodeGetFsmActive60s",
		"NodeGetFsmInRate":                  "NodeGetFsmInRate",
		"NodeGetFsmOutRate":                 "NodeGetFsmOutRate",
		"NodeGetFsmRejected":                "NodeGetFsmRejected",
		"NodeGetFsmRejected60s":             "NodeGetFsmRejected60s",
		"NodeGetFsmRejectedTotal":           "NodeGetFsmRejectedTotal",
		"NodePutFsmActive":                  "NodePutFsmActive",
		"NodePutFsmActive60s":               "NodePutFsmActive60s",
		"NodePutFsmInRate":                  "NodePutFsmInRate",
		"NodePutFsmOutRate":                 "NodePutFsmOutRate",
		"NodePutFsmRejected":                "NodePutFsmRejected",
		"NodePutFsmRejected60s":             "NodePutFsmRejected60s",
		"NodePutFsmRejectedTotal":           "NodePutFsmRejectedTotal",
		"ReadRepairsPrimaryOutofdateOne":    "ReadRepairsPrimaryOutofdateOne",
		"ReadRepairsPrimaryOutofdateCount":  "ReadRepairsPrimaryOutofdateCount",
		"ReadRepairsPrimaryNotfoundOne":     "ReadRepairsPrimaryNotfoundOne",
		"ReadRepairsPrimaryNotfoundCount":   "ReadRepairsPrimaryNotfoundCount",
		"ReadRepairsFallbackOutofdateOne":   "ReadRepairsFallbackOutofdateOne",
		"ReadRepairsFallbackOutofdateCount": "ReadRepairsFallbackOutofdateCount",
		"ReadRepairsFallbackNotfoundOne":    "ReadRepairsFallbackNotfoundOne",
		"ReadRepairsFallbackNotfoundCount":  "ReadRepairsFallbackNotfoundCount",
		"PipelineActive":                    "PipelineActive",
		"PipelineCreateCount":               "PipelineCreateCount",
		"PipelineCreateOne":                 "PipelineCreateOne",
		"PipelineCreateErrorCount":          "PipelineCreateErrorCount",
		"PipelineCreateErrorOne":            "PipelineCreateErrorOne",
		"CpuNprocs":                         "CpuNprocs",
		"CpuAvg1":                           "CpuAvg1",
		"CpuAvg5":                           "CpuAvg5",
		"CpuAvg15":                          "CpuAvg15",
		"MemTotal":                          "MemTotal",
		"MemAllocated":                      "MemAllocated",
		"SysGlobalHeapsSize":                "SysGlobalHeapsSize",
		"SysProcessCount":                   "SysProcessCount",
		"SysThreadPoolSize":                 "SysThreadPoolSize",
		"SysWordsize":                       "SysWordsize",
		"RingNumPartitions":                 "RingNumPartitions",
		"RingCreationSize":                  "RingCreationSize",
		"MemoryTotal":                       "MemoryTotal",
		"MemoryProcesses":                   "MemoryProcesses",
		"MemoryProcessesUsed":               "MemoryProcessesUsed",
		"MemorySystem":                      "MemorySystem",
		"MemoryAtom":                        "MemoryAtom",
		"MemoryAtomUsed":                    "MemoryAtomUsed",
		"MemoryBinary":                      "MemoryBinary",
		"MemoryCode":                        "MemoryCode",
		"MemoryEts":                         "MemoryEts",
		"RiakCoreStatTs":                    "RiakCoreStatTs",
		"IgnoredGossipTotal":                "IgnoredGossipTotal",
		"RingsReconciledTotal":              "RingsReconciledTotal",
		"RingsReconciled":                   "RingsReconciled",
		"GossipReceived":                    "GossipReceived",
		"RejectedHandoffs":                  "RejectedHandoffs",
		"HandoffTimeouts":                   "HandoffTimeouts",
		"DroppedVnodeRequestsTotal":         "DroppedVnodeRequestsTotal",
		"ConvergeDelayMin":                  "ConvergeDelayMin",
		"ConvergeDelayMax":                  "ConvergeDelayMax",
		"ConvergeDelayMean":                 "ConvergeDelayMean",
		"ConvergeDelayLast":                 "ConvergeDelayLast",
		"RebalanceDelayMin":                 "RebalanceDelayMin",
		"RebalanceDelayMax":                 "RebalanceDelayMax",
		"RebalanceDelayMean":                "RebalanceDelayMean",
		"RebalanceDelayLast":                "RebalanceDelayLast",
		"RiakKvVnodesRunning":               "RiakKvVnodesRunning",
		"RiakKvVnodeqMin":                   "RiakKvVnodeqMin",
		"RiakKvVnodeqMedian":                "RiakKvVnodeqMedian",
		"RiakKvVnodeqMean":                  "RiakKvVnodeqMean",
		"RiakKvVnodeqMax":                   "RiakKvVnodeqMax",
		"RiakKvVnodeqTotal":                 "RiakKvVnodeqTotal",
		"RiakPipeVnodesRunning":             "RiakPipeVnodesRunning",
		"RiakPipeVnodeqMin":                 "RiakPipeVnodeqMin",
		"RiakPipeVnodeqMedian":              "RiakPipeVnodeqMedian",
		"RiakPipeVnodeqMean":                "RiakPipeVnodeqMean",
		"RiakPipeVnodeqMax":                 "RiakPipeVnodeqMax",
		"RiakPipeVnodeqTotal":               "RiakPipeVnodeqTotal",
		"ConnectedNodesCount":               "ConnectedNodesCount",
		"RingMembersCount":                  "RingMembersCount",
	},
}

func AllMetricKeys() (ret []string) {
	ret = make([]string, 0)
	for prefix, mapping := range metricMapping {
		for k, _ := range mapping {
			ret = append(ret, prefix+"."+k)
		}
	}
	sort.Strings(ret)
	return
}

func MakeMetric(t, internalKey string, v int64, tags map[string]string) (m *Metric) {
	mapping, ok := metricMapping[t]
	if !ok {
		panic("key " + t + " can not defined in metric mapping")
	}
	key, ok := mapping[internalKey]

	if !ok {
		panic("key " + internalKey + " not defined")
	}
	if tags == nil {
		tags = map[string]string{}
	}
	m = &Metric{Key: t + "." + key, Value: v, Tags: tags}
	return
}

func (m *Metric) Ascii(t time.Time, hostname string) (r string) {
	r = fmt.Sprintf("%s %d %d", m.Key, t.Unix(), m.Value)
	m.Tags["host"] = hostname
	for k, v := range m.Tags {
		if len(v) > 0 {
			r = r + " " + k + "=" + v
		}
	}
	return
}

func (m *Metric) NormalizeTag(s string) (r string) {
	re := regexp.MustCompile("(^\\()|(\\)$)")
	re2 := regexp.MustCompile("[^\\w]+")
	return re2.ReplaceAllString(re.ReplaceAllString(s, ""), "_")
}

func (m *Metric) OpenTSDB(t time.Time, hostname string) (r string) {
	r = fmt.Sprintf("put %s %d %d", m.Key, t.Unix(), m.Value)
	if hostname == "" {
		panic("hostname must not be nil")
	}
	r = r + " host=" + hostname
	for k, v := range m.Tags {
		if len(v) > 0 {
			r = r + " " + k + "=" + m.NormalizeTag(v)
		}
	}
	return
}

func (m *Metric) Graphite(t time.Time, hostname string) (r string) {
	key := m.Key
	if cpu_id, ok := m.Tags["cpu_id"]; ok {
		key = strings.Replace(key, "cpu", "cpu"+cpu_id, 1)
	}
	if strings.HasPrefix(key, "disk.") {
		if name, ok := m.Tags["name"]; ok {
			key = strings.Replace(key, "disk.", "disk."+name+".", 1)
		}
	}
	if strings.HasPrefix(key, "df.") {
		if name, ok := m.Tags["file_system"]; ok {
			if strings.HasPrefix(name, "/") {
				key = strings.Replace(key, "df.", "df."+strings.Join(strings.Split(name, "/")[1:], ".")+".", 1)
			}
		}
	}
	if strings.HasPrefix(key, "processes.") {
		pid, _ := m.Tags["pid"]
		name, _ := m.Tags["name"]
		if name != "" && pid != "" {
			key = strings.Replace(key, "processes.", "processes."+name+"."+pid+".", 1)
		}
	}
	r = fmt.Sprintf("metrix.hosts.%s.%s %d %d", hostname, key, m.Value, t.Unix())
	return
}
