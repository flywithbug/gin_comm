package gin_comm

import (
	"encoding/json"
	"net/http"
)

// Response 返回数据
type Response struct {
	Data map[string]interface{} `json:"data"`
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
}

// ToJSON json
func (s Response) ToJSON() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// NewResponse json
func NewResponse() *Response {
	res := new(Response)
	res.Code = http.StatusOK
	res.Data = make(map[string]interface{})
	res.Msg = ""
	return res
}

//ErrorResponse a
func ErrorResponse(errno int, msg string) *Response {
	r := NewResponse()
	r.Code = errno
	r.SetErrorInfo(errno, msg)
	return r
}

//SetErrorInfo a
func (s *Response) SetErrorInfo(errno int, msg string) {
	s.Code = errno
	s.Msg = msg
	//s.Data["msg"]=msg
}

//SetSuccessInfo   a
func (s *Response) SetSuccessInfo(code int, msg string) {
	s.Code = code
	s.Msg = msg
	//s.Data["msg"]=msg
}

//SetSuccess   a
func (s *Response) SetSuccess() {
	s.Code = http.StatusOK
	s.Msg = "success"
}

//SetResponseDataInfo   a
func (s *Response) SetResponseDataInfo(key string, value string) {
	s.Data[key] = value
}

//AddResponseInfo   a
func (s *Response) AddResponseInfo(key string, val interface{}) {
	s.Data[key] = val
}
