package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	student_demo "github.com/lossfaller/kitex_demo/kitex_gen/student_demo/studentservice"
	"log"
	"net"
)

var (
	severList = []string{":9998"}
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range severList {
		go func(s string) {
			addr, _ := net.ResolveTCPAddr("tcp", s)
			svr := student_demo.NewServer(new(StudentServiceImpl), server.WithRegistry(r),
				server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student_demo"}),
				server.WithServiceAddr(addr))

			err = svr.Run()
			if err != nil {
				log.Println(err.Error())
			}
		}(s)
	}
	var c chan struct{} //nil channel
	<-c
}
