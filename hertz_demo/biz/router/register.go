// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	student_demo "github.com/lossfaller/hertz_demo/biz/router/student_demo"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	student_demo.Register(r)
}