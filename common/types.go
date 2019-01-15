package common

import "encoding/json"

type POSTBody struct {
	Appid     string      `json:"appid"`
	Crypto    string      `json:"crypto"`
	Data      interface{} `json:"data"`
	Hash      string      `json:"hash"`
	Sig       string      `json:"sig"`
	Timestamp int64       `json:"timestamp"`
	Path      string      // implication
}

type ReturnBody struct {
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Status    int         `json:"status"`
	SendAgain bool        `json"sendAgain"`
}

type ReturnBodySig struct {
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Status    int         `json:"status"`
	Crypto    string      `json:"crypto"`
	Timestamp int64       `json:"timestamp"`
	Sig       SigJSON     `json:"sig"`
}

// parse data
func (it *POSTBody) Extract(data interface{}) {
	dataBytes, _ := json.Marshal(it.Data)
	json.Unmarshal(dataBytes, &data)
}

// normal post
func (it *POSTBody) Normalize() {
	it.Appid = "app"
	it.Crypto = "ecc"
	it.Hash = "sha3"
}

// parse result
func (it *ReturnBody) Extract(res interface{}) {
	resBytes, _ := json.Marshal(it.Result)
	json.Unmarshal(resBytes, &res)
}

// normal response
func (it *ReturnBody) Normalize(res interface{}) {
	it.Message = "OK"
	it.Result = res
	it.Status = 0
	it.SendAgain = false
}
