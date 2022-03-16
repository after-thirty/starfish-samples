module github.com/gotrx/starfish-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/gotrx/mysql v1.0.1
	github.com/gotrx/starfish v1.0.2-rc1
)

replace (
	github.com/gotrx/starfish v1.0.2-rc1 => /Users/scottlewis/dksl/git/1/starfish
	github.com/gotrx/mysql v1.0.1 => /Users/scottlewis/dksl/current/mysql
)