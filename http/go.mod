module github.com/transaction-mesh/starfish-samples

go 1.15

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/mock v1.5.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/nacos-group/nacos-sdk-go v1.0.9 // indirect
	github.com/prometheus/common v0.28.0 // indirect
	github.com/transaction-mesh/mysql v1.0.1
	github.com/transaction-mesh/starfish v1.0.2-rc1
	github.com/ugorji/go v1.2.6 // indirect
	go.etcd.io/etcd/api/v3 v3.5.1 // indirect
	go.etcd.io/etcd/client/v2 v2.305.0 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/net v0.0.0-20211105192438-b53810dc28af // indirect
	google.golang.org/genproto v0.0.0-20211104193956-4c6863e31247 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/transaction-mesh/mysql => ../../mysql
	github.com/transaction-mesh/starfish => ../../starfish
)
