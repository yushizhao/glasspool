package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/go-resty/resty"
	"github.com/yushizhao/glasspool/addr"
	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/order"
)

func callback(url string, body common.ReturnBody) {

	returnBody := common.ReturnBodySig{
		Message:   body.Message,
		Result:    body.Result,
		Status:    body.Status,
		Crypto:    "ecc",
		Timestamp: common.Timestamp(),
		Sig:       common.SigJSON{},
		SendAgain: false,
	}

	if common.ECC {
		dataBytes, _ := json.Marshal(body.Result)

		var f interface{}
		json.Unmarshal(dataBytes, &f)
		data := f.(map[string]interface{})

		data["timestamp"] = returnBody.Timestamp

		msg := common.MapMessage(data)
		sig, err := common.ECCSignature([]byte(msg), common.KEY)
		returnBody.Sig = sig
		if err != nil {
			log.Printf("callback ECCSignature %v: %v", url, err)
			return
		}
	}

	log.Printf("callback " + url)
	log.Printf(common.JSONstring(returnBody))

	resp, err := resty.R().
		// SetHeader("Content-Type", common.MIME).
		SetBody(returnBody).
		Post(url)
	if err != nil {
		log.Printf("callback err %v: %v", url, err)
		return
	}

	log.Printf("URL: %v", url)
	// explore response object
	log.Printf("Error: %v", err)
	log.Printf("Response Status Code: %v", resp.StatusCode())
	log.Printf("Response Status: %v", resp.Status())
	log.Printf("Response Body: %v", resp) // or resp.String() or string(resp.Body())
}

func safeRead(req *restful.Request) (*common.POSTBody, error) {
	var f interface{}
	err := req.ReadEntity(&f)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(common.JSONstring(f))

	dataBytes, err := json.Marshal(f)
	postBody := new(common.POSTBody)
	json.Unmarshal(dataBytes, postBody)
	return postBody, err
}

func addressesNewHandler(req *restful.Request, resp *restful.Response) {
	returnBody := new(common.ReturnBody)

	postBody, err := safeRead(req)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	log.Println(common.JSONstring(postBody))

	postAddressesNew := new(addr.POSTAddressesNew)
	postBody.Extract(postAddressesNew)

	postAddressesNewResult, err := postAddressesNew.Process()
	if err != nil {
		// To Do: error code
		io.WriteString(resp, common.JSONstring(returnBody))
	} else {
		returnBody.Normalize(postAddressesNewResult)

		io.WriteString(resp, common.JSONstring(returnBody))
	}

	postBody.Path = "/api/v1/addresses/new"
	common.Gdb.POSTED([]byte(common.JSONstring(postBody)), []byte(common.JSONstring(returnBody)))
	k := postAddressesNewResult.Type + postAddressesNewResult.Address
	v := postAddressesNew.Callback
	common.Gdb.DEPOSIT([]byte(k), []byte(v))

	if common.AUTODEPOSIT {
		tx, err := postAddressesNewResult.AutoDeposit()
		if err != nil {
			log.Printf("AUTODEPOSIT: %v", err)
			return
		}
		// callback?
		o, err := tx.Submit("DEPOSIT")
		if err != nil {
			log.Printf("AUTODEPOSIT: %v", err)
			return
		}
		common.Gdb.ORDER([]byte(o.Index), []byte(common.JSONstring(o)))
		log.Println(common.JSONstring(o))
	}
}

func addressesVerifyHandler(req *restful.Request, resp *restful.Response) {
	a := req.PathParameter("addr")
	returnBody := new(common.ReturnBody)

	postBody, err := safeRead(req)

	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(common.JSONstring(postBody))

	postAddressesVerify := new(addr.POSTAddressesVerify)
	postBody.Extract(postAddressesVerify)

	postAddressesVerifyResult := postAddressesVerify.Process(a)
	returnBody.Normalize(postAddressesVerifyResult)
	io.WriteString(resp, common.JSONstring(returnBody))
}

// update callback?
func transactionsHandler(req *restful.Request, resp *restful.Response) {
	returnBody := new(common.ReturnBody)

	postBody, err := safeRead(req)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(common.JSONstring(postBody))

	postTransactions := new(order.POSTTransactions)
	postBody.Extract(postTransactions)

	postTransactionsResults, err := postTransactions.Process("WITHDRAW")
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	returnBody.Normalize(postTransactionsResults)
	io.WriteString(resp, common.JSONstring(returnBody))

	postBody.Path = "/api/v1/transactions"

	k := postTransactions.Type + postTransactions.To
	v := postTransactions.Callback
	common.Gdb.WITHDRAW([]byte(k), []byte(v))
	common.Gdb.POSTED([]byte(common.JSONstring(postBody)), []byte(common.JSONstring(returnBody)))
	common.Gdb.ORDER([]byte(postTransactionsResults.Index), []byte(common.JSONstring(postTransactionsResults)))
}

func orderIDHandler(req *restful.Request, resp *restful.Response) {
	orderID := req.PathParameter("orderId")
	orderBytes := common.Gdb.QueryOrder(orderID)
	if orderBytes == nil {
		io.WriteString(resp, "orderId: "+orderID+" not exist\n")
		return
	}

	order := new(order.POSTTransactionsResult)
	returnBody := new(common.ReturnBody)
	json.Unmarshal(orderBytes, &order)
	returnBody.Normalize(order)
	io.WriteString(resp, common.JSONstring(returnBody))
}

func configHandler(req *restful.Request, resp *restful.Response) {
	returnBody := new(common.ReturnBody)
	returnBody.Normalize("dummy")
	io.WriteString(resp, common.JSONstring(returnBody))
}

func depositHandler(req *restful.Request, resp *restful.Response) {
	txData := new(order.TxData)
	returnBody := new(common.ReturnBody)
	err := req.ReadEntity(txData)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(common.JSONstring(txData))

	// target first point
	key := txData.Type + txData.To[0].Address
	urlBytes := common.Gdb.Callback("DEPOSIT", key)
	if urlBytes == nil {
		io.WriteString(resp, "no callback found")
		return
	}
	url := string(urlBytes)

	orderData, err := txData.Submit("DEPOSIT")
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	common.Gdb.ORDER([]byte(orderData.Index), []byte(common.JSONstring(orderData)))
	io.WriteString(resp, common.JSONstring(orderData))

	returnBody.Normalize(orderData)

	callback(url, *returnBody)
}

func getDepositHandler(req *restful.Request, resp *restful.Response) {
	t := req.PathParameter("type")
	address := req.PathParameter("addr")
	value := req.PathParameter("value")

	urlBytes := common.Gdb.Callback("DEPOSIT", t+address)
	if urlBytes == nil {
		io.WriteString(resp, "no callback found")
		return
	}
	url := string(urlBytes)

	postTransactions := new(order.POSTTransactions)
	postTransactions.Type = t
	postTransactions.To = address
	postTransactions.Value = value
	postTransactions.Callback = url
	postTransactions.ExtraData = ""

	postTransactionsResults, err := postTransactions.Process("DEPOSIT")
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	// "/deposit/{type}/{addr}/{value}"
	postTransactionsResults.ExtraData = "/deposit/" + t + "/" + address + "/" + value

	io.WriteString(resp, common.JSONstring(postTransactionsResults))
	// Something like db POST?
	common.Gdb.ORDER([]byte(postTransactionsResults.Index), []byte(common.JSONstring(postTransactionsResults)))
}
