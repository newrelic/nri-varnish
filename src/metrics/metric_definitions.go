package metrics

// varnishDefinition represents the data to be fed to the VarnishSample event
type varnishDefinition struct {
	MgtUptime                interface{} `stat_name:"MGT.uptime" metric_name:"mgt.uptimeInMilliseconds" source_type:"Gauge"`
	MgtChildStart            interface{} `stat_name:"child_start" metric_name:"mgt.childStart" source_type:"Prate"`
	MgtChildExit             interface{} `stat_name:"child_exit" metric_name:"mgt.childExit" source_type:"Prate"`
	MgtChildStop             interface{} `stat_name:"child_stop" metric_name:"mgt.childStop" source_type:"Prate"`
	MgtChildDied             interface{} `stat_name:"child_died" metric_name:"mgt.childDied" source_type:"Prate"`
	MgtChildDump             interface{} `stat_name:"child_dump" metric_name:"mgt.childDump" source_type:"Prate"`
	MgtChildPanic            interface{} `stat_name:"child_panic" metric_name:"mgt.childPanic" source_type:"Prate"`
	MainSumms                interface{} `stat_name:"summs" metric_name:"main.summs" source_type:"Rate"`
	MainUptime               interface{} `stat_name:"MAIN.uptime" metric_name:"main.uptimeInMilliseconds" source_type:"Gauge"`
	SessionConnections       interface{} `stat_name:"sess_conn" metric_name:"session.connections" source_type:"Prate"`
	SessionDrops             interface{} `stat_name:"sess_drop" metric_name:"session.drops" source_type:"Prate"`
	SessionFail              interface{} `stat_name:"sess_fail" metric_name:"session.fail" source_type:"Prate"`
	Net400Errors             interface{} `stat_name:"client_req_400" metric_name:"net.400Errors" source_type:"Prate"`
	Net417Errors             interface{} `stat_name:"client_req_417" metric_name:"net.417Errors" source_type:"Prate"`
	NetRequests              interface{} `stat_name:"client_req" metric_name:"net.requests" source_type:"Prate"`
	CacheHits                interface{} `stat_name:"cache_hit" metric_name:"cache.hits" source_type:"Prate"`
	CacheGraceHits           interface{} `stat_name:"cache_hit_grace" metric_name:"cache.graceHits" source_type:"Prate"`
	CachePassHits            interface{} `stat_name:"cache_hitpass" metric_name:"cache.passHits" source_type:"Prate"`
	CacheMissHits            interface{} `stat_name:"cache_hitmiss" metric_name:"cache.missHits" source_type:"Prate"`
	CacheMisses              interface{} `stat_name:"cache_miss" metric_name:"cache.misses" source_type:"Prate"`
	BackendConSuccess        interface{} `stat_name:"backend_conn" metric_name:"backend.connectionSuccess" source_type:"Prate"`
	BackendConUnHealthy      interface{} `stat_name:"backend_unhealthy" metric_name:"backend.connectionUnHealthy" source_type:"Prate"`
	BackendConBusy           interface{} `stat_name:"backend_busy" metric_name:"backend.connectionBusy" source_type:"Prate"`
	BackendConFails          interface{} `stat_name:"backend_fail" metric_name:"backend.connectionFails" source_type:"Prate"`
	BackendConReuses         interface{} `stat_name:"backend_reuse" metric_name:"backend.connectionReuses" source_type:"Prate"`
	BackendConRecycles       interface{} `stat_name:"backend_recycle" metric_name:"backend.connectionRecycles" source_type:"Prate"`
	BackendConRetries        interface{} `stat_name:"backend_retry" metric_name:"backend.connectionRetries" source_type:"Prate"`
	BackendRespUncacheable   interface{} `stat_name:"beresp_uncacheable" metric_name:"backend.respUncacheable" source_type:"Prate"`
	BackendRespShortlived    interface{} `stat_name:"beresp_shortlived" metric_name:"backend.respShortlived" source_type:"Prate"`
	FetchHead                interface{} `stat_name:"fetch_head" metric_name:"fetch.head" source_type:"Prate"`
	FetchContentLength       interface{} `stat_name:"fetch_length" metric_name:"fetch.contentLength" source_type:"Prate"`
	FetchChuncked            interface{} `stat_name:"fetch_chunked" metric_name:"fetch.chuncked" source_type:"Prate"`
	FetchEOF                 interface{} `stat_name:"fetch_eof" metric_name:"fetch.eof" source_type:"Prate"`
	FetchBad                 interface{} `stat_name:"fetch_bad" metric_name:"fetch.bad" source_type:"Prate"`
	FetchNoBody              interface{} `stat_name:"fetch_none" metric_name:"fetch.noBody" source_type:"Prate"`
	FetchNoBody1xx           interface{} `stat_name:"fetch_1xx" metric_name:"fetch.noBody1xx" source_type:"Prate"`
	FetchNoBody204           interface{} `stat_name:"fetch_204" metric_name:"fetch.noBody204" source_type:"Prate"`
	FetchNoBody304           interface{} `stat_name:"fetch_304" metric_name:"fetch.noBody304" source_type:"Prate"`
	FetchFailed              interface{} `stat_name:"fetch_failed" metric_name:"fetch.failed" source_type:"Prate"`
	FetchNoThreadFail        interface{} `stat_name:"fetch_no_thread" metric_name:"fetch.noThreadFail" source_type:"Prate"`
	MainPools                interface{} `stat_name:"pools" metric_name:"main.pools" source_type:"Gauge"`
	MainThreads              interface{} `stat_name:"threads" metric_name:"main.threads" source_type:"Gauge"`
	MainThreadsLimited       interface{} `stat_name:"threads_limited" metric_name:"main.threadsLimited" source_type:"Prate"`
	MainThreadsCreated       interface{} `stat_name:"threads_created" metric_name:"main.threadsCreated" source_type:"Prate"`
	MainThreadsDestroyed     interface{} `stat_name:"threads_destroyed" metric_name:"main.threadsDestroyed" source_type:"Prate"`
	MainThreadsFailed        interface{} `stat_name:"threads_failed" metric_name:"main.threadsFailed" source_type:"Prate"`
	MainSessQueueLength      interface{} `stat_name:"thread_queue_len" metric_name:"main.sessQueueLength" source_type:"Gauge"`
	MainBusySleep            interface{} `stat_name:"busy_sleep" metric_name:"main.busySleep" source_type:"Prate"`
	MainBusyWakeup           interface{} `stat_name:"busy_wakeup" metric_name:"main.busyWakeup" source_type:"Prate"`
	MainBusyKilled           interface{} `stat_name:"busy_killed" metric_name:"main.busyKilled" source_type:"Prate"`
	SessQueued               interface{} `stat_name:"sess_queued" metric_name:"sess.queued" source_type:"Prate"`
	SessDropped              interface{} `stat_name:"sess_dropped" metric_name:"sess.dropped" source_type:"Prate"`
	MainReqDropped           interface{} `stat_name:"req_dropped" metric_name:"main.reqDropped" source_type:"Rate"`
	MainObjects              interface{} `stat_name:"n_object" metric_name:"main.objects" source_type:"Gauge"`
	MainUnresurrectedObjects interface{} `stat_name:"n_vampireobject" metric_name:"main.unresurrectedObjects" source_type:"Gauge"`
	MainObjectcores          interface{} `stat_name:"n_objectcore" metric_name:"main.objectcores" source_type:"Gauge"`
	MainObjectheads          interface{} `stat_name:"n_objecthead" metric_name:"main.objectheads" source_type:"Gauge"`
	MainBackends             interface{} `stat_name:"n_backend" metric_name:"main.backends" source_type:"Gauge"`
	MainExpired              interface{} `stat_name:"n_expired" metric_name:"main.expired" source_type:"Gauge"`
	MainPipe                 interface{} `stat_name:"n_pipe" metric_name:"main.pipe" source_type:"Rate"`
	MainPipeLimited          interface{} `stat_name:"pipe_limited" metric_name:"main.pipe_limited" source_type:"Rate"`
	LruNuked                 interface{} `stat_name:"n_lru_nuked" metric_name:"lru.nuked" source_type:"Prate"`
	LruMoved                 interface{} `stat_name:"n_lru_moved" metric_name:"lru.moved" source_type:"Prate"`
	LruLimited               interface{} `stat_name:"n_lru_limited" metric_name:"lru.limited" source_type:"Prate"`
	NetHTTPOverflow          interface{} `stat_name:"losthdr" metric_name:"net.httpOverflow" source_type:"Prate"`
	MainSessions             interface{} `stat_name:"s_sess" metric_name:"main.sessions" source_type:"Prate"`
	MainPipeSessions         interface{} `stat_name:"s_pipe" metric_name:"main.pipeSessions" source_type:"Prate"`
	MainPassedReq            interface{} `stat_name:"s_pass" metric_name:"main.passedRequests" source_type:"Prate"`
	BackendFetches           interface{} `stat_name:"s_fetch" metric_name:"backend.fetches" source_type:"Prate"`
	MainSyntheticResp        interface{} `stat_name:"s_synth" metric_name:"main.syntheticResponses" source_type:"Prate"`
	NetReqHeader             interface{} `stat_name:"s_req_hdrbytes" metric_name:"net.request.headerInBytes" source_type:"Prate"`
	NetReqBody               interface{} `stat_name:"s_req_bodybytes" metric_name:"net.request.bodyInBytes" source_type:"Prate"`
	NetRespHeader            interface{} `stat_name:"s_resp_hdrbytes" metric_name:"net.response.headerInBytes" source_type:"Prate"`
	NetRespBody              interface{} `stat_name:"s_resp_bodybytes" metric_name:"net.response.bodyInBytes" source_type:"Prate"`
	NetPipereqHeader         interface{} `stat_name:"s_pipe_hdrbytes" metric_name:"net.pipereq.headerInBytes" source_type:"Prate"`
	NetPipeIn                interface{} `stat_name:"s_pipe_in" metric_name:"net.pipe.inInBytes" source_type:"Prate"`
	NetPipeOut               interface{} `stat_name:"s_pipe_out" metric_name:"net.pipe.outInBytes" source_type:"Prate"`
	SessClosed               interface{} `stat_name:"sess_closed" metric_name:"sess.closed" source_type:"Prate"`
	SessClosedError          interface{} `stat_name:"sess_closed_err" metric_name:"sess.closedError" source_type:"Prate"`
	SessReadAhead            interface{} `stat_name:"sess_readahead" metric_name:"sess.readAhead" source_type:"Prate"`
	SessHerd                 interface{} `stat_name:"sess_herd" metric_name:"sess.herd" source_type:"Prate"`
	SessFailEconnaborted     interface{} `stat_name:"sess_fail_econnaborted" metric_name:"sess.sessFailEconnaborted" source_type:"Prate"`
	SessFailEmfile           interface{} `stat_name:"sess_fail_emfile" metric_name:"sess.sessFailEmfile" source_type:"Prate"`
	SessFailEbadf            interface{} `stat_name:"sess_fail_ebadf" metric_name:"sess.sessFailEbadf" source_type:"Prate"`
	SessFailEnomem           interface{} `stat_name:"sess_fail_enomem" metric_name:"sess.sessFailEnomem" source_type:"Prate"`
	SessFailEintr            interface{} `stat_name:"sess_fail_eintr" metric_name:"sess.sessFailEintr" source_type:"Prate"`
	SessFailOther            interface{} `stat_name:"sess_fail_other" metric_name:"sess.sessFailOther" source_type:"Prate"`
	SessClientClose          interface{} `stat_name:"sc_rem_close" metric_name:"sess.clientClose" source_type:"Prate"`
	SessClientReqClose       interface{} `stat_name:"sc_req_close" metric_name:"sess.clientReqClose" source_type:"Prate"`
	SessReqHTTP10Close       interface{} `stat_name:"sc_req_http10" metric_name:"sess.requestHTTP10Close" source_type:"Prate"`
	SessBadClose             interface{} `stat_name:"sc_rx_bad" metric_name:"sess.badClose" source_type:"Prate"`
	SessBodyFailClose        interface{} `stat_name:"sc_rx_body" metric_name:"sess.bodyFailClose" source_type:"Prate"`
	SessJunkClose            interface{} `stat_name:"sc_rx_junk" metric_name:"sess.junkClose" source_type:"Prate"`
	SessOverflowClose        interface{} `stat_name:"sc_rx_overflow" metric_name:"sess.overflowClose" source_type:"Prate"`
	SessTimeoutClose         interface{} `stat_name:"sc_rx_timeout" metric_name:"sess.timeoutClose" source_type:"Prate"`
	SessCloseIdle            interface{} `stat_name:"sc_rx_close_idle" metric_name:"sess.closeIdle" source_type:"Prate"`
	SessPipeTxnClose         interface{} `stat_name:"sc_tx_pipe" metric_name:"sess.pipeTxnClose" source_type:"Prate"`
	SessErrorTxnClose        interface{} `stat_name:"sc_tx_error" metric_name:"sess.errorTxnClose" source_type:"Prate"`
	SessEOFTxnClose          interface{} `stat_name:"sc_tx_eof" metric_name:"sess.eofTxnClose" source_type:"Prate"`
	SessBackendClose         interface{} `stat_name:"sc_resp_close" metric_name:"sess.backendClose" source_type:"Prate"`
	SessOverloadClose        interface{} `stat_name:"sc_overload" metric_name:"sess.overloadClose" source_type:"Prate"`
	SessPipeOverflowClose    interface{} `stat_name:"sc_pipe_overflow" metric_name:"sess.pipeOverflowClose" source_type:"Prate"`
	SessShortRangeClose      interface{} `stat_name:"sc_range_short" metric_name:"sess.shortRangeClose" source_type:"Prate"`
	SessReqHTTP20Close       interface{} `stat_name:"sc_req_http20" metric_name:"sess.requestHTTP20Close" source_type:"Prate"`
	SessVclFailClose         interface{} `stat_name:"sc_vcl_failure" metric_name:"sess.vclFailClose" source_type:"Prate"`
	WorkspaceDeliveryFail    interface{} `stat_name:"client_resp_500" metric_name:"workspace.deliveryFail" source_type:"Prate"`
	WorkspaceBackendOverflow interface{} `stat_name:"ws_backend_overflow" metric_name:"workspace.backendOverflow" source_type:"Prate"`
	WorkspaceClientOverflow  interface{} `stat_name:"ws_client_overflow" metric_name:"workspace.clientOverflow" source_type:"Prate"`
	WorkspaceThreadOverflow  interface{} `stat_name:"ws_thread_overflow" metric_name:"workspace.threadOverflow" source_type:"Prate"`
	WorkspaceSessionOverflow interface{} `stat_name:"ws_session_overflow" metric_name:"workspace.sessionOverflow" source_type:"Prate"`
	ShmRecords               interface{} `stat_name:"shm_records" metric_name:"shm.records" source_type:"Prate"`
	ShmWrites                interface{} `stat_name:"shm_writes" metric_name:"shm.writes" source_type:"Prate"`
	ShmFlushes               interface{} `stat_name:"shm_flushes" metric_name:"shm.flushes" source_type:"Prate"`
	ShmContentions           interface{} `stat_name:"shm_cont" metric_name:"shm.contentions" source_type:"Prate"`
	ShmCycles                interface{} `stat_name:"shm_cycles" metric_name:"shm.cycles" source_type:"Prate"`
	BackendRequests          interface{} `stat_name:"backend_req" metric_name:"backend.requests" source_type:"Prate"`
	MainVclLoaded            interface{} `stat_name:"n_vcl" metric_name:"main.vclLoaded" source_type:"Gauge"`
	MainVclAvailable         interface{} `stat_name:"n_vcl_avail" metric_name:"main.vclAvailable" source_type:"Gauge"`
	MainVclDiscarded         interface{} `stat_name:"n_vcl_discard" metric_name:"main.vclDiscarded" source_type:"Gauge"`
	MainVclFails             interface{} `stat_name:"vcl_fail" metric_name:"main.vclFails" source_type:"Prate"`
	MainBans                 interface{} `stat_name:"bans" metric_name:"main.bans" source_type:"Gauge"`
	BansCompleted            interface{} `stat_name:"bans_completed" metric_name:"bans.completed" source_type:"Gauge"`
	BansObj                  interface{} `stat_name:"bans_obj" metric_name:"bans.obj" source_type:"Gauge"`
	BansReq                  interface{} `stat_name:"bans_req" metric_name:"bans.req" source_type:"Gauge"`
	BansAdded                interface{} `stat_name:"bans_added" metric_name:"bans.added" source_type:"Prate"`
	BansDeleted              interface{} `stat_name:"bans_deleted" metric_name:"bans.deleted" source_type:"Prate"`
	BansTested               interface{} `stat_name:"bans_tested" metric_name:"bans.tested" source_type:"Prate"`
	BansLookupKilled         interface{} `stat_name:"bans_obj_killed" metric_name:"bans.lookupKilled" source_type:"Prate"`
	BansLurkerTested         interface{} `stat_name:"bans_lurker_tested" metric_name:"bans.lurkerTested" source_type:"Prate"`
	BansLookupTestsTested    interface{} `stat_name:"bans_tests_tested" metric_name:"bans.lookupTestsTested" source_type:"Prate"`
	BansLurkerTestsTested    interface{} `stat_name:"bans_lurker_tests_tested" metric_name:"bans.lurkerTestsTested" source_type:"Prate"`
	BansLurkerKilled         interface{} `stat_name:"bans_lurker_obj_killed" metric_name:"bans.lurkerKilled" source_type:"Prate"`
	BansCutoffLurkerKilled   interface{} `stat_name:"bans_lurker_obj_killed_cutoff" metric_name:"bans.cutoffLurkerKilled" source_type:"Prate"`
	BansDups                 interface{} `stat_name:"bans_dups" metric_name:"bans.dups" source_type:"Prate"`
	BansLurkerCon            interface{} `stat_name:"bans_lurker_contention" metric_name:"bans.lurkerCon" source_type:"Prate"`
	BansPersisted            interface{} `stat_name:"bans_persisted_bytes" metric_name:"bans.persistedInBytes" source_type:"Gauge"`
	BansFragmentation        interface{} `stat_name:"bans_persisted_fragmentation" metric_name:"bans.fragmentationInBytes" source_type:"Gauge"`
	MainPurgeOperations      interface{} `stat_name:"n_purges" metric_name:"main.purgeOperations" source_type:"Gauge"`
	MainPurgeObjects         interface{} `stat_name:"n_obj_purged" metric_name:"main.purgeObjects" source_type:"Gauge"`
	MainExpiredMailed        interface{} `stat_name:"exp_mailed" metric_name:"main.expiredMailed" source_type:"Prate"`
	MainExpiredReceived      interface{} `stat_name:"exp_received" metric_name:"main.expiredReceived" source_type:"Prate"`
	HcbNoLock                interface{} `stat_name:"hcb_nolock" metric_name:"hcb.noLock" source_type:"Prate"`
	HcbLock                  interface{} `stat_name:"hcb_lock" metric_name:"hcb.lock" source_type:"Prate"`
	HcbInsers                interface{} `stat_name:"hcb_insert" metric_name:"hcb.inserts" source_type:"Prate"`
	EsiErrors                interface{} `stat_name:"esi_errors" metric_name:"esi.errors" source_type:"Prate"`
	EsiWarnings              interface{} `stat_name:"esi_warnings" metric_name:"esi.warnings" source_type:"Prate"`
	MainVmodsLoaded          interface{} `stat_name:"vmods" metric_name:"main.vmodsLoaded" source_type:"Gauge"`
	MainGzip                 interface{} `stat_name:"n_gzip" metric_name:"main.gzip" source_type:"Prate"`
	MainGunzip               interface{} `stat_name:"n_gunzip" metric_name:"main.gunzip" source_type:"Prate"`
	MainGunzipTest           interface{} `stat_name:"n_test_gunzip" metric_name:"main.gunzipTest" source_type:"Prate"`

	locks          map[string]*lockDefinition
	mempools       map[string]*mempoolDefinition
	storages       map[string]*storageDefinition
	book           map[string]*bookDefinition
	store          map[string]*storeDefinition
	massiveStorage map[string]*massiveStorageDefinition
}

// lockDefinition represents the data for a VarnishLockSample event
type lockDefinition struct {
	Created    interface{} `stat_name:"creat" metric_name:"lock.created" source_type:"Prate"`
	Destroyed  interface{} `stat_name:"destroy" metric_name:"lock.destroyed" source_type:"Prate"`
	LockOps    interface{} `stat_name:"locks" metric_name:"lock.locks" source_type:"Prate"`
	DbgBusy    interface{} `stat_name:"dbg_busy" metric_name:"lock.dbg_busy" source_type:"Prate"`
	DbgTryFail interface{} `stat_name:"dbg_try_fail" metric_name:"lock.dbg_try_fail" source_type:"Prate"`
}

// mempoolDefinition represents the data for a VarnishMempoolSample event
type mempoolDefinition struct {
	Live          interface{} `stat_name:"live" metric_name:"mempool.live" source_type:"Gauge"`
	Pool          interface{} `stat_name:"pool" metric_name:"mempool.pool" source_type:"Gauge"`
	ReqSize       interface{} `stat_name:"sz_wanted" metric_name:"mempool.requestSizeInBytes" source_type:"Gauge"`
	AllocatedSize interface{} `stat_name:"sz_actual" metric_name:"mempool.allocatedSizeInBytes" source_type:"Gauge"`
	Allocs        interface{} `stat_name:"allocs" metric_name:"mempool.allocs" source_type:"Prate"`
	Frees         interface{} `stat_name:"frees" metric_name:"mempool.frees" source_type:"Prate"`
	Recycles      interface{} `stat_name:"recycle" metric_name:"mempool.recycles" source_type:"Prate"`
	Timeouts      interface{} `stat_name:"timeout" metric_name:"mempool.timeouts" source_type:"Prate"`
	TooSmall      interface{} `stat_name:"toosmall" metric_name:"mempool.tooSmall" source_type:"Prate"`
	Surplus       interface{} `stat_name:"surplus" metric_name:"mempool.surplus" source_type:"Prate"`
	RanDry        interface{} `stat_name:"randry" metric_name:"mempool.ranDry" source_type:"Prate"`
}

// bookDefinition represents the data for a VarnishBookSample event
type bookDefinition struct {
	Alloc        interface{} `stat_name:"g_bytes" metric_name:"book.allocInBytes" source_type:"Gauge"`
	Available    interface{} `stat_name:"g_space" metric_name:"book.availableInBytes" source_type:"Gauge"`
	ThreadQueued interface{} `stat_name:"c_waterlevel_queue" metric_name:"book.threadQueued" source_type:"Prate"`
	PurgeObjects interface{} `stat_name:"c_waterlevel_purge" metric_name:"book.purgeObjects" source_type:"Prate"`
}

type storeDefinition struct {
	Objects      interface{} `stat_name:"g_objects" metric_name:"store.numOfObjects" source_type:"Gauge"`
	AioQueue     interface{} `stat_name:"c_aio_queue" metric_name:"store.numOfAioQueue" source_type:"Prate"`
	AioBytes     interface{} `stat_name:"c_aio_finished_bytes" metric_name:"store.numOfAioBytes" source_type:"Prate"`
	AioRead      interface{} `stat_name:"c_aio_finished_read" metric_name:"store.numOfAioRead" source_type:"Prate"`
	AioWrite     interface{} `stat_name:"c_aio_finished_write" metric_name:"store.numOfAioWrite" source_type:"Prate"`
	ThreadQueue  interface{} `stat_name:"c_waterlevel_queue" metric_name:"store.threadQueue" source_type:"Prate"`
	PurgeObjects interface{} `stat_name:"c_waterlevel_purge" metric_name:"store.purgeObjects" source_type:"Prate"`
	YkeysReg     interface{} `stat_name:"g_ykey_keys" metric_name:"store.numOfYkeysReg" source_type:"Gauge"`
	YkeysPurged  interface{} `stat_name:"c_ykey_purged" metric_name:"store.numOfYkeysPurged" source_type:"Prate"`
}

// storageDefinition represents the data for a VarnishStorageSample event
type storageDefinition struct {
	AllocReqs       interface{} `stat_name:"c_req" metric_name:"storage.allocReqs" source_type:"Prate"`
	AllocFails      interface{} `stat_name:"c_fail" metric_name:"storage.allocFails" source_type:"Prate"`
	Alloc           interface{} `stat_name:"c_bytes" metric_name:"storage.allocInBytes" source_type:"Prate"`
	Free            interface{} `stat_name:"c_freed" metric_name:"storage.freeInBytes" source_type:"Prate"`
	AllocOustanding interface{} `stat_name:"g_alloc" metric_name:"storage.allocOustanding" source_type:"Gauge"`
	Outstanding     interface{} `stat_name:"g_bytes" metric_name:"storage.outstandingInBytes" source_type:"Gauge"`
	Available       interface{} `stat_name:"g_space" metric_name:"storage.availableInBytes" source_type:"Gauge"`
}

type massiveStorageDefinition struct {
	YkeysReg    interface{} `stat_name:"g_ykey_keys" metric_name:"massiveStorage.numOfYkeysReg" source_type:"Gauge"`
	YkeysPurged interface{} `stat_name:"c_ykey_purged" metric_name:"massiveStorage.numOfYkeysPurged" source_type:"Prate"`
}

// backendDefinition represents the data to be set for a Backend Entity and VarnishBackendSample event
type backendDefinition struct {
	Happy                  interface{} `stat_name:"happy" metric_name:"backend.happy" source_type:"Gauge"`
	ReqHeader              interface{} `stat_name:"bereq_hdrbytes" metric_name:"net.backend.requestHeaderInBytes" source_type:"Prate"`
	ReqBody                interface{} `stat_name:"bereq_bodybytes" metric_name:"net.backend.requestBodyInBytes" source_type:"Prate"`
	RespHeader             interface{} `stat_name:"beresp_hdrbytes" metric_name:"net.backend.responseHeaderInBytes" source_type:"Prate"`
	RespBody               interface{} `stat_name:"beresp_bodybytes" metric_name:"net.backend.responseBodyInBytes" source_type:"Prate"`
	PipeHeader             interface{} `stat_name:"pipe_hdrbytes" metric_name:"net.backend.pipeHeaderInBytes" source_type:"Prate"`
	PipeOut                interface{} `stat_name:"pipe_out" metric_name:"net.backend.pipeOutInBytes" source_type:"Prate"`
	PipeIn                 interface{} `stat_name:"pipe_in" metric_name:"net.backend.pipeInInBytes" source_type:"Prate"`
	Connections            interface{} `stat_name:"conn" metric_name:"backend.connections" source_type:"Gauge"`
	Req                    interface{} `stat_name:"req" metric_name:"net.backend.requests" source_type:"Prate"`
	UnhealtyFetches        interface{} `stat_name:"unhealthy" metric_name:"backend.unhealthyFetches" source_type:"Prate"`
	BusyFetches            interface{} `stat_name:"busy" metric_name:"backend.busyFetches" source_type:"Prate"`
	ConFailed              interface{} `stat_name:"fail" metric_name:"backend.connectionsFailed" source_type:"Prate"`
	ConFailedEaccess       interface{} `stat_name:"fail_eacces" metric_name:"backend.connectionsFailedEacces" source_type:"Prate"`
	ConFailedEconnrefused  interface{} `stat_name:"fail_econnrefused" metric_name:"backend.connectionsFailedEconnrefused" source_type:"Prate"`
	ConFailedEtimedout     interface{} `stat_name:"fail_etimedout" metric_name:"backend.connectionsFailedEtimedout" source_type:"Prate"`
	ConFailedEaddrnotavail interface{} `stat_name:"fail_eaddrnotavail" metric_name:"backend.connectionsFailedEaddrnotavail" source_type:"Prate"`
	ConFailedEnetunreach   interface{} `stat_name:"fail_enetunreach" metric_name:"backend.connectionsFailedEnetunreach" source_type:"Prate"`
	ConFailedOther         interface{} `stat_name:"fail_other" metric_name:"backend.connectionsFailedOther" source_type:"Prate"`
	ConNotAttempted        interface{} `stat_name:"helddown" metric_name:"backend.connectionsNotAttempted" source_type:"Prate"`
}
