module github.com/gotrx/starfish-samples

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gotrx/starfish v1.0.0-rc2
	github.com/gotrx/starfish-samples/tcc v0.0.0-00010101000000-000000000000
)

replace github.com/gotrx/starfish => ../../starfish

replace github.com/gotrx/starfish-samples/tcc => ../../starfish-samples/tcc
