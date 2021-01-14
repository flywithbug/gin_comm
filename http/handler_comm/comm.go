package handler_comm

import "github.com/gin-gonic/gin"

type StateType int

const (
	RouterTypeNormal StateType = iota
	RouterTypeNeedAuth
)

type MethodType string

const (
	POST   MethodType = "POST"
	GET    MethodType = "GET"
	PUT    MethodType = "PUT"
	DELETE MethodType = "DELETE"
)

type GinHandleFunc struct {
	Handler     gin.HandlerFunc
	Method      MethodType
	Route       string
	RouterGroup *gin.RouterGroup
	RouterType  StateType
}
