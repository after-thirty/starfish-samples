/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dao

import (
	"context"
	"database/sql"
)

import (
	"gorm.io/gorm"
)

type Inventory struct {
	Sysno           uint64
	ProductSysno    uint64
	AccountQty      int32
	AvailableQty    int32
	AllocatedQty    int32
	AdjustLockedQty int32
}

type Dao struct {
	DB *gorm.DB
}

type AllocateInventoryReq struct {
	ProductSysNo int64
	Qty          int32
}

func (dao *Dao) AllocateInventory(ctx context.Context, reqs []*AllocateInventoryReq) error {
	tx := dao.DB.WithContext(ctx).Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})

	for _, req := range reqs {
		if err := tx.Model(&Inventory{}).
			Where("product_sysno = ? and available_qty >= ?", req.ProductSysNo, req.Qty).
			UpdateColumns(map[string]interface{}{
				"available_qty": gorm.Expr("available_qty - ?", req.Qty),
				"allocated_qty": gorm.Expr("allocated_qty + ?", req.Qty),
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
