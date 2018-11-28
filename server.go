package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/node/btc"
	"github.com/yushizhao/glasspool/node/eth"
)

func main() {
	common.InitDB()
	eth.Init()
	btc.Init()

	go orderStateUpdate()

	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.POST("/api/v1/addresses/new").To(addressesNewHandler))
	ws.Route(ws.POST("/api/v1/addresses/{addr}/verify").To(addressesVerifyHandler))
	ws.Route(ws.POST("/api/v1/transactions").To(transactionsHandler))
	ws.Route(ws.GET("/api/v1/transactions/{orderId}").To(orderIDHandler))

	ws.Route(ws.GET("/deposit/{type}/{addr}/{value}").To(getDepositHandler))
	ws.Route(ws.POST("/deposit").To(depositHandler))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
