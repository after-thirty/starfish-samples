package main

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-mesh/starfish-samples/tcc/service"
	"github.com/transaction-mesh/starfish/pkg/client"
	"github.com/transaction-mesh/starfish/pkg/client/config"
	"github.com/transaction-mesh/starfish/pkg/client/tcc"
	"github.com/transaction-mesh/starfish/pkg/client/tm"
)

func main() {
	r := gin.Default()

	config.InitConf("config/client.yml")
	client.NewRpcClient()
	tcc.InitTCCResourceManager()

	tm.Implement(service.ProxySvc)
	tcc.ImplementTCC(service.TccProxyServiceA)
	tcc.ImplementTCC(service.TccProxyServiceB)
	tcc.ImplementTCC(service.TccProxyServiceC)

	r.GET("/commit", func(c *gin.Context) {
		service.ProxySvc.TCCCommitted(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/rollback", func(c *gin.Context) {
		service.ProxySvc.TCCCanceled(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
