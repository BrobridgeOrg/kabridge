module github.com/BrobridgeOrg/kabridge

go 1.21.4

require github.com/segmentio/kafka-go v0.4.47

require (
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.19 // indirect
)

//replace github.com/segmentio/kafka-go => ../../opensource/kafka-go
