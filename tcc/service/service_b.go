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
