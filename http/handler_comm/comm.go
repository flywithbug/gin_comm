package handler_comm

import "github.com/gin-gonic/gin"

type StateType int

const (
	RouterTypeNormal StateType = iota
	RouterTypeNeedAuth
)

type MethodType string

const (
	MethodPOST   MethodType = "POST"
	MethodGET    MethodType = "GET"
	MethodPUT    MethodType = "PUT"
	MethodDELETE MethodType = "DELETE"
)

type GinHandleFunc struct {
	Handler     gin.HandlerFunc
	Method      MethodType
	Route       string
	RouterGroup *gin.RouterGroup
	RouterType  StateType
}
