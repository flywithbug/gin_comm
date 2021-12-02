package router

import "github.com/gin-gonic/gin"

//StateType a
type StateType int

const (
	//RouterTypeNormal a
	RouterTypeNormal StateType = iota
	//RouterTypeNeedAuth a
	RouterTypeNeedAuth
)

//MethodType a
type MethodType string

const (
	POST    MethodType = "POST"
	GET     MethodType = "GET"
	PUT     MethodType = "PUT"
	DELETE  MethodType = "DELETE"
	HEADER  MethodType = "HEADER"
	ANY     MethodType = "ANY"
	OPTIONS MethodType = "OPTIONS"
)

type GinHandleFunc struct {
	Handler     gin.HandlerFunc
	Method      MethodType
	Route       string
	RouterGroup *gin.RouterGroup
	RouterType  StateType
}
