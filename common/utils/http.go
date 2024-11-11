package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result ...
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResult(code int, msg string, data any) *Result {
	return &Result{Code: code, Msg: msg, Data: data}
}

func (r *Result) SetCode(code int) {
	r.Code = code
}

func (r *Result) SetMsg(msg string) {
	r.Msg = msg
}

func (r *Result) SetData(data any) {
	r.Data = data
}

// Success ...
func (r *Result) Success(data any) {
	r.SetCode(http.StatusOK)
	r.SetMsg("success")
	r.SetData(data)
}

// Error ...
func (r *Result) Error(code int, msg string) {
	r.SetCode(code)
	r.SetMsg(msg)
}

// BadRequest ...
func (r *Result) BadRequest(msg string) {
	r.SetCode(http.StatusBadGateway)
	r.SetMsg(msg)
}

// InternalServerError ...
func (r *Result) InternalServerError(msg string) {
	r.SetCode(http.StatusInternalServerError)
	r.SetMsg(msg)
}

func (r *Result) HTML(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "error.html", gin.H{
		"title":   "404 - ohUrlShortener",
		"code":    http.StatusNotFound,
		"message": "不支持的打开方式",
		"label":   "Status Not Found",
	})
}
