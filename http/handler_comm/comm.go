package handler_comm


import "github.com/gin-gonic/gin"

type StateType int

const (
	RouterTypeNormal StateType = iota
	RouterTypeNeedAuth
)



type GinHandleFunc struct {
	Handler    gin.HandlerFunc
	Method     string
	Route      string
	RouterGroup      *gin.RouterGroup
}
