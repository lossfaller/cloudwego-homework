package caller

import (
	"github.com/cloudwego/kitex/client"
	bgeneric "github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lossfaller/hertz_demo/constants"
	"github.com/lossfaller/hertz_demo/kitex_gen/student_demo/studentservice"
)

var Rpccli studentservice.Client
var GeneralCli bgeneric.Client

func ClientInit() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	Rpccli, err = studentservice.NewClient("student_demo", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

var Provider *generic.ThriftContentProvider

// 热更新泛化
func GeneralCliInit() {

	var err error
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	Provider, err = generic.NewThriftContentProviderWithDynamicGo(constants.OldIdlContent, map[string]string{})
	if err != nil {
		panic(err)
	}
	// 构造http类型的泛化调用
	g, err := generic.HTTPThriftGeneric(Provider)
	if err != nil {
		panic(err)
	}
	GeneralCli, err = bgeneric.NewClient("student_demo", g, client.WithResolver(r))
	if err != nil {
		panic(err)
	}

}

////普通泛化
//func GeneralCliInit() {
//	var err error
//	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
//	if err != nil {
//		panic(err)
//	}
//	p, err := generic.NewThriftFileProvider("/Users/bytedance/go/src/github.com/lossfaller/hertz_demo/idl/student_demo.thrift")
//	if err != nil {
//		panic(err)
//	}
//	// 构造http类型的泛化调用
//	g, err := generic.HTTPThriftGeneric(p)
//	if err != nil {
//		panic(err)
//	}
//	GeneralCli, err = bgeneric.NewClient("student_demo", g, client.WithResolver(r))
//	if err != nil {
//		panic(err)
//	}
//
//}
