module github.com/transaction-mesh/starfish-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gogf/gf v1.15.6
	github.com/google/uuid v1.2.0
	github.com/transaction-mesh/mysql v1.0.0-rc4
	github.com/transaction-mesh/starfish v1.0.2-rc1
	gorm.io/gorm v1.21.8

)

replace github.com/transaction-mesh/mysql => ../../mysql

replace github.com/transaction-mesh/starfish => ../../starfish
