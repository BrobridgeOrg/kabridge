package main

import (
	addOffsetsToTxnAPI "github.com/segmentio/kafka-go/protocol/addoffsetstotxn"
	addPartitionsToTxnAPI "github.com/segmentio/kafka-go/protocol/addpartitionstotxn"
	alterConfigsAPI "github.com/segmentio/kafka-go/protocol/alterconfigs"
	alterPartitionReassignmentsAPI "github.com/segmentio/kafka-go/protocol/alterpartitionreassignments"

	//	alterReplicaLogDirsAPI "github.com/segmentio/kafka-go/protocol/alterreplicalogdirs"
	apiVersionsAPI "github.com/segmentio/kafka-go/protocol/apiversions"
	//	controlledShutdownAPI "github.com/segmentio/kafka-go/protocol/controlledshutdown"
	createAclsAPI "github.com/segmentio/kafka-go/protocol/createacls"

	//	createDelegationTokenAPI "github.com/segmentio/kafka-go/protocol/createdelegationtoken"
	createPartitionsAPI "github.com/segmentio/kafka-go/protocol/createpartitions"
	createTopicsAPI "github.com/segmentio/kafka-go/protocol/createtopics"
	deleteAclsAPI "github.com/segmentio/kafka-go/protocol/deleteacls"
	deleteGroupsAPI "github.com/segmentio/kafka-go/protocol/deletegroups"

	//	deleteRecordsAPI "github.com/segmentio/kafka-go/protocol/deleterecords"
	deleteTopicsAPI "github.com/segmentio/kafka-go/protocol/deletetopics"
	describeAclsAPI "github.com/segmentio/kafka-go/protocol/describeacls"
	describeConfigsAPI "github.com/segmentio/kafka-go/protocol/describeconfigs"

	//	describeDelegationTokenAPI "github.com/segmentio/kafka-go/protocol/describedelegationtoken"
	describeGroupsAPI "github.com/segmentio/kafka-go/protocol/describegroups"
	//	describeLogDirsAPI "github.com/segmentio/kafka-go/protocol/describelogdirs"
	electLeadersAPI "github.com/segmentio/kafka-go/protocol/electleaders"
	endTxnAPI "github.com/segmentio/kafka-go/protocol/endtxn"

	//	expireDelegationTokenAPI "github.com/segmentio/kafka-go/protocol/expiredelegationtoken"
	fetchAPI "github.com/segmentio/kafka-go/protocol/fetch"
	findCoordinatorAPI "github.com/segmentio/kafka-go/protocol/findcoordinator"
	heartbeatAPI "github.com/segmentio/kafka-go/protocol/heartbeat"
	incrementalAlterConfigsAPI "github.com/segmentio/kafka-go/protocol/incrementalalterconfigs"
	initProducerIdAPI "github.com/segmentio/kafka-go/protocol/initproducerid"
	joinGroupAPI "github.com/segmentio/kafka-go/protocol/joingroup"

	//	leaderAndIsrAPI "github.com/segmentio/kafka-go/protocol/leaderandisr"
	leaveGroupAPI "github.com/segmentio/kafka-go/protocol/leavegroup"
	listGroupsAPI "github.com/segmentio/kafka-go/protocol/listgroups"
	listOffsetsAPI "github.com/segmentio/kafka-go/protocol/listoffsets"
	listPartitionReassignmentsAPI "github.com/segmentio/kafka-go/protocol/listpartitionreassignments"
	metadataAPI "github.com/segmentio/kafka-go/protocol/metadata"
	offsetCommitAPI "github.com/segmentio/kafka-go/protocol/offsetcommit"
	offsetDeleteAPI "github.com/segmentio/kafka-go/protocol/offsetdelete"
	offsetFetchAPI "github.com/segmentio/kafka-go/protocol/offsetfetch"

	//	offsetForLeaderEpochAPI "github.com/segmentio/kafka-go/protocol/offsetforleaderepoch"
	produceAPI "github.com/segmentio/kafka-go/protocol/produce"
	//	renewDelegationTokenAPI "github.com/segmentio/kafka-go/protocol/renewdelegationtoken"
	saslAuthenticateAPI "github.com/segmentio/kafka-go/protocol/saslauthenticate"
	saslHandshakeAPI "github.com/segmentio/kafka-go/protocol/saslhandshake"

	//	stopReplicaAPI "github.com/segmentio/kafka-go/protocol/stopreplica"
	syncGroupAPI "github.com/segmentio/kafka-go/protocol/syncgroup"
	txnOffsetCommitAPI "github.com/segmentio/kafka-go/protocol/txnoffsetcommit"
	//	updateMetadataAPI "github.com/segmentio/kafka-go/protocol/updatemetadata"
	//	writeTxnMarkersAPI "github.com/segmentio/kafka-go/protocol/writetxnmarkers"
)

type apiKey int16

const (
	produce                     apiKey = 0
	fetch                       apiKey = 1
	listOffsets                 apiKey = 2
	metadata                    apiKey = 3
	leaderAndIsr                apiKey = 4
	stopReplica                 apiKey = 5
	updateMetadata              apiKey = 6
	controlledShutdown          apiKey = 7
	offsetCommit                apiKey = 8
	offsetFetch                 apiKey = 9
	findCoordinator             apiKey = 10
	joinGroup                   apiKey = 11
	heartbeat                   apiKey = 12
	leaveGroup                  apiKey = 13
	syncGroup                   apiKey = 14
	describeGroups              apiKey = 15
	listGroups                  apiKey = 16
	saslHandshake               apiKey = 17
	apiVersions                 apiKey = 18
	createTopics                apiKey = 19
	deleteTopics                apiKey = 20
	deleteRecords               apiKey = 21
	initProducerId              apiKey = 22
	offsetForLeaderEpoch        apiKey = 23
	addPartitionsToTxn          apiKey = 24
	addOffsetsToTxn             apiKey = 25
	endTxn                      apiKey = 26
	writeTxnMarkers             apiKey = 27
	txnOffsetCommit             apiKey = 28
	describeAcls                apiKey = 29
	createAcls                  apiKey = 30
	deleteAcls                  apiKey = 31
	describeConfigs             apiKey = 32
	alterConfigs                apiKey = 33
	alterReplicaLogDirs         apiKey = 34
	describeLogDirs             apiKey = 35
	saslAuthenticate            apiKey = 36
	createPartitions            apiKey = 37
	createDelegationToken       apiKey = 38
	renewDelegationToken        apiKey = 39
	expireDelegationToken       apiKey = 40
	describeDelegationToken     apiKey = 41
	deleteGroups                apiKey = 42
	electLeaders                apiKey = 43
	incrementalAlterConfigs     apiKey = 44
	alterPartitionReassignments apiKey = 45
	listPartitionReassignments  apiKey = 46
	offsetDelete                apiKey = 47
)

var apiKeys = map[apiKey]string{
	produce:                     "Produce",
	fetch:                       "Fetch",
	listOffsets:                 "ListOffsets",
	metadata:                    "Metadata",
	leaderAndIsr:                "LeaderAndIsr",
	stopReplica:                 "StopReplica",
	updateMetadata:              "UpdateMetadata",
	controlledShutdown:          "ControlledShutdown",
	offsetCommit:                "OffsetCommit",
	offsetFetch:                 "OffsetFetch",
	findCoordinator:             "FindCoordinator",
	joinGroup:                   "JoinGroup",
	heartbeat:                   "Heartbeat",
	leaveGroup:                  "LeaveGroup",
	syncGroup:                   "SyncGroup",
	describeGroups:              "DescribeGroups",
	listGroups:                  "ListGroups",
	saslHandshake:               "SaslHandshake",
	apiVersions:                 "ApiVersions",
	createTopics:                "CreateTopics",
	deleteTopics:                "DeleteTopics",
	deleteRecords:               "DeleteRecords",
	initProducerId:              "InitProducerId",
	offsetForLeaderEpoch:        "OffsetForLeaderEpoch",
	addPartitionsToTxn:          "AddPartitionsToTxn",
	addOffsetsToTxn:             "AddOffsetsToTxn",
	endTxn:                      "EndTxn",
	writeTxnMarkers:             "WriteTxnMarkers",
	txnOffsetCommit:             "TxnOffsetCommit",
	describeAcls:                "DescribeAcls",
	createAcls:                  "CreateAcls",
	deleteAcls:                  "DeleteAcls",
	describeConfigs:             "DescribeConfigs",
	alterConfigs:                "AlterConfigs",
	alterReplicaLogDirs:         "AlterReplicaLogDirs",
	describeLogDirs:             "DescribeLogDirs",
	saslAuthenticate:            "SaslAuthenticate",
	createPartitions:            "CreatePartitions",
	createDelegationToken:       "CreateDelegationToken",
	renewDelegationToken:        "RenewDelegationToken",
	expireDelegationToken:       "ExpireDelegationToken",
	describeDelegationToken:     "DescribeDelegationToken",
	deleteGroups:                "DeleteGroups",
	electLeaders:                "ElectLeaders",
	incrementalAlterConfigs:     "IncrementalAlterConfigs",
	alterPartitionReassignments: "AlterPartitionReassignments",
	listPartitionReassignments:  "ListPartitionReassignments",
	offsetDelete:                "OffsetDelete",
}

type APIVersion struct {
	min int16
	max int16
}

var APIVersions = map[apiKey]APIVersion{
	produce:                     {0, 8},
	fetch:                       {0, 11},
	listOffsets:                 {0, 5},
	metadata:                    {0, 8},
	leaderAndIsr:                {0, 5},
	stopReplica:                 {0, 3},
	updateMetadata:              {0, 7},
	controlledShutdown:          {0, 2},
	offsetCommit:                {0, 8},
	offsetFetch:                 {0, 7},
	findCoordinator:             {0, 3},
	joinGroup:                   {0, 8},
	heartbeat:                   {0, 4},
	leaveGroup:                  {0, 4},
	syncGroup:                   {0, 5},
	describeGroups:              {0, 5},
	listGroups:                  {0, 2},
	saslHandshake:               {0, 1},
	apiVersions:                 {0, 3},
	createTopics:                {0, 5},
	deleteTopics:                {0, 4},
	deleteRecords:               {0, 1},
	initProducerId:              {0, 3},
	offsetForLeaderEpoch:        {0, 4},
	addPartitionsToTxn:          {0, 3},
	addOffsetsToTxn:             {0, 3},
	endTxn:                      {0, 3},
	writeTxnMarkers:             {0, 1},
	txnOffsetCommit:             {0, 3},
	describeAcls:                {0, 2},
	createAcls:                  {0, 2},
	deleteAcls:                  {0, 2},
	describeConfigs:             {0, 3},
	alterConfigs:                {0, 1},
	alterReplicaLogDirs:         {0, 2},
	describeLogDirs:             {0, 2},
	saslAuthenticate:            {0, 2},
	createPartitions:            {0, 2},
	createDelegationToken:       {0, 3},
	renewDelegationToken:        {0, 2},
	expireDelegationToken:       {0, 2},
	describeDelegationToken:     {0, 2},
	deleteGroups:                {0, 2},
	electLeaders:                {0, 3},
	incrementalAlterConfigs:     {0, 1},
	alterPartitionReassignments: {0, 1},
	listPartitionReassignments:  {0, 1},
	offsetDelete:                {0, 1},
}

var apiRequests = map[apiKey]interface{}{
	produce:     produceAPI.Request{},
	fetch:       fetchAPI.Request{},
	listOffsets: listOffsetsAPI.Request{},
	metadata:    metadataAPI.Request{},
	//	leaderAndIsr:                leaderAndIsrAPI.Request{},
	//	stopReplica:                 stopReplicaAPI.Request{},
	//	updateMetadata:              updateMetadataAPI.Request{},
	//	controlledShutdown:          controlledShutdownAPI.Request{},
	offsetCommit:    offsetCommitAPI.Request{},
	offsetFetch:     offsetFetchAPI.Request{},
	findCoordinator: findCoordinatorAPI.Request{},
	joinGroup:       joinGroupAPI.Request{},
	heartbeat:       heartbeatAPI.Request{},
	leaveGroup:      leaveGroupAPI.Request{},
	syncGroup:       syncGroupAPI.Request{},
	describeGroups:  describeGroupsAPI.Request{},
	listGroups:      listGroupsAPI.Request{},
	saslHandshake:   saslHandshakeAPI.Request{},
	apiVersions:     apiVersionsAPI.Request{},
	createTopics:    createTopicsAPI.Request{},
	deleteTopics:    deleteTopicsAPI.Request{},
	//	deleteRecords:               deleteRecordsAPI.Request{},
	initProducerId: initProducerIdAPI.Request{},
	//	offsetForLeaderEpoch:        offsetForLeaderEpochAPI.Request{},
	addPartitionsToTxn: addPartitionsToTxnAPI.Request{},
	addOffsetsToTxn:    addOffsetsToTxnAPI.Request{},
	endTxn:             endTxnAPI.Request{},
	//	writeTxnMarkers:             writeTxnMarkersAPI.Request{},
	txnOffsetCommit: txnOffsetCommitAPI.Request{},
	describeAcls:    describeAclsAPI.Request{},
	createAcls:      createAclsAPI.Request{},
	deleteAcls:      deleteAclsAPI.Request{},
	describeConfigs: describeConfigsAPI.Request{},
	alterConfigs:    alterConfigsAPI.Request{},
	//	alterReplicaLogDirs:         alterReplicaLogDirsAPI.Request{},
	//	describeLogDirs:             describeLogDirsAPI.Request{},
	saslAuthenticate: saslAuthenticateAPI.Request{},
	createPartitions: createPartitionsAPI.Request{},
	//	createDelegationToken:       createDelegationTokenAPI.Request{},
	//	renewDelegationToken:        renewDelegationTokenAPI.Request{},
	//	expireDelegationToken:       expireDelegationTokenAPI.Request{},
	//	describeDelegationToken:     describeDelegationTokenAPI.Request{},
	deleteGroups:                deleteGroupsAPI.Request{},
	electLeaders:                electLeadersAPI.Request{},
	incrementalAlterConfigs:     incrementalAlterConfigsAPI.Request{},
	alterPartitionReassignments: alterPartitionReassignmentsAPI.Request{},
	listPartitionReassignments:  listPartitionReassignmentsAPI.Request{},
	offsetDelete:                offsetDeleteAPI.Request{},
}

var apiResponses = map[apiKey]interface{}{
	produce:     produceAPI.Response{},
	fetch:       fetchAPI.Response{},
	listOffsets: listOffsetsAPI.Response{},
	metadata:    metadataAPI.Response{},
	//	leaderAndIsr:                leaderAndIsrAPI.Response{},
	//	stopReplica:                 stopReplicaAPI.Response{},
	//	updateMetadata:              updateMetadataAPI.Response{},
	//	controlledShutdown:          controlledShutdownAPI.Response{},
	offsetCommit:    offsetCommitAPI.Response{},
	offsetFetch:     offsetFetchAPI.Response{},
	findCoordinator: findCoordinatorAPI.Response{},
	joinGroup:       joinGroupAPI.Response{},
	heartbeat:       heartbeatAPI.Response{},
	leaveGroup:      leaveGroupAPI.Response{},
	syncGroup:       syncGroupAPI.Response{},
	describeGroups:  describeGroupsAPI.Response{},
	listGroups:      listGroupsAPI.Response{},
	saslHandshake:   saslHandshakeAPI.Response{},
	apiVersions:     apiVersionsAPI.Response{},
	createTopics:    createTopicsAPI.Response{},
	deleteTopics:    deleteTopicsAPI.Response{},
	//	deleteRecords:               deleteRecordsAPI.Response{},
	initProducerId: initProducerIdAPI.Response{},
	//	offsetForLeaderEpoch:        offsetForLeaderEpochAPI.Response{},
	addPartitionsToTxn: addPartitionsToTxnAPI.Response{},
	addOffsetsToTxn:    addOffsetsToTxnAPI.Response{},
	endTxn:             endTxnAPI.Response{},
	//	writeTxnMarkers:             writeTxnMarkersAPI.Response{},
	txnOffsetCommit: txnOffsetCommitAPI.Response{},
	describeAcls:    describeAclsAPI.Response{},
	createAcls:      createAclsAPI.Response{},
	deleteAcls:      deleteAclsAPI.Response{},
	describeConfigs: describeConfigsAPI.Response{},
	alterConfigs:    alterConfigsAPI.Response{},
	//	alterReplicaLogDirs:         alterReplicaLogDirsAPI.Response{},
	//	describeLogDirs:             describeLogDirsAPI.Response{},
	saslAuthenticate: saslAuthenticateAPI.Response{},
	createPartitions: createPartitionsAPI.Response{},
	//	createDelegationToken:       createDelegationTokenAPI.Response{},
	//	renewDelegationToken:        renewDelegationTokenAPI.Response{},
	//	expireDelegationToken:       expireDelegationTokenAPI.Response{},
	//	describeDelegationToken:     describeDelegationTokenAPI.Response{},
	deleteGroups:                deleteGroupsAPI.Response{},
	electLeaders:                electLeadersAPI.Response{},
	incrementalAlterConfigs:     incrementalAlterConfigsAPI.Response{},
	alterPartitionReassignments: alterPartitionReassignmentsAPI.Response{},
	listPartitionReassignments:  listPartitionReassignmentsAPI.Response{},
	offsetDelete:                offsetDeleteAPI.Response{},
}
