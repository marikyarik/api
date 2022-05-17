package controllers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ApiController struct {

}

func NewApiController() *ApiController {
	return &ApiController{}
}

func (a *ApiController) GetHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	headers := make(map[string]string)
	ctx.Request.Header.VisitAll(func (key, value []byte) {
		headers[string(key)] = string(value)
	})

	reqHeadersBytes, err := json.Marshal(headers);
	
	if err != nil {
		ctx.Error("Could not Marshal Req Headers", fasthttp.StatusInternalServerError)
		return
	} 

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(reqHeadersBytes)
}