# metrix

[![Build Status](https://travis-ci.org/dynport/metrix.png?branch=master)](https://travis-ci.org/dynport/metrix)

Metrics collector written in golang

## Requirements

You need to have go 1.1 and some build tools installed. This can be done in ubuntu like this (make sure `universe` is enabled):
	
	apt-get install -y bzr git-core build-essential
	cd /opt
	curl -OL https://go.googlecode.com/files/go1.1.1.linux-amd64.tar.gz
	tar xvfz go1.1.1.linux-amd64.tar.gz
	export GOROOT=/opt/go
	export GOPATH=/go

## Installation
    
Building of the binary needs to only be done once. You can just copy it to any linux system and run it without any extra dependencies.

    git clone git@github.com:dynport/metrix.git /tmp/metrix
    cd /tmp/metrix
    make install_dependencies
    make
    make install    

## Examples

    # Post load and memory metrics to opentsdb
    metrix --memory --loadavg --opentsdb=<opentsdbhost>:<opentsdbport>

    # Post memory and loadavg metrix to opentsdb (every 60s), log output to syslog
    # 
    # /etc/cron.d/metrix
    * * * * * root /usr/local/bin/metrix --memory --loadavg --opentsdb=<opentsdbhost>:<opentsdbport> 2>&1 | logger -t metrix

## Used Keys

You can get a list of used keys for a specific metric by adding the `--keys` option. All keys should directly match to what can
be found in `man proc` or to the keys used in the stats output (examples: riak, redis, elasticsearch, etc.)

### Loadavg
    $ metrix --loadavg --keys
    load.Load1m
    load.Load5m
    load.Load15m

### Cpu

The Tags `cpu_id=1` (for cpu specific counters) or `total=total` (for the global counters) are added for OpenTSDB

    $ metrix --cpu --keys
    cpu.Ctxt
    cpu.Btime
    cpu.Processes
    cpu.ProcsRunning
    cpu.ProcsBlocked
    cpu.User
    cpu.Nice
    cpu.System
    cpu.Idle
    cpu.IOWait
    cpu.IRC
    cpu.SoftIRQ

### Disk Usage (fetched with df -k and df -i)

All metrics are tagged with `file_system` and `mounted_on` when writing to OpenTSDB.

    $ metrix --df --keys
    df.space.Total
    df.space.Used
    df.space.Available
    df.space.Use
    df.inode.Total
    df.inode.Used
    df.inode.Available
    df.inode.Use

### Network counters (fetched with netstat -s)
    $ metrix --net --keys
    net.ip.TotalPacketsReceived
    net.ip.Forwarded
    net.ip.IncomingPacketsDiscarded
    net.ip.IncomingPacketsDelivered
    net.ip.RequestsSentOut
    net.tcp.ActiveConnectionsOpenings
    net.tcp.PassiveConnectionsOpenings
    net.tcp.FailedConnectionAttempts
    net.tcp.ConnectionResetsReceived
    net.tcp.ConnectionsEstablished
    net.tcp.SegmentsReceived
    net.tcp.SegmentsSendOut
    net.tcp.SegmentsTransmitted
    net.tcp.BadSegmentsReceived
    net.tcp.ResetsSent
    net.udp.PacketsReceived
    net.udp.PacketsToUnknownPortRecived
    net.udp.PacketReceiveErrors
    net.udp.PacketsSent
    net.ip.InOctets
    net.ip.OutOctets

### Open Files (fetched with lsof)

     $ ./bin/metrix --files --keys
     files.Open

### Process

All metrics are tagged with `pid`, parent pid (`ppid`), `name` and raw `comm` value of the process when writing to OpenTSDB.

    $ metrix --processes --keys
    processes.Pid
    processes.Ppid
    processes.Pgrp
    processes.Session
    processes.TtyNr
    processes.Tpgid
    processes.Flags
    processes.Minflt
    processes.Cminflt
    processes.Majflt
    processes.Cmajflt
    processes.Utime
    processes.Stime
    processes.Cutime
    processes.Sctime
    processes.Priority
    processes.Nice
    processes.NumThreads
    processes.Itrealvalue
    processes.Starttime
    processes.Vsize
    processes.RSS
    processes.RSSlim
    processes.Startcode
    processes.Endcode
    processes.Startstac
    processes.GuestTime
    processes.CguestTime

### Disk
All metrics are tagged with `name` of the disk when writing to OpenTSDB.

    $ metrix --disk --keys
    disk.ReadsCompleted
    disk.ReadsMerged
    disk.SectorsRead
    disk.MillisecondsRead
    disk.WritesCompleted
    disk.WritesMerged
    disk.SectorsWritten
    disk.MillisecondsWritten
    disk.IosInProgress
    disk.MillisecondsIO
    disk.WeightedMillisecondsIO

### Redis

All redis metrics are tagged with the `pid` and the `port` on which redis is running. For the `redis.db` db id is tagged with `db`.

    $ metrix --redis --keys
    redis.UptimeInSeconds
    redis.memory.UsedMemory
    redis.memory.UsedMemoryRSS
    redis.memory.UsedMemoryPeak
    redis.memory.UsedMemoryLua
    redis.clients.ConnectedClients
    redis.clients.ClientLongestOutputList
    redis.clients.ClientBiggestInputBuf
    redis.clients.BlockedClients
    redis.stats.TotalConnectionsReceived
    redis.stats.TotalCommandsProcessed
    redis.stats.InstantaneousOpsPerSec
    redis.stats.RejectedConnections
    redis.stats.ExpiredKeys
    redis.stats.EvictedKeys
    redis.stats.KeyspaceHits
    redis.stats.KeyspaceMisses
    redis.stats.PubsubChannels
    redis.stats.PubsubPatterns
    redis.stats.LatestForkUsec
    redis.replication.ConnectedSlaves
    redis.db.Keys
    redis.db.Expires

### Riak

    $ metrix --redis --keys
    riak.VNodeGets
    riak.VNodeGets
    riak.VNodeGetsTotal
    riak.VNodePuts
    riak.VNodePutsTotal
    riak.VNodeIndexReads
    riak.VNodeIndexReadsTotal
    riak.VNodeIndexWrites
    riak.VNodeIndexWritesTotal
    riak.VNodeIndexWritesPostings
    riak.VNodeIndexWritesPostings_total
    riak.VNodeIndexDeletes
    riak.VNodeIndexDeletes_total
    riak.VNodeIndexDeletes_postings
    riak.VNodeIndexDeletes_postings_total
    riak.NodeGets
    riak.NodeGetsTotal
    riak.NodeGet_fsm_siblings_mean
    riak.NodegetFsmSiblingsMedian
    riak.NodegetFsmSiblings95
    riak.NodegetFsmSiblings99
    riak.NodegetFsmSiblings100
    riak.NodegetFsmObjsizeMean
    riak.NodegetFsmObjsizeMedian
    riak.NodegetFsmObjsize95
    riak.NodegetFsmObjsize99
    riak.NodegetFsmObjsize100
    riak.NodegetFsmTimeMean
    riak.NodegetFsmTimeMedian
    riak.NodegetFsmTime95
    riak.NodegetFsmTime99
    riak.NodegetFsmTime100
    riak.NodePuts
    riak.NodePutsTotal
    riak.NodePutFsmTimeMean
    riak.NodePutFsmTimeMedian
    riak.Node_put_fsm_time_95
    riak.Node_put_fsm_time_99
    riak.Node_put_fsm_time_100
    riak.ReadRepairs
    riak.ReadRepairsTotal
    riak.CoordRedirsTotal
    riak.ExecutingMappers
    riak.PrecommitFail
    riak.PostcommitFail
    riak.PbcActive
    riak.PbcConnects
    riak.PbcConnectsTotal
    riak.NodeGetFsmActive
    riak.NodeGetFsmActive60s
    riak.NodeGetFsmInRate
    riak.NodeGetFsmOutRate
    riak.NodeGetFsmRejected
    riak.NodeGetFsmRejected60s
    riak.NodeGetFsmRejectedTotal
    riak.NodePutFsmActive
    riak.NodePutFsmActive60s
    riak.NodePutFsmInRate
    riak.NodePutFsmOutRate
    riak.NodePutFsmRejected
    riak.NodePutFsmRejected60s
    riak.NodePutFsmRejectedTotal
    riak.ReadRepairsPrimaryOutofdateOne
    riak.ReadRepairsPrimaryOutofdateCount
    riak.ReadRepairsPrimaryNotfoundOne
    riak.ReadRepairsPrimaryNotfoundCount
    riak.ReadRepairsFallbackOutofdateOne
    riak.ReadRepairsFallbackOutofdateCount
    riak.ReadRepairsFallbackNotfoundOne
    riak.ReadRepairsFallbackNotfoundCount
    riak.PipelineActive
    riak.PipelineCreateCount
    riak.PipelineCreateOne
    riak.PipelineCreateErrorCount
    riak.PipelineCreateErrorOne
    riak.CpuNprocs
    riak.CpuAvg1
    riak.CpuAvg5
    riak.CpuAvg15
    riak.MemTotal
    riak.MemAllocated
    riak.SysGlobalHeapsSize
    riak.SysProcessCount
    riak.SysThreadPoolSize
    riak.SysWordsize
    riak.RingNumPartitions
    riak.RingCreationSize
    riak.MemoryTotal
    riak.MemoryProcesses
    riak.MemoryProcessesUsed
    riak.MemorySystem
    riak.MemoryAtom
    riak.MemoryAtomUsed
    riak.MemoryBinary
    riak.MemoryCode
    riak.MemoryEts
    riak.RiakCoreStatTs
    riak.IgnoredGossipTotal
    riak.RingsReconciledTotal
    riak.RingsReconciled
    riak.GossipReceived
    riak.RejectedHandoffs
    riak.HandoffTimeouts
    riak.DroppedVnodeRequestsTotal
    riak.ConvergeDelayMin
    riak.ConvergeDelayMax
    riak.ConvergeDelayMean
    riak.ConvergeDelayLast
    riak.RebalanceDelayMin
    riak.RebalanceDelayMax
    riak.RebalanceDelayMean
    riak.RebalanceDelayLast
    riak.RiakKvVnodesRunning
    riak.RiakKvVnodeqMin
    riak.RiakKvVnodeqMedian
    riak.RiakKvVnodeqMean
    riak.RiakKvVnodeqMax
    riak.RiakKvVnodeqTotal
    riak.RiakPipeVnodesRunning
    riak.RiakPipeVnodeqMin
    riak.RiakPipeVnodeqMedian
    riak.RiakPipeVnodeqMean
    riak.RiakPipeVnodeqMax
    riak.RiakPipeVnodeqTotal
    riak.ConnectedNodesCount
    riak.RingMembersCount

### Nginx

    $ metrics --nginx --keys
    nginx.ActiveConnections
    nginx.Accepts
    nginx.Handled
    nginx.Requests
    nginx.Reading
    nginx.Writing
    nginx.Waiting

### ElasticSearch

All indexes starting with `elasticsearch.indices` are tagged with the `index_name` of the specific index.

    $ metrix --elasticsearch --keys
    elasticsearch.shards.Total
    elasticsearch.shards.Successful
    elasticsearch.shards.Failed
    elasticsearch.indices.index.SizeInBytes
    elasticsearch.indices.index.PrimarySizeInBytes
    elasticsearch.indices.translog.Operations
    elasticsearch.indices.docs.NumDocs
    elasticsearch.indices.docs.MaxDoc
    elasticsearch.indices.docs.DeletedDocs
    elasticsearch.indices.merges.Current
    elasticsearch.indices.merges.CurrentDocs
    elasticsearch.indices.merges.CurrentSizeInBytes
    elasticsearch.indices.merges.Total
    elasticsearch.indices.merges.TotalTimeInMillis
    elasticsearch.indices.merges.TotalDocs
    elasticsearch.indices.merges.TotalSizeInBytes
    elasticsearch.indices.refresh.Total
    elasticsearch.indices.refresh.TotalTimeInMillis
    elasticsearch.indices.flush.Total
    elasticsearch.indices.flush.TotalTimeInMillis

### PostgreSQL

Used tags for some metrics are `table`, `index`, `database`, `pid`, `state` and `waiting`.

    $ metrix --postgres=127.0.0.1 --keys
    postgres.StatActivity
    postgres.tables.SeqScan
    postgres.tables.SeqTupRead
    postgres.tables.IdxScan
    postgres.tables.IdxTupFetch
    postgres.tables.NTupIns
    postgres.tables.NTupUpd
    postgres.tables.NTupDel
    postgres.tables.NTupHotUpd
    postgres.tables.NLiveTup
    postgres.tables.NDeadTup
    postgres.tables.VacuumCount
    postgres.tables.AutoVacuumAcount
    postgres.tables.AnalyzeCount
    postgres.tables.AutoAnalyzeCount
    postgres.databases.NumBackends
    postgres.databases.XactCommit
    postgres.databases.XactRollback
    postgres.databases.BlksRead
    postgres.databases.BlksHit
    postgres.databases.TupReturned
    postgres.databases.TupFetched
    postgres.databases.TupInserted
    postgres.databases.TupUpdated
    postgres.databases.TupDeleted
    postgres.databases.Conflicts
    postgres.databases.TempFiles
    postgres.databases.TempBytes
    postgres.databases.Deadlocks
    postgres.databases.BlkReadTime
    postgres.databases.BlkWriteTime
    postgres.indexes.IdxScan
    postgres.indexes.IdxTupRead
    postgres.indexes.IdxTupFetch

### PgBouncer

Used tags are `database`, `state`, `address`, `port`, `local_address`, `local_port` and `type`.

	$ ./bin/metrix --pgbouncer=127.0.0.1:6432 --keys
	pgbouncer.connections.Total
	pgbouncer.memory.Free
	pgbouncer.memory.MemTotal
	pgbouncer.memory.Size
	pgbouncer.memory.Used
	pgbouncer.pools.ClientsActive
	pgbouncer.pools.ClientsWaiting
	pgbouncer.pools.MaxWait
	pgbouncer.pools.ServersActive
	pgbouncer.pools.ServersIdle
	pgbouncer.pools.ServersLogin
	pgbouncer.pools.ServersTested
	pgbouncer.pools.ServersUsed
	pgbouncer.sockets.PktAvail
	pgbouncer.sockets.PktPos
	pgbouncer.sockets.PktRemain
	pgbouncer.sockets.RecvPos
	pgbouncer.sockets.SendAvail
	pgbouncer.sockets.SendPos
	pgbouncer.sockets.SendRemain
	pgbouncer.stats.AvgQuery
	pgbouncer.stats.AvgRecv
	pgbouncer.stats.AvgReq
	pgbouncer.stats.AvgSent
	pgbouncer.stats.TotalQueryTime
	pgbouncer.stats.TotalReceived
	pgbouncer.stats.TotalRequests
	pgbouncer.stats.TotalSent

## Help
    USAGE: ./bin/metrix
      --help           	Print this usage page                                               
      --keys           	Only list all known keys                                            
      --opentsdb       	Report metrics to OpenTSDB host.                                    
                        EXAMPLE: opentsdb.host:4242                                         

      --graphite       	Report metrics to Graphite host.                                    
                        EXAMPLE: graphite.host:2003                                         

      --hostname       	Hostname to be used for tagging. If blank the local hostname is used
      --loadavg        	Collect loadvg metrics                                              
      --memory         	Collect memory metrics                                              
      --cpu            	Collect cpu metrics                                                 
      --disk           	Collect disk usage metrics                                          
      --processes      	Collect metrics for processes                                       
      --net            	Collect network metrics                                             
      --df             	Collect disk free space metrics                                     
      --riak           	Collect riak metrics                                                
                        DEFAULT: http://127.0.0.1:8098/stats                                

      --elasticsearch  	Collect ElasticSearch metrics                                       
                        DEFAULT: http://127.0.0.1:9200/_status                              

      --redis          	Collect redis metrics                                               
                        DEFAULT: 127.0.0.1:6379                                             

      --postgres       	Collect postgres metrics.                                           
                        EXAMPLE: psql://user:pwd@host/db                                    

      --pgbouncer      	Collect pgbouncer metrics                                           
                        DEFAULT: 127.0.0.1:6432                                             

      --nginx          	Collect nginx metrics                                               
                        DEFAULT: http://127.0.0.1:8080                                      
