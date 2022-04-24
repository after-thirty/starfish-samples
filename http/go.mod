module github.com/transaction-mesh/starfish-samples

go 1.15

require (
	github.com/dubbogo/tools v1.0.9 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.3.0
	github.com/transaction-mesh/mysql v1.0.1
	github.com/transaction-mesh/starfish v1.0.2-rc1
)

replace (
	github.com/transaction-mesh/mysql => ../../mysql
	github.com/transaction-mesh/starfish => ../../starfish
)
