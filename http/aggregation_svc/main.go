package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gotrx/starfish/pkg/client"
	"github.com/gotrx/starfish/pkg/client/config"
	"github.com/gotrx/starfish/pkg/client/tm"

	"github.com/gotrx/starfish-samples/aggregation_svc/svc"
)

var configPath = "/Users/scottlewis/dksl/temp/starfish-samples/http/aggregation_svc/conf/client.yml"

func main() {
	r := gin.Default()
	config.InitConf(configPath)
	client.NewRpcClient()
	tm.Implement(svc.ProxySvc)

	r.GET("/createSoCommit", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, false)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.GET("/createSoRollback", func(c *gin.Context) {

		svc.ProxySvc.CreateSo(c, true)

		c.JSON(200, gin.H{
			"success": true,
			"message": "success",
		})
	})

	r.Run(":8003")
}
