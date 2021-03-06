package main

import (
	"api/controllers"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

func main() {
	log.Println("Init controllers")
	a := controllers.NewApiController()

	log.Println("Init router")
	r := router.New()

	r.GET("/api/headers", a.GetHeaders)

	r.PanicHandler  = func (ctx *fasthttp.RequestCtx, data interface{})  {
		log.Println(data)
		ctx.SetStatusCode(500)
		ctx.SetBodyString("Internal server error")
	}

	log.Println("Starting webserver")

	ln, err := reuseport.Listen("tcp4", ":80")
	if err != nil {
		log.Fatalf("error in reuseport listener: %s", err.Error())
	}

	if err = fasthttp.Serve(ln, r.Handler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err.Error())
	}
}
