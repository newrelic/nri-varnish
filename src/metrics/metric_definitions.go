package metrics

// varnishDefinition represents the data to be fed to the VarnishSample event
type varnishDefinition struct {
	MgtUptime                interface{} `stat_name:"uptime" metric_name:"mgt.uptimeInMilliseconds" source_type:"Rate"`
	MgtChildStart            interface{} `stat_name:"child_start" metric_name:"mgt.childStart" source_type:"Rate"`
	MgtChildExit             interface{} `stat_name:"child_exit" metric_name:"mgt.childExit" source_type:"Rate"`
	MgtChildStop             interface{} `stat_name:"child_stop" metric_name:"mgt.childStop" source_type:"Rate"`
	MgtChildDied             interface{} `stat_name:"child_died" metric_name:"mgt.childDied" source_type:"Rate"`
	MgtChildDump             interface{} `stat_name:"child_dump" metric_name:"mgt.childDump" source_type:"Rate"`
	MgtChildPanic            interface{} `stat_name:"child_panic" metric_name:"mgt.childPanic" source_type:"Rate"`
	MainSumms                interface{} `stat_name:"summs" metric_name:"main.summs" source_type:"Rate"`
	MainUptime               interface{} `stat_name:"uptime" metric_name:"main.uptimeInMilliseconds" source_type:"Rate"`
	SessionConnections       interface{} `stat_name:"sess_conn" metric_name:"session.connections" source_type:"Rate"`
	SessionDrops             interface{} `stat_name:"sess_drop" metric_name:"session.drops" source_type:"Rate"`
	SessionFail              interface{} `stat_name:"sess_fail" metric_name:"session.fail" source_type:"Rate"`
	Net400Errors             interface{} `stat_name:"client_req_400" metric_name:"net.400Errors" source_type:"Rate"`
	Net417Errors             interface{} `stat_name:"client_req_417" metric_name:"net.417Errors" source_type:"Rate"`
	NetRequests              interface{} `stat_name:"client_req" metric_name:"net.requests" source_type:"Rate"`
	CacheHits                interface{} `stat_name:"cache_hit" metric_name:"cache.hits" source_type:"Rate"`
	CacheGraceHits           interface{} `stat_name:"cache_hit_grace" metric_name:"cache.graceHits" source_type:"Rate"`
	CachePassHits            interface{} `stat_name:"cache_hitpass" metric_name:"cache.passHits" source_type:"Rate"`
	CacheMissHits            interface{} `stat_name:"cache_hitmiss" metric_name:"cache.missHits" source_type:"Rate"`
	CacheMisses              interface{} `stat_name:"cache_miss" metric_name:"cache.misses" source_type:"Rate"`
	BackendConSuccess        interface{} `stat_name:"backend_conn" metric_name:"backend.conSuccess" source_type:"Rate"`
	BackendConUnHealthy      interface{} `stat_name:"backend_unhealthy" metric_name:"backend.conUnHealthy" source_type:"Rate"`
	BackendConBusy           interface{} `stat_name:"backend_busy" metric_name:"backend.conBusy" source_type:"Rate"`
	BackendConFails          interface{} `stat_name:"backend_fail" metric_name:"backend.conFails" source_type:"Rate"`
	BackendConReuses         interface{} `stat_name:"backend_reuse" metric_name:"backend.conReuses" source_type:"Rate"`
	BackendConRecycles       interface{} `stat_name:"backend_recycle" metric_name:"backend.conRecycles" source_type:"Rate"`
	BackendConRetries        interface{} `stat_name:"backend_retry" metric_name:"backend.conRetries" source_type:"Rate"`
	FetchHead                interface{} `stat_name:"fetch_head" metric_name:"fetch.head" source_type:"Rate"`
	FetchContentLength       interface{} `stat_name:"fetch_length" metric_name:"fetch.contentLength" source_type:"Rate"`
	FetchChuncked            interface{} `stat_name:"fetch_chunked" metric_name:"fetch.chuncked" source_type:"Rate"`
	FetchEOF                 interface{} `stat_name:"fetch_eof" metric_name:"fetch.eof" source_type:"Rate"`
	FetchBad                 interface{} `stat_name:"fetch_bad" metric_name:"fetch.bad" source_type:"Rate"`
	FetchNoBody              interface{} `stat_name:"fetch_none" metric_name:"fetch.noBody" source_type:"Rate"`
	FetchNoBody1xx           interface{} `stat_name:"fetch_1xx" metric_name:"fetch.noBody1xx" source_type:"Rate"`
	FetchNoBody204           interface{} `stat_name:"fetch_204" metric_name:"fetch.noBody204" source_type:"Rate"`
	FetchNoBody304           interface{} `stat_name:"fetch_304" metric_name:"fetch.noBody304" source_type:"Rate"`
	FetchFailed              interface{} `stat_name:"fetch_failed" metric_name:"fetch.failed" source_type:"Rate"`
	FetchNoThreadFail        interface{} `stat_name:"fetch_no_thread" metric_name:"fetch.noThreadFail" source_type:"Rate"`
	MainPools                interface{} `stat_name:"pools" metric_name:"main.pools" source_type:"Gauge"`
	MainThreads              interface{} `stat_name:"threads" metric_name:"main.threads" source_type:"Gauge"`
	MainThreadsLimited       interface{} `stat_name:"threads_limited" metric_name:"main.threadsLimited" source_type:"Rate"`
	MainThreadsCreated       interface{} `stat_name:"threads_created" metric_name:"main.threadsCreated" source_type:"Rate"`
	MainThreadsDestroyed     interface{} `stat_name:"threads_destroyed" metric_name:"main.threadsDestroyed" source_type:"Rate"`
	MainThreadsFailed        interface{} `stat_name:"threads_failed" metric_name:"main.threadsFailed" source_type:"Rate"`
	MainSessQueueLength      interface{} `stat_name:"thread_queue_len" metric_name:"main.sessQueueLength" source_type:"Gauge"`
	MainBusySleep            interface{} `stat_name:"busy_sleep" metric_name:"main.busySleep" source_type:"Rate"`
	MainBusyWakeup           interface{} `stat_name:"busy_wakeup" metric_name:"main.busyWakeup" source_type:"Rate"`
	MainBusyKilled           interface{} `stat_name:"busy_killed" metric_name:"main.busyKilled" source_type:"Rate"`
	SessQueued               interface{} `stat_name:"sess_queued" metric_name:"sess.queued" source_type:"Rate"`
	SessDropped              interface{} `stat_name:"sess_dropped" metric_name:"sess.dropped" source_type:"Rate"`
	MainReqDropped           interface{} `stat_name:"req_dropped" metric_name:"main.reqDropped" source_type:"Rate"`
	MainObjects              interface{} `stat_name:"n_object" metric_name:"main.objects" source_type:"Gauge"`
	MainUnresurrectedObjects interface{} `stat_name:"n_vampireobject" metric_name:"main.unresurrectedObjects" source_type:"Gauge"`
	MainObjectcores          interface{} `stat_name:"n_objectcore" metric_name:"main.objectcores" source_type:"Gauge"`
	MainObjectheads          interface{} `stat_name:"n_objecthead" metric_name:"main.objectheads" source_type:"Gauge"`
	MainBackends             interface{} `stat_name:"n_backend" metric_name:"main.backends" source_type:"Gauge"`
	MainExpired              interface{} `stat_name:"n_expired" metric_name:"main.expired" source_type:"Rate"`
	LruNuked                 interface{} `stat_name:"n_lru_nuked" metric_name:"lru.nuked" source_type:"Rate"`
	LruMoved                 interface{} `stat_name:"n_lru_moved" metric_name:"lru.moved" source_type:"Rate"`
	LruLimited               interface{} `stat_name:"n_lru_limited" metric_name:"lru.limited" source_type:"Rate"`
	NetHTTPOverflow          interface{} `stat_name:"losthdr" metric_name:"net.httpOverflow" source_type:"Rate"`
	MainSessions             interface{} `stat_name:"s_sess" metric_name:"main.sessions" source_type:"Rate"`
	MainPipeSessions         interface{} `stat_name:"s_pipe" metric_name:"main.pipeSessions" source_type:"Rate"`
	MainPassedReq            interface{} `stat_name:"s_pass" metric_name:"main.passedReq" source_type:"Rate"`
	BackendFetches           interface{} `stat_name:"s_fetch" metric_name:"backend.fetches" source_type:"Rate"`
	MainSyntheticResp        interface{} `stat_name:"s_synth" metric_name:"main.syntheticResp" source_type:"Rate"`
	NetReqHeader             interface{} `stat_name:"s_req_hdrbytes" metric_name:"net.req.headerInBytes" source_type:"Rate"`
	NetReqBody               interface{} `stat_name:"s_req_bodybytes" metric_name:"net.req.bodyInBytes" source_type:"Rate"`
	NetRespHeader            interface{} `stat_name:"s_resp_hdrbytes" metric_name:"net.resp.headerInBytes" source_type:"Rate"`
	NetRespBody              interface{} `stat_name:"s_resp_bodybytes" metric_name:"net.resp.bodyInBytes" source_type:"Rate"`
	NetPipereqHeader         interface{} `stat_name:"s_pipe_hdrbytes" metric_name:"net.pipereq.headerInBytes" source_type:"Rate"`
	NetPipeIn                interface{} `stat_name:"s_pipe_in" metric_name:"net.pipe.inInBytes" source_type:"Rate"`
	NetPipeOut               interface{} `stat_name:"s_pipe_out" metric_name:"net.pipe.outInBytes" source_type:"Rate"`
	SessClosed               interface{} `stat_name:"sess_closed" metric_name:"sess.closed" source_type:"Rate"`
	SessClosedError          interface{} `stat_name:"sess_closed_err" metric_name:"sess.closedError" source_type:"Rate"`
	SessReadAhead            interface{} `stat_name:"sess_readahead" metric_name:"sess.readAhead" source_type:"Rate"`
	SessHerd                 interface{} `stat_name:"sess_herd" metric_name:"sess.herd" source_type:"Rate"`
	SessClientClose          interface{} `stat_name:"sc_rem_close" metric_name:"sess.clientClose" source_type:"Rate"`
	SessClientReqClose       interface{} `stat_name:"sc_req_close" metric_name:"sess.clientReqClose" source_type:"Rate"`
	SessReqHTTP10Close       interface{} `stat_name:"sc_req_http10" metric_name:"sess.reqHTTP10Close" source_type:"Rate"`
	SessBadClose             interface{} `stat_name:"sc_rx_bad" metric_name:"sess.badClose" source_type:"Rate"`
	SessBodyFailClose        interface{} `stat_name:"sc_rx_body" metric_name:"sess.bodyFailClose" source_type:"Rate"`
	SessJunkClose            interface{} `stat_name:"sc_rx_junk" metric_name:"sess.junkClose" source_type:"Rate"`
	SessOverflowClose        interface{} `stat_name:"sc_rx_overflow" metric_name:"sess.overflowClose" source_type:"Rate"`
	SessTimeoutClose         interface{} `stat_name:"sc_rx_timeout" metric_name:"sess.timeoutClose" source_type:"Rate"`
	SessPipeTxnClose         interface{} `stat_name:"sc_tx_pipe" metric_name:"sess.pipeTxnClose" source_type:"Rate"`
	SessErrorTxnClose        interface{} `stat_name:"sc_tx_error" metric_name:"sess.errorTxnClose" source_type:"Rate"`
	SessEOFTxnClose          interface{} `stat_name:"sc_tx_eof" metric_name:"sess.eofTxnClose" source_type:"Rate"`
	SessBackendClose         interface{} `stat_name:"sc_resp_close" metric_name:"sess.backendClose" source_type:"Rate"`
	SessOverloadClose        interface{} `stat_name:"sc_overload" metric_name:"sess.overloadClose" source_type:"Rate"`
	SessPipeOverflowClose    interface{} `stat_name:"sc_pipe_overflow" metric_name:"sess.pipeOverflowClose" source_type:"Rate"`
	SessShortRangeClose      interface{} `stat_name:"sc_range_short" metric_name:"sess.shortRangeClose" source_type:"Rate"`
	SessReqHTTP20Close       interface{} `stat_name:"sc_req_http20" metric_name:"sess.reqHTTP20Close" source_type:"Rate"`
	SessVclFailClose         interface{} `stat_name:"sc_vcl_failure" metric_name:"sess.vclFailClose" source_type:"Rate"`
	WorkspaceDeliveryFail    interface{} `stat_name:"client_resp_500" metric_name:"workspace.deliveryFail" source_type:"Rate"`
	WorkspaceBackendOverflow interface{} `stat_name:"ws_backend_overflow" metric_name:"workspace.backendOverflow" source_type:"Rate"`
	WorkspaceClientOverflow  interface{} `stat_name:"ws_client_overflow" metric_name:"workspace.clientOverflow" source_type:"Rate"`
	WorkspaceThreadOverflow  interface{} `stat_name:"ws_thread_overflow" metric_name:"workspace.threadOverflow" source_type:"Rate"`
	WorkspaceSessionOverflow interface{} `stat_name:"ws_session_overflow" metric_name:"workspace.sessionOverflow" source_type:"Rate"`
	ShmRecords               interface{} `stat_name:"shm_records" metric_name:"shm.records" source_type:"Rate"`
	ShmWrites                interface{} `stat_name:"shm_writes" metric_name:"shm.writes" source_type:"Rate"`
	ShmFlushes               interface{} `stat_name:"shm_flushes" metric_name:"shm.flushes" source_type:"Rate"`
	ShmContentions           interface{} `stat_name:"shm_cont" metric_name:"shm.contentions" source_type:"Rate"`
	ShmCycles                interface{} `stat_name:"shm_cycles" metric_name:"shm.cycles" source_type:"Rate"`
	BackendRequests          interface{} `stat_name:"backend_req" metric_name:"backend.requests" source_type:"Rate"`
	MainVclLoaded            interface{} `stat_name:"n_vcl" metric_name:"main.vclLoaded" source_type:"Gauge"`
	MainVclAvailable         interface{} `stat_name:"n_vcl_avail" metric_name:"main.vclAvailable" source_type:"Gauge"`
	MainVclDiscarded         interface{} `stat_name:"n_vcl_discard" metric_name:"main.vclDiscarded" source_type:"Gauge"`
	MainVclFails             interface{} `stat_name:"vcl_fail" metric_name:"main.vclFails" source_type:"Rate"`
	MainBans                 interface{} `stat_name:"bans" metric_name:"main.bans" source_type:"Gauge"`
	BansCompleted            interface{} `stat_name:"bans_completed" metric_name:"bans.completed" source_type:"Gauge"`
	BansObj                  interface{} `stat_name:"bans_obj" metric_name:"bans.obj" source_type:"Gauge"`
	BansReq                  interface{} `stat_name:"bans_req" metric_name:"bans.req" source_type:"Gauge"`
	BansAdded                interface{} `stat_name:"bans_added" metric_name:"bans.added" source_type:"Rate"`
	BansDeleted              interface{} `stat_name:"bans_deleted" metric_name:"bans.deleted" source_type:"Rate"`
	BansTested               interface{} `stat_name:"bans_tested" metric_name:"bans.tested" source_type:"Rate"`
	BansLookupKilled         interface{} `stat_name:"bans_obj_killed" metric_name:"bans.lookupKilled" source_type:"Rate"`
	BansLurkerTested         interface{} `stat_name:"bans_lurker_tested" metric_name:"bans.lurkerTested" source_type:"Rate"`
	BansLookupTestsTested    interface{} `stat_name:"bans_tests_tested" metric_name:"bans.lookupTestsTested" source_type:"Rate"`
	BansLurkerTestsTested    interface{} `stat_name:"bans_lurker_tests_tested" metric_name:"bans.lurkerTestsTested" source_type:"Rate"`
	BansLurkerKilled         interface{} `stat_name:"bans_lurker_obj_killed" metric_name:"bans.lurkerKilled" source_type:"Rate"`
	BansCutoffLurkerKilled   interface{} `stat_name:"bans_lurker_obj_killed_cutoff" metric_name:"bans.cutoffLurkerKilled" source_type:"Rate"`
	BansDups                 interface{} `stat_name:"bans_dups" metric_name:"bans.dups" source_type:"Rate"`
	BansLurkerCon            interface{} `stat_name:"bans_lurker_contention" metric_name:"bans.lurkerCon" source_type:"Rate"`
	BansPersisted            interface{} `stat_name:"bans_persisted_bytes" metric_name:"bans.persistedInBytes" source_type:"Gauge"`
	BansFragmentation        interface{} `stat_name:"bans_persisted_fragmentation" metric_name:"bans.fragmentationInBytes" source_type:"Gauge"`
	MainPurgeOperations      interface{} `stat_name:"n_purges" metric_name:"main.purgeOperations" source_type:"Rate"`
	MainPurgeObjects         interface{} `stat_name:"n_obj_purged" metric_name:"main.purgeObjects" source_type:"Rate"`
	MainExpiredMailed        interface{} `stat_name:"exp_mailed" metric_name:"main.expiredMailed" source_type:"Rate"`
	MainExpiredReceived      interface{} `stat_name:"exp_received" metric_name:"main.expiredReceived" source_type:"Rate"`
	HcbNoLock                interface{} `stat_name:"hcb_nolock" metric_name:"hcb.noLock" source_type:"Rate"`
	HcbLock                  interface{} `stat_name:"hcb_lock" metric_name:"hcb.lock" source_type:"Rate"`
	HcbInsers                interface{} `stat_name:"hcb_insert" metric_name:"hcb.inserts" source_type:"Rate"`
	EsiErrors                interface{} `stat_name:"esi_errors" metric_name:"esi.errors" source_type:"Rate"`
	EsiWarnings              interface{} `stat_name:"esi_warnings" metric_name:"esi.warnings" source_type:"Rate"`
	MainVmodsLoaded          interface{} `stat_name:"vmods" metric_name:"main.vmodsLoaded" source_type:"Gauge"`
	MainGzip                 interface{} `stat_name:"n_gzip" metric_name:"main.gzip" source_type:"Rate"`
	MainGunzip               interface{} `stat_name:"n_gunzip" metric_name:"main.gunzip" source_type:"Rate"`
	MainGunzipTest           interface{} `stat_name:"n_test_gunzip" metric_name:"main.gunzipTest" source_type:"Rate"`

	locks    map[string]*lockDefinition
	mempools map[string]*mempoolDefinition
	storages map[string]*storageDefinition
}

// lockDefinition represents the data for a VarnishLockSample event
type lockDefinition struct {
	Name      string      `metric_name:"lock" source_type:"Attribute"`
	Created   interface{} `stat_name:"creat" metric_name:"lock.created" source_type:"Rate"`
	Destroyed interface{} `stat_name:"destroy" metric_name:"lock.destroyed" source_type:"Rate"`
	LockOps   interface{} `stat_name:"locks" metric_name:"lock.locks" source_type:"Rate"`
}

// mempoolDefinition represents the data for a VarnishMempoolSample event
type mempoolDefinition struct {
	Name          string      `metric_name:"memoryPool" source_type:"Attribute"`
	Live          interface{} `stat_name:"live" metric_name:"mempool.live" source_type:"Gauge"`
	Pool          interface{} `stat_name:"pool" metric_name:"mempool.pool" source_type:"Gauge"`
	ReqSize       interface{} `stat_name:"sz_wanted" metric_name:"mempool.reqSizeInBytes" source_type:"Gauge"`
	AllocatedSize interface{} `stat_name:"sz_actual" metric_name:"mempool.allocatedSizeInBytes" source_type:"Gauge"`
	Allocs        interface{} `stat_name:"allocs" metric_name:"mempool.allocs" source_type:"Rate"`
	Frees         interface{} `stat_name:"frees" metric_name:"mempool.frees" source_type:"Rate"`
	Recycles      interface{} `stat_name:"recycle" metric_name:"mempool.recycles" source_type:"Rate"`
	Timeouts      interface{} `stat_name:"timeout" metric_name:"mempool.timeouts" source_type:"Rate"`
	TooSmall      interface{} `stat_name:"toosmall" metric_name:"mempool.tooSmall" source_type:"Rate"`
	Surplus       interface{} `stat_name:"surplus" metric_name:"mempool.surplus" source_type:"Rate"`
	RanDry        interface{} `stat_name:"randry" metric_name:"mempool.ranDry" source_type:"Rate"`
}

// storageDefinition represents the data for a VarnishStorageSample event
type storageDefinition struct {
	Name            string      `metric_name:"storage" source_type:"Attribute"`
	AllocReqs       interface{} `stat_name:"c_req" metric_name:"storage.allocReqs" source_type:"Rate"`
	AllocFails      interface{} `stat_name:"c_fail" metric_name:"storage.allocFails" source_type:"Rate"`
	Alloc           interface{} `stat_name:"c_bytes" metric_name:"storage.allocInBytes" source_type:"Rate"`
	Free            interface{} `stat_name:"c_freed" metric_name:"storage.freeInBytes" source_type:"Rate"`
	AllocOustanding interface{} `stat_name:"g_alloc" metric_name:"storage.allocOustanding" source_type:"Gauge"`
	Outstanding     interface{} `stat_name:"g_bytes" metric_name:"storage.outstandingInBytes" source_type:"Gauge"`
	Available       interface{} `stat_name:"g_space" metric_name:"storage.availableInBytes" source_type:"Gauge"`
}

// backendDefinition represents the data to be set for a Backend Entity and VarnishBackendSample event
type backendDefinition struct {
	Happy           interface{} `stat_name:"happy" metric_name:"backend.happy" source_type:"Gauge"`
	ReqHeader       interface{} `stat_name:"bereq_hdrbytes" metric_name:"net.backend.reqHeaderInBytes" source_type:"Rate"`
	ReqBody         interface{} `stat_name:"bereq_bodybytes" metric_name:"net.backend.reqBodyInBytes" source_type:"Rate"`
	RespHeader      interface{} `stat_name:"beresp_hdrbytes" metric_name:"net.backend.respHeaderInBytes" source_type:"Rate"`
	RespBody        interface{} `stat_name:"beresp_bodybytes" metric_name:"net.backend.respBodyInBytes" source_type:"Rate"`
	PipeHeader      interface{} `stat_name:"pipe_hdrbytes" metric_name:"net.backend.pipeHeaderInBytes" source_type:"Rate"`
	PipeOut         interface{} `stat_name:"pipe_out" metric_name:"net.backend.pipeOutInBytes" source_type:"Rate"`
	PipeIn          interface{} `stat_name:"pipe_in" metric_name:"net.backend.pipeInInBytes" source_type:"Rate"`
	Connections     interface{} `stat_name:"conn" metric_name:"backend.connections" source_type:"Gauge"`
	Req             interface{} `stat_name:"req" metric_name:"net.backend.req" source_type:"Rate"`
	UnhealtyFetches interface{} `stat_name:"unhealty" metric_name:"backend.unhealtyFetches" source_type:"Rate"`
	BusyFetches     interface{} `stat_name:"busy" metric_name:"backend.busyFetches" source_type:"Rate"`
	ConFailed       interface{} `stat_name:"fail" metric_name:"backend.conFailed" source_type:"Rate"`
	ConNotAttempted interface{} `stat_name:"helddown" metric_name:"backend.conNotAttempted" source_type:"Rate"`
}
