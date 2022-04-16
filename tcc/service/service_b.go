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
	"fmt"
)

import (
	"github.com/transaction-mesh/starfish/pkg/client/context"
	"github.com/transaction-mesh/starfish/pkg/client/tcc"
)

type ServiceB struct {
}

func (svc *ServiceB) Try(ctx *context.BusinessActionContext) (bool, error) {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B Tried!")
	return true, nil
}

func (svc *ServiceB) Confirm(ctx *context.BusinessActionContext) bool {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B confirmed!")
	return true
}

func (svc *ServiceB) Cancel(ctx *context.BusinessActionContext) bool {
	word := ctx.ActionContext["hello"]
	fmt.Println(word)
	fmt.Println("Service B canceled!")
	return true
}

var serviceB = &ServiceB{}

type TCCProxyServiceB struct {
	*ServiceB

	Try func(ctx *context.BusinessActionContext) (bool, error) `TCCActionName:"ServiceB"`
}

func (svc *TCCProxyServiceB) GetTCCService() tcc.TCCService {
	return svc.ServiceB
}

var TccProxyServiceB = &TCCProxyServiceB{
	ServiceB: serviceB,
}
