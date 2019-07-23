package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flywithbug/gin_comm/http/handler_comm"
	"strings"
)

func RegisterRouters(r *gin.Engine,routers []handler_comm.GinHandleFunc)  {
	for _,v := range routers {
		if v.RouterGroup != nil{
			doRouteGroupRegister(v.Method,strings.ToLower(v.Route),v.Handler,v.RouterGroup)
		}else {
			doRouteRegister(v.Method,strings.ToLower(v.Route),v.Handler,r)
		}
	}
}



func doRouteGroupRegister(method, route string, handler gin.HandlerFunc, rGroup *gin.RouterGroup) {
	switch strings.ToUpper(method) {
	case "POST":
		rGroup.POST(route, handler)
	case "PUT":
		rGroup.PUT(route, handler)
	case "HEAD":
		rGroup.HEAD(route, handler)
	case "DELETE":
		rGroup.DELETE(route, handler)
	case "OPTIONS":
		rGroup.OPTIONS(route, handler)
	case "ANY":
		rGroup.Any(route, handler)
	default:
		rGroup.GET(route, handler)
	}
}

//普通routerType==routerTypeNormal
func doRouteRegister(method, route string, handler gin.HandlerFunc, r *gin.Engine) {
	switch strings.ToUpper(method) {
	case "POST":
		r.POST(route, handler)
	case "PUT":
		r.PUT(route, handler)
	case "HEAD":
		r.HEAD(route, handler)
	case "DELETE":
		r.DELETE(route, handler)
	case "OPTIONS":
		r.OPTIONS(route, handler)
	case "ANY":
		r.Any(route, handler)
	default:
		r.GET(route, handler)
	}
}
