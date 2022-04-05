package main

import (
	"context"
	"net/http"
	"time"
)

import (
	"github.com/gin-gonic/gin"

	"github.com/transaction-mesh/mysql"

	"github.com/transaction-mesh/starfish/pkg/client"
	"github.com/transaction-mesh/starfish/pkg/client/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

import (
	dialector "github.com/transaction-mesh/starfish-samples/dialector/mysql"
	"github.com/transaction-mesh/starfish-samples/product_svc/dao"
)

const configPath = "conf/client.yml"

func main() {
	r := gin.Default()
	config.InitConf(configPath)
	client.NewRpcClient()
	mysql.InitDataResourceManager()
	mysql.RegisterResource(config.GetATConfig().DSN)

	db, err := gorm.Open(
		dialector.Open(config.GetATConfig().DSN),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			}})
	if err != nil {
		panic(err)
	}
	DB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(4 * time.Hour)

	d := &dao.Dao{DB: db}

	r.POST("/allocateInventory", func(c *gin.Context) {
		type req struct {
			Req []*dao.AllocateInventoryReq
		}
		var q req
		if err := c.ShouldBindJSON(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := d.AllocateInventory(
			context.WithValue(
				context.Background(),
				mysql.XID,
				c.Request.Header.Get("XID")),
			q.Req)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "fail",
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"message": "success",
			})
		}
	})

	r.Run(":8001")
}
