module recommendationS

go 1.18

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/neo4j/neo4j-go-driver/v4 v4.4.3
	go.mongodb.org/mongo-driver v1.9.1
)

require google.golang.org/protobuf v1.28.0 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2
)
