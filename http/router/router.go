package gin_comm

import (
	"strings"

	"github.com/gin-gonic/gin"
)

//RegisterRouters 注册到路由
func RegisterRouters(r *gin.Engine, routers []GinHandleFunc) {
	for _, v := range routers {
		if v.RouterGroup != nil {
			doRouteGroupRegister(string(v.Method), strings.ToLower(v.Route), v.Handler, v.RouterGroup)
		} else {
			doRouteRegister(string(v.Method), strings.ToLower(v.Route), v.Handler, r)
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
