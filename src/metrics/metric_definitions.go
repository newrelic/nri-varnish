package metrics

// varnishDefinition represents the data to be fed to the VarnishSample event
type varnishDefinition struct {
	mgtUptime                interface{} `metric_name:"mgt.uptimeInMilliseconds" source_type:"Rate"`
	mgtChildStart            interface{} `metric_name:"mgt.childStart" source_type:"Rate"`
	mgtChildExit             interface{} `metric_name:"mgt.childExit" source_type:"Rate"`
	mgtChildStop             interface{} `metric_name:"mgt.childStop" source_type:"Rate"`
	mgtChildDied             interface{} `metric_name:"mgt.childDied" source_type:"Rate"`
	mgtChildDump             interface{} `metric_name:"mgt.childDump" source_type:"Rate"`
	mgtChildPanic            interface{} `metric_name:"mgt.childPanic" source_type:"Rate"`
	mainSumms                interface{} `metric_name:"main.summs" source_type:"Rate"`
	mainUptime               interface{} `metric_name:"main.uptimeInMilliseconds" source_type:"Rate"`
	sessionConnections       interface{} `metric_name:"session.connections" source_type:"Rate"`
	sessionDrops             interface{} `metric_name:"session.drops" source_type:"Rate"`
	sessionFail              interface{} `metric_name:"session.fail" source_type:"Rate"`
	net400Errors             interface{} `metric_name:"net.400Errors" source_type:"Rate"`
	net417Errors             interface{} `metric_name:"net.417Errors" source_type:"Rate"`
	netRequests              interface{} `metric_name:"net.requests" source_type:"Rate"`
	cacheHits                interface{} `metric_name:"cache.hits" source_type:"Rate"`
	cacheGraceHits           interface{} `metric_name:"cache.graceHits" source_type:"Rate"`
	cachePassHits            interface{} `metric_name:"cache.passHits" source_type:"Rate"`
	cacheMissHits            interface{} `metric_name:"cache.missHits" source_type:"Rate"`
	cacheMisses              interface{} `metric_name:"cache.misses" source_type:"Rate"`
	backendConSuccess        interface{} `metric_name:"backend.conSuccess" source_type:"Rate"`
	backendConUnHealthy      interface{} `metric_name:"backend.conUnHealthy" source_type:"Rate"`
	backendConBusy           interface{} `metric_name:"backend.conBusy" source_type:"Rate"`
	backendConFails          interface{} `metric_name:"backend.conFails" source_type:"Rate"`
	backendConReuses         interface{} `metric_name:"backend.conReuses" source_type:"Rate"`
	backendConRecycles       interface{} `metric_name:"backend.conRecycles" source_type:"Rate"`
	backendConRetries        interface{} `metric_name:"backend.conRetries" source_type:"Rate"`
	fetchHead                interface{} `metric_name:"fetch.head" source_type:"Rate"`
	fetchContentLength       interface{} `metric_name:"fetch.contentLength" source_type:"Rate"`
	fetchChuncked            interface{} `metric_name:"fetch.chuncked" source_type:"Rate"`
	fetchEOF                 interface{} `metric_name:"fetch.eof" source_type:"Rate"`
	fetchBad                 interface{} `metric_name:"fetch.bad" source_type:"Rate"`
	fetchNoBody              interface{} `metric_name:"fetch.noBody" source_type:"Rate"`
	fetchNoBody1xx           interface{} `metric_name:"fetch.noBody1xx" source_type:"Rate"`
	fetchNoBody204           interface{} `metric_name:"fetch.noBody204" source_type:"Rate"`
	fetchNoBody304           interface{} `metric_name:"fetch.noBody304" source_type:"Rate"`
	fetchFailed              interface{} `metric_name:"fetch.failed" source_type:"Rate"`
	fetchNoThreadFail        interface{} `metric_name:"fetch.noThreadFail" source_type:"Rate"`
	mainPools                interface{} `metric_name:"main.pools" source_type:"Gauge"`
	mainThreads              interface{} `metric_name:"main.threads" source_type:"Gauge"`
	mainThreadsLimited       interface{} `metric_name:"main.threadsLimited" source_type:"Rate"`
	mainThreadsCreated       interface{} `metric_name:"main.threadsCreated" source_type:"Rate"`
	mainThreadsDestroyed     interface{} `metric_name:"main.threadsDestroyed" source_type:"Rate"`
	mainThreadsFailed        interface{} `metric_name:"main.threadsFailed" source_type:"Rate"`
	mainSessQueueLength      interface{} `metric_name:"main.sessQueueLength" source_type:"Gauge"`
	mainBusySleep            interface{} `metric_name:"main.busySleep" source_type:"Rate"`
	mainBusyWakeup           interface{} `metric_name:"main.busyWakeup" source_type:"Rate"`
	mainBusyKilled           interface{} `metric_name:"main.busyKilled" source_type:"Rate"`
	sessQueued               interface{} `metric_name:"sess.queued" source_type:"Rate"`
	sessDropped              interface{} `metric_name:"sess.dropped" source_type:"Rate"`
	mainReqDropped           interface{} `metric_name:"main.reqDropped" source_type:"Rate"`
	mainObjects              interface{} `metric_name:"main.objects" source_type:"Gauge"`
	mainUnresurrectedObjects interface{} `metric_name:"main.unresurrectedObjects" source_type:"Gauge"`
	mainObjectcores          interface{} `metric_name:"main.objectcores" source_type:"Gauge"`
	mainObjectheads          interface{} `metric_name:"main.objectheads" source_type:"Gauge"`
	mainBackends             interface{} `metric_name:"main.backends" source_type:"Gauge"`
	mainExpired              interface{} `metric_name:"main.expired" source_type:"Rate"`
	lruNuked                 interface{} `metric_name:"lru.nuked" source_type:"Rate"`
	lruMoved                 interface{} `metric_name:"lru.moved" source_type:"Rate"`
	lruLimited               interface{} `metric_name:"lru.limited" source_type:"Rate"`
	netHTTPOverflow          interface{} `metric_name:"net.httpOverflow" source_type:"Rate"`
	mainSessions             interface{} `metric_name:"main.sessions" source_type:"Rate"`
	mainPipeSessions         interface{} `metric_name:"main.pipeSessions" source_type:"Rate"`
	mainPassedReq            interface{} `metric_name:"main.passedReq" source_type:"Rate"`
	backendFetches           interface{} `metric_name:"backend.fetches" source_type:"Rate"`
	mainSyntheticResp        interface{} `metric_name:"main.syntheticResp" source_type:"Rate"`
	netReqHeader             interface{} `metric_name:"net.req.headerInBytes" source_type:"Rate"`
	netReqBody               interface{} `metric_name:"net.req.bodyInBytes" source_type:"Rate"`
	netRespHeader            interface{} `metric_name:"net.resp.headerInBytes" source_type:"Rate"`
	netRespBody              interface{} `metric_name:"net.resp.bodyInBytes" source_type:"Rate"`
	netPipereqHeader         interface{} `metric_name:"net.pipereq.headerInBytes" source_type:"Rate"`
	netPipeIn                interface{} `metric_name:"net.pipe.inInBytes" source_type:"Rate"`
	netPipeOut               interface{} `metric_name:"net.pipe.outInBytes" source_type:"Rate"`
	sessClosed               interface{} `metric_name:"sess.closed" source_type:"Rate"`
	sessClosedError          interface{} `metric_name:"sess.closedError" source_type:"Rate"`
	sessReadAhead            interface{} `metric_name:"sess.readAhead" source_type:"Rate"`
	sessHerd                 interface{} `metric_name:"sess.herd" source_type:"Rate"`
	sessClientClose          interface{} `metric_name:"sess.clientClose" source_type:"Rate"`
	sessClientReqClose       interface{} `metric_name:"sess.clientReqClose" source_type:"Rate"`
	sessReqHTTP10Close       interface{} `metric_name:"sess.reqHTTP10Close" source_type:"Rate"`
	sessBadClose             interface{} `metric_name:"sess.badClose" source_type:"Rate"`
	sessBodyFailClose        interface{} `metric_name:"sess.bodyFailClose" source_type:"Rate"`
	sessJunkClose            interface{} `metric_name:"sess.junkClose" source_type:"Rate"`
	sessOverflowClose        interface{} `metric_name:"sess.overflowClose" source_type:"Rate"`
	sessTimeoutClose         interface{} `metric_name:"sess.timeoutClose" source_type:"Rate"`
	sessPipeTxnClose         interface{} `metric_name:"sess.pipeTxnClose" source_type:"Rate"`
	sessErrorTxnClose        interface{} `metric_name:"sess.errorTxnClose" source_type:"Rate"`
	sessEOFTxnClose          interface{} `metric_name:"sess.eofTxnClose" source_type:"Rate"`
	sessBackendClose         interface{} `metric_name:"sess.backendClose" source_type:"Rate"`
	sessOverloadClose        interface{} `metric_name:"sess.overloadClose" source_type:"Rate"`
	sessPipeOverflowClose    interface{} `metric_name:"sess.pipeOverflowClose" source_type:"Rate"`
	sessShortRangeClose      interface{} `metric_name:"sess.shortRangeClose" source_type:"Rate"`
	sessReqHTTP20Close       interface{} `metric_name:"sess.reqHTTP20Close" source_type:"Rate"`
	sessVclFailClose         interface{} `metric_name:"sess.vclFailClose" source_type:"Rate"`
	workspaceDeliveryFail    interface{} `metric_name:"workspace.deliveryFail" source_type:"Rate"`
	workspaceBackendOverflow interface{} `metric_name:"workspace.backendOverflow" source_type:"Rate"`
	workspaceClientOverflow  interface{} `metric_name:"workspace.clientOverflow" source_type:"Rate"`
	workspaceThreadOverflow  interface{} `metric_name:"workspace.threadOverflow" source_type:"Rate"`
	workspaceSessionOverflow interface{} `metric_name:"workspace.sessionOverflow" source_type:"Rate"`
	shmRecords               interface{} `metric_name:"shm.records" source_type:"Rate"`
	shmWrites                interface{} `metric_name:"shm.writes" source_type:"Rate"`
	shmFlushes               interface{} `metric_name:"shm.flushes" source_type:"Rate"`
	shmContentions           interface{} `metric_name:"shm.contentions" source_type:"Rate"`
	shmCycles                interface{} `metric_name:"shm.cycles" source_type:"Rate"`
	backendRequests          interface{} `metric_name:"backend.requests" source_type:"Rate"`
	mainVclLoaded            interface{} `metric_name:"main.vclLoaded" source_type:"Gauge"`
	mainVclAvailable         interface{} `metric_name:"main.vclAvailable" source_type:"Gauge"`
	mainVclDiscarded         interface{} `metric_name:"main.vclDiscarded" source_type:"Gauge"`
	mainVclFails             interface{} `metric_name:"main.vclFails" source_type:"Rate"`
	mainBans                 interface{} `metric_name:"main.bans" source_type:"Gauge"`
	bansCompleted            interface{} `metric_name:"bans.completed" source_type:"Gauge"`
	bansObj                  interface{} `metric_name:"bans.obj" source_type:"Gauge"`
	bansReq                  interface{} `metric_name:"bans.req" source_type:"Gauge"`
	bansAdded                interface{} `metric_name:"bans.added" source_type:"Rate"`
	bansDeleted              interface{} `metric_name:"bans.deleted" source_type:"Rate"`
	bansTested               interface{} `metric_name:"bans.tested" source_type:"Rate"`
	bansLookupKilled         interface{} `metric_name:"bans.lookupKilled" source_type:"Rate"`
	bansLurkerTested         interface{} `metric_name:"bans.lurkerTested" source_type:"Rate"`
	bansLookupTestsTested    interface{} `metric_name:"bans.lookupTestsTested" source_type:"Rate"`
	bansLurkerTestsTested    interface{} `metric_name:"bans.lurkerTestsTested" source_type:"Rate"`
	bansLurkerKilled         interface{} `metric_name:"bans.lurkerKilled" source_type:"Rate"`
	bansCutoffLurkerKilled   interface{} `metric_name:"bans.cutoffLurkerKilled" source_type:"Rate"`
	bansDups                 interface{} `metric_name:"bans.dups" source_type:"Rate"`
	bansLurkerCon            interface{} `metric_name:"bans.lurkerCon" source_type:"Rate"`
	bansPersisted            interface{} `metric_name:"bans.persistedInBytes" source_type:"Gauge"`
	bansFragmentation        interface{} `metric_name:"bans.fragmentationInBytes" source_type:"Gauge"`
	mainPurgeOperations      interface{} `metric_name:"main.purgeOperations" source_type:"Rate"`
	mainPurgeObjects         interface{} `metric_name:"main.purgeObjects" source_type:"Rate"`
	mainExpiredMailed        interface{} `metric_name:"main.expiredMailed" source_type:"Rate"`
	mainExpiredReceived      interface{} `metric_name:"main.expiredReceived" source_type:"Rate"`
	hcbNoLock                interface{} `metric_name:"hcb.noLock" source_type:"Rate"`
	hcbLock                  interface{} `metric_name:"hcb.lock" source_type:"Rate"`
	hcbInsers                interface{} `metric_name:"hcb.inserts" source_type:"Rate"`
	esiErrors                interface{} `metric_name:"esi.errors" source_type:"Rate"`
	esiWarnings              interface{} `metric_name:"esi.warnings" source_type:"Rate"`
	mainVmodsLoaded          interface{} `metric_name:"main.vmodsLoaded" source_type:"Gauge"`
	mainGzip                 interface{} `metric_name:"main.gzip" source_type:"Rate"`
	mainGunzip               interface{} `metric_name:"main.gunzip" source_type:"Rate"`
	mainGunzipTest           interface{} `metric_name:"main.gunzipTest" source_type:"Rate"`

	locks    map[string]*lockDefinition
	mempools map[string]*mempoolDefinition
	storages map[string]*storageDefinition
}

// lockDefinition represents the data for a VarnishLockSample event
type lockDefinition struct {
	name      string      `metric_name:"lock" source_type:"Attribute"`
	created   interface{} `metric_name:"lock.created" source_type:"Rate"`
	destroyed interface{} `metric_name:"lock.destroyed" source_type:"Rate"`
	lockOps   interface{} `metric_name:"lock.locks" source_type:"Rate"`
}

// mempoolDefinition represents the data for a VarnishMempoolSample event
type mempoolDefinition struct {
	name          string      `metric_name:"memoryPool" source_type:"Attribute"`
	live          interface{} `metric_name:"mempool.live" source_type:"Gauge"`
	pool          interface{} `metric_name:"mempool.pool" source_type:"Gauge"`
	reqSize       interface{} `metric_name:"mempool.reqSizeInBytes" source_type:"Gauge"`
	allocatedSize interface{} `metric_name:"mempool.allocatedSizeInBytes" source_type:"Gauge"`
	allocs        interface{} `metric_name:"mempool.allocs" source_type:"Rate"`
	frees         interface{} `metric_name:"mempool.frees" source_type:"Rate"`
	recycles      interface{} `metric_name:"mempool.recycles" source_type:"Rate"`
	timeouts      interface{} `metric_name:"mempool.timeouts" source_type:"Rate"`
	tooSmall      interface{} `metric_name:"mempool.tooSmall" source_type:"Rate"`
	surplus       interface{} `metric_name:"mempool.surplus" source_type:"Rate"`
	ranDry        interface{} `metric_name:"mempool.ranDry" source_type:"Rate"`
}

// storageDefinition represents the data for a VarnishStorageSample event
type storageDefinition struct {
	name            string      `metric_name:"storage" source_type:"Attribute"`
	allocReqs       interface{} `metric_name:"storage.allocReqs" source_type:"Rate"`
	allocFails      interface{} `metric_name:"storage.allocFails" source_type:"Rate"`
	alloc           interface{} `metric_name:"storage.allocInBytes" source_type:"Rate"`
	free            interface{} `metric_name:"storage.freeInBytes" source_type:"Rate"`
	allocOustanding interface{} `metric_name:"storage.allocOustanding" source_type:"Gauge"`
	outstanding     interface{} `metric_name:"storage.outstandingInBytes" source_type:"Gauge"`
	available       interface{} `metric_name:"storage.availableInBytes" source_type:"Gauge"`
}

// backendDefinition represents the data to be set for a Backend Entity and VarnishBackendSample event
type backendDefinition struct {
	name                   string
	backendHappy           interface{} `metric_name:"backend.happy" source_type:"Gauge"`
	netBackendReqHeader    interface{} `metric_name:"net.backend.reqHeaderInBytes" source_type:"Rate"`
	netBackendReqBody      interface{} `metric_name:"net.backend.reqBodyInBytes" source_type:"Rate"`
	netBackendRespHeader   interface{} `metric_name:"net.backend.respHeaderInBytes" source_type:"Rate"`
	netBackendRespBody     interface{} `metric_name:"net.backend.respBodyInBytes" source_type:"Rate"`
	netBackendPipeHeader   interface{} `metric_name:"net.backend.pipeHeaderInBytes" source_type:"Rate"`
	netBackendPipeOut      interface{} `metric_name:"net.backend.pipeOutInBytes" source_type:"Rate"`
	netBackendPipeIn       interface{} `metric_name:"net.backend.pipeInInBytes" source_type:"Rate"`
	backendConnections     interface{} `metric_name:"backend.connections" source_type:"Gauge"`
	netBackendReq          interface{} `metric_name:"net.backend.req" source_type:"Rate"`
	backendUnhealtyFetches interface{} `metric_name:"backend.unhealtyFetches" source_type:"Rate"`
	backendBusyFetches     interface{} `metric_name:"backend.busyFetches" source_type:"Rate"`
	backendConFailed       interface{} `metric_name:"backend.conFailed" source_type:"Rate"`
	backendConNotAttempted interface{} `metric_name:"backend.conNotAttempted" source_type:"Rate"`
}
