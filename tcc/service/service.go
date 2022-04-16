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

package service

import (
	"context"
	"fmt"
)

import (
	context2 "github.com/transaction-mesh/starfish/pkg/client/context"
	"github.com/transaction-mesh/starfish/pkg/client/tm"
)

type Service struct {
}

func (svc *Service) TCCCommitted(ctx context.Context) error {
	rootContext := ctx.(*context2.RootContext)
	businessActionContextA := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	// 业务参数全部放到 ActionContext 里
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	businessActionContextB := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextB.ActionContext["hello"] = "hello world,this is from BusinessActionContext B"

	resultA, err := TccProxyServiceA.Try(businessActionContextA)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultB, err := TccProxyServiceB.Try(businessActionContextB)
	fmt.Printf("result B is :%v", resultB)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) TCCCanceled(ctx context.Context) error {
	rootContext := ctx.(*context2.RootContext)
	businessActionContextA := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextA.ActionContext["hello"] = "hello world,this is from BusinessActionContext A"

	businessActionContextC := &context2.BusinessActionContext{
		RootContext:   rootContext,
		ActionContext: make(map[string]interface{}),
	}
	businessActionContextC.ActionContext["hello"] = "hello world,this is from BusinessActionContext C"

	resultA, err := TccProxyServiceA.Try(businessActionContextA)
	fmt.Printf("result A is :%v", resultA)
	if err != nil {
		return err
	}

	resultC, err := TccProxyServiceC.Try(businessActionContextC)
	fmt.Printf("result C is :%v", resultC)
	if err != nil {
		return err
	}

	return nil
}

var service = &Service{}

type ProxyService struct {
	*Service
	TCCCommitted func(ctx context.Context) error
	TCCCanceled  func(ctx context.Context) error
}

func (svc *ProxyService) GetServiceProxy() interface{} {
	return svc.Service
}

func (svc *ProxyService) GetMethodTransactionInfo(methodName string) *tm.TransactionInfo {
	return methodTransactionInfo[methodName]
}

var methodTransactionInfo = make(map[string]*tm.TransactionInfo)

func init() {
	methodTransactionInfo["TCCCommitted"] = &tm.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_COMMITTED",
		Propagation: tm.REQUIRED,
	}
	methodTransactionInfo["TCCCanceled"] = &tm.TransactionInfo{
		TimeOut:     60000000,
		Name:        "TCC_TEST_CANCELED",
		Propagation: tm.REQUIRED,
	}
}

var ProxySvc = &ProxyService{
	Service: service,
}
