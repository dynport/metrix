package main

import (
	"encoding/json"
)

const RIAK = "riak"

func init() {
	parser.Add(RIAK, "http://127.0.0.1:8098/stats", "Collect riak metrics")
}

type RiakStatus struct {
	VNodeGets int64 `json:"vnode_gets"`
	VNodeGetsTotal int64 `json:"vnode_gets_total"`
	VNodePuts int64 `json:"vnode_puts"`
	VNodePutsTotal int64 `json:"vnode_puts_total"`
	VNodeIndexReads int64 `json:"vnode_index_reads"`
	VNodeIndexReadsTotal int64 `json:"vnode_index_reads_total"`
	VNodeIndexWrites int64 `json:"vnode_index_writes"`
	VNodeIndexWritesTotal int64 `json:"vnode_index_writes_total"`
	VNodeIndexWritesPostings int64 `json:"vnode_index_writes_postings"`
	VNodeIndexWritesPostings_total int64 `json:"vnode_index_writes_postings_total"`
	VNodeIndexDeletes int64 `json:"vnode_index_deletes"`
	VNodeIndexDeletes_total int64 `json:"vnode_index_deletes_total"`
	VNodeIndexDeletes_postings int64 `json:"vnode_index_deletes_postings"`
	VNodeIndexDeletes_postings_total int64 `json:"vnode_index_deletes_postings_total"`
	NodeGets int64 `json:"node_gets"`
	NodeGetsTotal int64 `json:"node_gets_total"`
	NodeGet_fsm_siblings_mean int64 `json:"node_get_fsm_siblings_mean"`
	NodegetFsmSiblingsMedian int64 `json:"node_get_fsm_siblings_median"`
	NodegetFsmSiblings95 int64 `json:"node_get_fsm_siblings_95"`
	NodegetFsmSiblings99 int64 `json:"node_get_fsm_siblings_99"`
	NodegetFsmSiblings100 int64 `json:"node_get_fsm_siblings_100"`
	NodegetFsmObjsizeMean int64 `json:"node_get_fsm_objsize_mean"`
	NodegetFsmObjsizeMedian int64 `json:"node_get_fsm_objsize_median"`
	NodegetFsmObjsize95 int64 `json:"node_get_fsm_objsize_95"`
	NodegetFsmObjsize99 int64 `json:"node_get_fsm_objsize_99"`
	NodegetFsmObjsize100 int64 `json:"node_get_fsm_objsize_100"`
	NodegetFsmTimeMean int64 `json:"node_get_fsm_time_mean"`
	NodegetFsmTimeMedian int64 `json:"node_get_fsm_time_median"`
	NodegetFsmTime95 int64 `json:"node_get_fsm_time_95"`
	NodegetFsmTime99 int64 `json:"node_get_fsm_time_99"`
	NodegetFsmTime100 int64 `json:"node_get_fsm_time_100"`
	NodePuts int64 `json:"node_puts"`
	NodePutsTotal int64 `json:"node_puts_total"`
	NodePutFsmTimeMean int64 `json:"node_put_fsm_time_mean"`
	NodePutFsmTimeMedian int64 `json:"node_put_fsm_time_median"`
	Node_put_fsm_time_95 int64 `json:"node_put_fsm_time_95"`
	Node_put_fsm_time_99 int64 `json:"node_put_fsm_time_99"`
	Node_put_fsm_time_100 int64 `json:"node_put_fsm_time_100"`
	ReadRepairs int64 `json:"read_repairs"`
	ReadRepairsTotal int64 `json:"read_repairs_total"`
	CoordRedirsTotal int64 `json:"coord_redirs_total"`
	ExecutingMappers int64 `json:"executing_mappers"`
	PrecommitFail int64 `json:"precommit_fail"`
	PostcommitFail int64 `json:"postcommit_fail"`
	PbcActive int64 `json:"pbc_active"`
	PbcConnects int64 `json:"pbc_connects"`
	PbcConnectsTotal int64 `json:"pbc_connects_total"`
	NodeGetFsmActive int64 `json:"node_get_fsm_active"`
	NodeGetFsmActive60s int64 `json:"node_get_fsm_active_60s"`
	NodeGetFsmInRate int64 `json:"node_get_fsm_in_rate"`
	NodeGetFsmOutRate int64 `json:"node_get_fsm_out_rate"`
	NodeGetFsmRejected int64 `json:"node_get_fsm_rejected"`
	NodeGetFsmRejected60s int64 `json:"node_get_fsm_rejected_60s"`
	NodeGetFsmRejectedTotal int64 `json:"node_get_fsm_rejected_total"`
	NodePutFsmActive int64 `json:"node_put_fsm_active"`
	NodePutFsmActive60s int64 `json:"node_put_fsm_active_60s"`
	NodePutFsmInRate int64 `json:"node_put_fsm_in_rate"`
	NodePutFsmOutRate int64 `json:"node_put_fsm_out_rate"`
	NodePutFsmRejected int64 `json:"node_put_fsm_rejected"`
	NodePutFsmRejected60s int64 `json:"node_put_fsm_rejected_60s"`
	NodePutFsmRejectedTotal int64 `json:"node_put_fsm_rejected_total"`
	ReadRepairsPrimaryOutofdateOne int64 `json:"read_repairs_primary_outofdate_one"`
	ReadRepairsPrimaryOutofdateCount int64 `json:"read_repairs_primary_outofdate_count"`
	ReadRepairsPrimaryNotfoundOne int64 `json:"read_repairs_primary_notfound_one"`
	ReadRepairsPrimaryNotfoundCount int64 `json:"read_repairs_primary_notfound_count"`
	ReadRepairsFallbackOutofdateOne int64 `json:"read_repairs_fallback_outofdate_one"`
	ReadRepairsFallbackOutofdateCount int64 `json:"read_repairs_fallback_outofdate_count"`
	ReadRepairsFallbackNotfoundOne int64 `json:"read_repairs_fallback_notfound_one"`
	ReadRepairsFallbackNotfoundCount int64 `json:"read_repairs_fallback_notfound_count"`

	PipelineActive int64 `json:"pipeline_active"`
	PipelineCreateCount int64 `json:"pipeline_create_count"`
	PipelineCreateOne int64 `json:"pipeline_create_one"`
	PipelineCreateErrorCount int64 `json:"pipeline_create_error_count"`
	PipelineCreateErrorOne int64 `json:"pipeline_create_error_one"`
	CpuNprocs int64 `json:"cpu_nprocs"`
	CpuAvg1 int64 `json:"cpu_avg1"`
	CpuAvg5 int64 `json:"cpu_avg5"`
	CpuAvg15 int64 `json:"cpu_avg15"`
	MemTotal int64 `json:"mem_total"`
	MemAllocated int64 `json:"mem_allocated"`
	SysGlobalHeapsSize int64 `json:"sys_global_heaps_size"`
	SysProcessCount int64 `json:"sys_process_count"`
	SysThreadPoolSize int64 `json:"sys_thread_pool_size"`
	SysWordsize int64 `json:"sys_wordsize"`
	RingNumPartitions int64 `json:"ring_num_partitions"`
	RingCreationSize int64 `json:"ring_creation_size"`
	MemoryTotal int64 `json:"memory_total"`
	MemoryProcesses int64 `json:"memory_processes"`
	MemoryProcessesUsed int64 `json:"memory_processes_used"`
	MemorySystem int64 `json:"memory_system"`
	MemoryAtom int64 `json:"memory_atom"`
	MemoryAtomUsed int64 `json:"memory_atom_used"`
	MemoryBinary int64 `json:"memory_binary"`
	MemoryCode int64 `json:"memory_code"`
	MemoryEts int64 `json:"memory_ets"`
	RiakCoreStatTs int64 `json:"riak_core_stat_ts"`
	IgnoredGossipTotal int64 `json:"ignored_gossip_total"`
	RingsReconciledTotal int64 `json:"rings_reconciled_total"`
	RingsReconciled int64 `json:"rings_reconciled"`
	GossipReceived int64 `json:"gossip_received"`
	RejectedHandoffs int64 `json:"rejected_handoffs"`
	HandoffTimeouts int64 `json:"handoff_timeouts"`
	DroppedVnodeRequestsTotal int64 `json:"dropped_vnode_requests_total"`
	ConvergeDelayMin int64 `json:"converge_delay_min"`
	ConvergeDelayMax int64 `json:"converge_delay_max"`
	ConvergeDelayMean int64 `json:"converge_delay_mean"`
	ConvergeDelayLast int64 `json:"converge_delay_last"`
	RebalanceDelayMin int64 `json:"rebalance_delay_min"`
	RebalanceDelayMax int64 `json:"rebalance_delay_max"`
	RebalanceDelayMean int64 `json:"rebalance_delay_mean"`
	RebalanceDelayLast int64 `json:"rebalance_delay_last"`
	RiakKvVnodesRunning int64 `json:"riak_kv_vnodes_running"`
	RiakKvVnodeqMin int64 `json:"riak_kv_vnodeq_min"`
	RiakKvVnodeqMedian int64 `json:"riak_kv_vnodeq_median"`
	RiakKvVnodeqMean int64 `json:"riak_kv_vnodeq_mean"`
	RiakKvVnodeqMax int64 `json:"riak_kv_vnodeq_max"`
	RiakKvVnodeqTotal int64 `json:"riak_kv_vnodeq_total"`
	RiakPipeVnodesRunning int64 `json:"riak_pipe_vnodes_running"`
	RiakPipeVnodeqMin int64 `json:"riak_pipe_vnodeq_min"`
	RiakPipeVnodeqMedian int64 `json:"riak_pipe_vnodeq_median"`
	RiakPipeVnodeqMean int64 `json:"riak_pipe_vnodeq_mean"`
	RiakPipeVnodeqMax int64 `json:"riak_pipe_vnodeq_max"`
	RiakPipeVnodeqTotal int64 `json:"riak_pipe_vnodeq_total"`

	ConnectedNodes []string `json:"connected_nodes"`
	RingMembers []string `json:"ring_members"`
}

func (self *Riak) Keys() []string {
	return []string {
		"VNodeGets", "VNodeGets", "VNodeGetsTotal", "VNodePuts", "VNodePutsTotal", "VNodeIndexReads", "VNodeIndexReadsTotal", "VNodeIndexWrites",
		"VNodeIndexWritesTotal", "VNodeIndexWritesPostings", "VNodeIndexWritesPostings_total", "VNodeIndexDeletes", "VNodeIndexDeletes_total",
		"VNodeIndexDeletes_postings", "VNodeIndexDeletes_postings_total", "NodeGets", "NodeGetsTotal", "NodeGet_fsm_siblings_mean", "NodegetFsmSiblingsMedian",
		"NodegetFsmSiblings95", "NodegetFsmSiblings99", "NodegetFsmSiblings100", "NodegetFsmObjsizeMean", "NodegetFsmObjsizeMedian", "NodegetFsmObjsize95",
		"NodegetFsmObjsize99", "NodegetFsmObjsize100", "NodegetFsmTimeMean", "NodegetFsmTimeMedian", "NodegetFsmTime95", "NodegetFsmTime99",
		"NodegetFsmTime100", "NodePuts", "NodePutsTotal", "NodePutFsmTimeMean", "NodePutFsmTimeMedian", "Node_put_fsm_time_95", "Node_put_fsm_time_99",
		"Node_put_fsm_time_100", "ReadRepairs", "ReadRepairsTotal", "CoordRedirsTotal", "ExecutingMappers", "PrecommitFail", "PostcommitFail", "PbcActive",
		"PbcConnects", "PbcConnectsTotal", "NodeGetFsmActive", "NodeGetFsmActive60s", "NodeGetFsmInRate", "NodeGetFsmOutRate", "NodeGetFsmRejected",
		"NodeGetFsmRejected60s", "NodeGetFsmRejectedTotal", "NodePutFsmActive", "NodePutFsmActive60s", "NodePutFsmInRate", "NodePutFsmOutRate",
		"NodePutFsmRejected", "NodePutFsmRejected60s", "NodePutFsmRejectedTotal", "ReadRepairsPrimaryOutofdateOne", "ReadRepairsPrimaryOutofdateCount",
		"ReadRepairsPrimaryNotfoundOne", "ReadRepairsPrimaryNotfoundCount", "ReadRepairsFallbackOutofdateOne", "ReadRepairsFallbackOutofdateCount",
		"ReadRepairsFallbackNotfoundOne", "ReadRepairsFallbackNotfoundCount", "PipelineActive", "PipelineCreateCount", "PipelineCreateOne",
		"PipelineCreateErrorCount", "PipelineCreateErrorOne", "CpuNprocs", "CpuAvg1", "CpuAvg5", "CpuAvg15", "MemTotal", "MemAllocated", "SysGlobalHeapsSize",
		"SysProcessCount", "SysThreadPoolSize", "SysWordsize", "RingNumPartitions", "RingCreationSize", "MemoryTotal", "MemoryProcesses",
		"MemoryProcessesUsed", "MemorySystem", "MemoryAtom", "MemoryAtomUsed", "MemoryBinary", "MemoryCode", "MemoryEts", "RiakCoreStatTs",
		"IgnoredGossipTotal", "RingsReconciledTotal", "RingsReconciled", "GossipReceived", "RejectedHandoffs", "HandoffTimeouts", "DroppedVnodeRequestsTotal",
		"ConvergeDelayMin", "ConvergeDelayMax", "ConvergeDelayMean", "ConvergeDelayLast", "RebalanceDelayMin", "RebalanceDelayMax", "RebalanceDelayMean",
		"RebalanceDelayLast", "RiakKvVnodesRunning", "RiakKvVnodeqMin", "RiakKvVnodeqMedian", "RiakKvVnodeqMean", "RiakKvVnodeqMax", "RiakKvVnodeqTotal",
		"RiakPipeVnodesRunning", "RiakPipeVnodeqMin", "RiakPipeVnodeqMedian", "RiakPipeVnodeqMean", "RiakPipeVnodeqMax", "RiakPipeVnodeqTotal",
		"ConnectedNodesCount", "RingMembersCount",
	}
}

func (self *Riak) Prefix() string {
	return "riak"
}

func (r* RiakStatus) Collect(c* MetricsCollection) {
	c.Add("VNodeGets", r.VNodeGets)
	c.Add("VNodeGets", r.VNodeGets)
	c.Add("VNodeGetsTotal", r.VNodeGetsTotal)
	c.Add("VNodePuts", r.VNodePuts)
	c.Add("VNodePutsTotal", r.VNodePutsTotal)
	c.Add("VNodeIndexReads", r.VNodeIndexReads)
	c.Add("VNodeIndexReadsTotal", r.VNodeIndexReadsTotal)
	c.Add("VNodeIndexWrites", r.VNodeIndexWrites)
	c.Add("VNodeIndexWritesTotal", r.VNodeIndexWritesTotal)
	c.Add("VNodeIndexWritesPostings", r.VNodeIndexWritesPostings)
	c.Add("VNodeIndexWritesPostings_total", r.VNodeIndexWritesPostings_total)
	c.Add("VNodeIndexDeletes", r.VNodeIndexDeletes)
	c.Add("VNodeIndexDeletes_total", r.VNodeIndexDeletes_total)
	c.Add("VNodeIndexDeletes_postings", r.VNodeIndexDeletes_postings)
	c.Add("VNodeIndexDeletes_postings_total", r.VNodeIndexDeletes_postings_total)
	c.Add("NodeGets", r.NodeGets)
	c.Add("NodeGetsTotal", r.NodeGetsTotal)
	c.Add("NodeGet_fsm_siblings_mean", r.NodeGet_fsm_siblings_mean)
	c.Add("NodegetFsmSiblingsMedian", r.NodegetFsmSiblingsMedian)
	c.Add("NodegetFsmSiblings95", r.NodegetFsmSiblings95)
	c.Add("NodegetFsmSiblings99", r.NodegetFsmSiblings99)
	c.Add("NodegetFsmSiblings100", r.NodegetFsmSiblings100)
	c.Add("NodegetFsmObjsizeMean", r.NodegetFsmObjsizeMean)
	c.Add("NodegetFsmObjsizeMedian", r.NodegetFsmObjsizeMedian)
	c.Add("NodegetFsmObjsize95", r.NodegetFsmObjsize95)
	c.Add("NodegetFsmObjsize99", r.NodegetFsmObjsize99)
	c.Add("NodegetFsmObjsize100", r.NodegetFsmObjsize100)
	c.Add("NodegetFsmTimeMean", r.NodegetFsmTimeMean)
	c.Add("NodegetFsmTimeMedian", r.NodegetFsmTimeMedian)
	c.Add("NodegetFsmTime95", r.NodegetFsmTime95)
	c.Add("NodegetFsmTime99", r.NodegetFsmTime99)
	c.Add("NodegetFsmTime100", r.NodegetFsmTime100)
	c.Add("NodePuts", r.NodePuts)
	c.Add("NodePutsTotal", r.NodePutsTotal)
	c.Add("NodePutFsmTimeMean", r.NodePutFsmTimeMean)
	c.Add("NodePutFsmTimeMedian", r.NodePutFsmTimeMedian)
	c.Add("Node_put_fsm_time_95", r.Node_put_fsm_time_95)
	c.Add("Node_put_fsm_time_99", r.Node_put_fsm_time_99)
	c.Add("Node_put_fsm_time_100", r.Node_put_fsm_time_100)
	c.Add("ReadRepairs", r.ReadRepairs)
	c.Add("ReadRepairsTotal", r.ReadRepairsTotal)
	c.Add("CoordRedirsTotal", r.CoordRedirsTotal)
	c.Add("ExecutingMappers", r.ExecutingMappers)
	c.Add("PrecommitFail", r.PrecommitFail)
	c.Add("PostcommitFail", r.PostcommitFail)
	c.Add("PbcActive", r.PbcActive)
	c.Add("PbcConnects", r.PbcConnects)
	c.Add("PbcConnectsTotal", r.PbcConnectsTotal)
	c.Add("NodeGetFsmActive", r.NodeGetFsmActive)
	c.Add("NodeGetFsmActive60s", r.NodeGetFsmActive60s)
	c.Add("NodeGetFsmInRate", r.NodeGetFsmInRate)
	c.Add("NodeGetFsmOutRate", r.NodeGetFsmOutRate)
	c.Add("NodeGetFsmRejected", r.NodeGetFsmRejected)
	c.Add("NodeGetFsmRejected60s", r.NodeGetFsmRejected60s)
	c.Add("NodeGetFsmRejectedTotal", r.NodeGetFsmRejectedTotal)
	c.Add("NodePutFsmActive", r.NodePutFsmActive)
	c.Add("NodePutFsmActive60s", r.NodePutFsmActive60s)
	c.Add("NodePutFsmInRate", r.NodePutFsmInRate)
	c.Add("NodePutFsmOutRate", r.NodePutFsmOutRate)
	c.Add("NodePutFsmRejected", r.NodePutFsmRejected)
	c.Add("NodePutFsmRejected60s", r.NodePutFsmRejected60s)
	c.Add("NodePutFsmRejectedTotal", r.NodePutFsmRejectedTotal)
	c.Add("ReadRepairsPrimaryOutofdateOne", r.ReadRepairsPrimaryOutofdateOne)
	c.Add("ReadRepairsPrimaryOutofdateCount", r.ReadRepairsPrimaryOutofdateCount)
	c.Add("ReadRepairsPrimaryNotfoundOne", r.ReadRepairsPrimaryNotfoundOne)
	c.Add("ReadRepairsPrimaryNotfoundCount", r.ReadRepairsPrimaryNotfoundCount)
	c.Add("ReadRepairsFallbackOutofdateOne", r.ReadRepairsFallbackOutofdateOne)
	c.Add("ReadRepairsFallbackOutofdateCount", r.ReadRepairsFallbackOutofdateCount)
	c.Add("ReadRepairsFallbackNotfoundOne", r.ReadRepairsFallbackNotfoundOne)
	c.Add("ReadRepairsFallbackNotfoundCount", r.ReadRepairsFallbackNotfoundCount)
	c.Add("PipelineActive", r.PipelineActive)
	c.Add("PipelineCreateCount", r.PipelineCreateCount)
	c.Add("PipelineCreateOne", r.PipelineCreateOne)
	c.Add("PipelineCreateErrorCount", r.PipelineCreateErrorCount)
	c.Add("PipelineCreateErrorOne", r.PipelineCreateErrorOne)
	c.Add("CpuNprocs", r.CpuNprocs)
	c.Add("CpuAvg1", r.CpuAvg1)
	c.Add("CpuAvg5", r.CpuAvg5)
	c.Add("CpuAvg15", r.CpuAvg15)
	c.Add("MemTotal", r.MemTotal)
	c.Add("MemAllocated", r.MemAllocated)
	c.Add("SysGlobalHeapsSize", r.SysGlobalHeapsSize)
	c.Add("SysProcessCount", r.SysProcessCount)
	c.Add("SysThreadPoolSize", r.SysThreadPoolSize)
	c.Add("SysWordsize", r.SysWordsize)
	c.Add("RingNumPartitions", r.RingNumPartitions)
	c.Add("RingCreationSize", r.RingCreationSize)
	c.Add("MemoryTotal", r.MemoryTotal)
	c.Add("MemoryProcesses", r.MemoryProcesses)
	c.Add("MemoryProcessesUsed", r.MemoryProcessesUsed)
	c.Add("MemorySystem", r.MemorySystem)
	c.Add("MemoryAtom", r.MemoryAtom)
	c.Add("MemoryAtomUsed", r.MemoryAtomUsed)
	c.Add("MemoryBinary", r.MemoryBinary)
	c.Add("MemoryCode", r.MemoryCode)
	c.Add("MemoryEts", r.MemoryEts)
	c.Add("RiakCoreStatTs", r.RiakCoreStatTs)
	c.Add("IgnoredGossipTotal", r.IgnoredGossipTotal)
	c.Add("RingsReconciledTotal", r.RingsReconciledTotal)
	c.Add("RingsReconciled", r.RingsReconciled)
	c.Add("GossipReceived", r.GossipReceived)
	c.Add("RejectedHandoffs", r.RejectedHandoffs)
	c.Add("HandoffTimeouts", r.HandoffTimeouts)
	c.Add("DroppedVnodeRequestsTotal", r.DroppedVnodeRequestsTotal)
	c.Add("ConvergeDelayMin", r.ConvergeDelayMin)
	c.Add("ConvergeDelayMax", r.ConvergeDelayMax)
	c.Add("ConvergeDelayMean", r.ConvergeDelayMean)
	c.Add("ConvergeDelayLast", r.ConvergeDelayLast)
	c.Add("RebalanceDelayMin", r.RebalanceDelayMin)
	c.Add("RebalanceDelayMax", r.RebalanceDelayMax)
	c.Add("RebalanceDelayMean", r.RebalanceDelayMean)
	c.Add("RebalanceDelayLast", r.RebalanceDelayLast)
	c.Add("RiakKvVnodesRunning", r.RiakKvVnodesRunning)
	c.Add("RiakKvVnodeqMin", r.RiakKvVnodeqMin)
	c.Add("RiakKvVnodeqMedian", r.RiakKvVnodeqMedian)
	c.Add("RiakKvVnodeqMean", r.RiakKvVnodeqMean)
	c.Add("RiakKvVnodeqMax", r.RiakKvVnodeqMax)
	c.Add("RiakKvVnodeqTotal", r.RiakKvVnodeqTotal)
	c.Add("RiakPipeVnodesRunning", r.RiakPipeVnodesRunning)
	c.Add("RiakPipeVnodeqMin", r.RiakPipeVnodeqMin)
	c.Add("RiakPipeVnodeqMedian", r.RiakPipeVnodeqMedian)
	c.Add("RiakPipeVnodeqMean", r.RiakPipeVnodeqMean)
	c.Add("RiakPipeVnodeqMax", r.RiakPipeVnodeqMax)
	c.Add("RiakPipeVnodeqTotal", r.RiakPipeVnodeqTotal)
	c.Add("ConnectedNodesCount", int64(len(r.ConnectedNodes)))
	c.Add("RingMembersCount", int64(len(r.RingMembers)))
	return
}

type Riak struct {
	Address string
	Raw []byte
}

func (r* Riak) Collect(c* MetricsCollection) (e error) {
	if len(r.Raw) == 0 {
		r.Raw, e = FetchURL(r.Address)
		if e != nil {
			return
		}
	}
	status, e := r.ParseRiakStatus(r.Raw)
	if e != nil {
		return
	}
	status.Collect(c)
	return
}

func (r* Riak) ParseRiakStatus(b []byte) (s* RiakStatus, e error) {
	s = &RiakStatus{}
	e = json.Unmarshal(b, s)
	return
}
