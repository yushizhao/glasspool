package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/node/btc"
	"github.com/yushizhao/glasspool/node/eos"
	"github.com/yushizhao/glasspool/node/eth"
	"github.com/yushizhao/glasspool/node/gnt"
	"github.com/yushizhao/glasspool/node/ltc"
	"github.com/yushizhao/glasspool/node/usdc"
	"github.com/yushizhao/glasspool/node/usdt"
	"github.com/yushizhao/glasspool/node/vet"
	"github.com/yushizhao/glasspool/node/xrp"
	"github.com/yushizhao/glasspool/node/zrx"
)

func main() {
	var bind = flag.String("bind", ":8080", "ip:port")
	var key = flag.String("key", "", "glasspool private key")
	flag.Parse()

	common.KEY = *key

	common.InitDB()
	eth.Init()
	vet.Init()
	usdc.Init()
	gnt.Init()
	zrx.Init()
	btc.Init()
	ltc.Init()
	usdt.Init()
	eos.Init()
	xrp.Init()

	go orderStateUpdate()

	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.POST("/api/v1/addresses/new").To(addressesNewHandler))
	ws.Route(ws.POST("/api/v1/addresses/{addr}/verify").To(addressesVerifyHandler))
	ws.Route(ws.POST("/api/v1/transactions").To(transactionsHandler))
	ws.Route(ws.GET("/api/v1/transactions/{orderId}").To(orderIDHandler))
	ws.Route(ws.POST("/api/v1/config/callback").To(configHandler))

	ws.Route(ws.GET("/deposit/{type}/{addr}/{value}").To(getDepositHandler))
	ws.Route(ws.POST("/deposit").To(depositHandler))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(*bind, nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
