package order

import (
	"fmt"
	"strconv"

	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/node/btc"
	"github.com/yushizhao/glasspool/node/eos"
	"github.com/yushizhao/glasspool/node/eth"
	"github.com/yushizhao/glasspool/node/usdc"
	"github.com/yushizhao/glasspool/node/usdt"
)

type POSTTransactions struct {
	Callback  string `json:"callback"`
	ExtraData string `json:"extraData"`
	To        string `json:"to"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

// used as order as well
type POSTTransactionsResult struct {
	BizType       string `json:"bizType"`
	CoinType      string `json:"coinType"`
	Confirmations int    `json:"confirmations"`
	CreateAt      int64  `json:"create_at"`
	Data          TxData `json:"data"`
	ExtraData     string `json:"extraData"`
	Fee           string `json:"fee"`
	Hash          string `json:"hash"`
	Index         string `json:"id"`
	State         string `json:"state"`
	To            string `json:"to"`
	UpdateAt      int64  `json:"update_at"`
	Value         string `json:"value"`
}

type TxData struct {
	BlockHash       string    `json:"blockHash"`
	BlockNumber     int64     `json:"blockNumber"`
	Confirmations   int       `json:"confirmations"`
	Fee             float64   `json:"fee"`
	From            []TxPoint `json:"from"`
	Hash            string    `json:"hash"`
	State           string    `json:"state"`
	TimestampBegin  int64     `json:"timestampBegin"`
	TimestampFinish int64     `json:"timestampFinish"`
	To              []TxPoint `json:"to"`
	Type            string    `json:"type"`
}

type TxPoint struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

func (input POSTTransactions) Process(bizType string) (output POSTTransactionsResult, err error) {
	txdata := new(TxData)
	txdata.State = "init"

	switch input.Type {
	case "ETH":
		txdata.Type = input.Type

		eth.UpdateBlockNumber()
		txdata.TimestampBegin = eth.CurrentTime
		txdata.TimestampFinish = txdata.TimestampBegin
		txdata.BlockNumber = eth.CurrentHeight
		txdata.BlockHash = eth.CurrentHash
		txdata.Hash = eth.HashInt(txdata.TimestampBegin)
		txdata.Confirmations = 0

		txpointTo := TxPoint{input.To, input.Value}
		txdata.To = []TxPoint{txpointTo}

		txdata.Fee = eth.Fee

		v, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return output, err
		}
		valueFrom := v + txdata.Fee
		txpointFrom := TxPoint{Address: eth.COINBASE, Value: common.Float64string(valueFrom)}
		txdata.From = []TxPoint{txpointFrom}

	case "USDC":
		txdata.Type = input.Type

		usdc.UpdateBlockNumber()
		txdata.TimestampBegin = usdc.CurrentTime
		txdata.TimestampFinish = txdata.TimestampBegin
		txdata.BlockNumber = usdc.CurrentHeight
		txdata.BlockHash = usdc.CurrentHash
		txdata.Hash = usdc.HashInt(txdata.TimestampBegin)
		txdata.Confirmations = 0

		txpointTo := TxPoint{input.To, input.Value}
		txdata.To = []TxPoint{txpointTo}

		txdata.Fee = usdc.Fee

		v, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return output, err
		}
		valueFrom := v + txdata.Fee
		txpointFrom := TxPoint{Address: usdc.COINBASE, Value: common.Float64string(valueFrom)}
		txdata.From = []TxPoint{txpointFrom}

	case "BTC":
		txdata.Type = input.Type

		btc.UpdateBlockNumber()
		txdata.TimestampBegin = btc.CurrentTime
		txdata.TimestampFinish = txdata.TimestampBegin
		txdata.BlockNumber = btc.CurrentHeight
		txdata.BlockHash = btc.CurrentHash
		txdata.Hash = btc.HashInt(txdata.TimestampBegin)
		txdata.Confirmations = 0
		// more plausible? e.g. N to N tx
		txpointTo := TxPoint{input.To, input.Value}
		txdata.To = []TxPoint{txpointTo}

		txdata.Fee = btc.Fee

		v, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return output, err
		}
		valueFrom := v + txdata.Fee
		txpointFrom := TxPoint{Address: btc.COINBASE, Value: common.Float64string(valueFrom)}
		txdata.From = []TxPoint{txpointFrom}

	case "USDT":
		txdata.Type = input.Type

		usdt.UpdateBlockNumber()
		txdata.TimestampBegin = usdt.CurrentTime
		txdata.TimestampFinish = txdata.TimestampBegin
		txdata.BlockNumber = usdt.CurrentHeight
		txdata.BlockHash = usdt.CurrentHash
		txdata.Hash = usdt.HashInt(txdata.TimestampBegin)
		txdata.Confirmations = 0
		// more plausible? e.g. N to N tx
		txpointTo := TxPoint{input.To, input.Value}
		txdata.To = []TxPoint{txpointTo}

		txdata.Fee = usdt.Fee

		v, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return output, err
		}
		valueFrom := v + txdata.Fee
		txpointFrom := TxPoint{Address: usdt.COINBASE, Value: common.Float64string(valueFrom)}
		txdata.From = []TxPoint{txpointFrom}

	// case "CYB":
	// 	txdata.Type = input.Type

	// 	cyb.UpdateBlockNumber()
	// 	txdata.TimestampBegin = cyb.CurrentTime
	// 	txdata.TimestampFinish = txdata.TimestampBegin
	// 	txdata.BlockNumber = cyb.CurrentHeight
	// 	txdata.BlockHash = cyb.CurrentHash
	// 	txdata.Hash = cyb.HashInt(txdata.TimestampBegin)
	// 	txdata.Confirmations = 0
	// 	// more plausible? e.g. N to N tx
	// 	txpointTo := TxPoint{input.To, input.Value}
	// 	txdata.To = []TxPoint{txpointTo}

	// 	txdata.Fee = cyb.Fee

	// 	v, err := strconv.ParseFloat(input.Value, 64)
	// 	if err != nil {
	// 		return output, err
	// 	}
	// 	valueFrom := v + txdata.Fee
	// 	txpointFrom := TxPoint{Address: cyb.COINBASE, Value: common.Float64string(valueFrom)}
	// 	txdata.From = []TxPoint{txpointFrom}

	case "EOS":
		txdata.Type = input.Type

		eos.UpdateBlockNumber()
		txdata.TimestampBegin = eos.CurrentTime
		txdata.TimestampFinish = txdata.TimestampBegin
		txdata.BlockNumber = eos.CurrentHeight
		txdata.BlockHash = eos.CurrentHash
		txdata.Hash = eos.HashInt(txdata.TimestampBegin)
		txdata.Confirmations = 0
		// more plausible? e.g. N to N tx
		txpointTo := TxPoint{input.To, input.Value}
		txdata.To = []TxPoint{txpointTo}

		txdata.Fee = eos.Fee

		v, err := strconv.ParseFloat(input.Value, 64)
		if err != nil {
			return output, err
		}
		valueFrom := v + txdata.Fee
		txpointFrom := TxPoint{Address: eos.COINBASE, Value: common.Float64string(valueFrom)}
		txdata.From = []TxPoint{txpointFrom}

	default:
		return output, fmt.Errorf("unexpected type: %v", input.Type)
	}
	output, err = txdata.Submit(bizType)

	return output, err
}

func (input TxData) Submit(bizType string) (output POSTTransactionsResult, err error) {
	switch input.Type {
	case "ETH", "Ethereum", "ethereum":
		output.CoinType = "ETH"
		// target the first point
		output.To = input.To[0].Address
		output.Value = input.To[0].Value
	case "USDC":
		output.CoinType = "USDC"
		// target the first point
		output.To = input.To[0].Address
		output.Value = input.To[0].Value
	case "BTC", "Bitcoin", "bitcoin":
		output.CoinType = "BTC"
		// target the first point
		output.To = input.To[0].Address
		output.Value = input.To[0].Value
	case "USDT":
		output.CoinType = "USDT"
		// target the first point
		output.To = input.To[0].Address
		output.Value = input.To[0].Value
	// case "CYB":
	// 	output.CoinType = "CYB"
	// 	// target the first point
	// 	output.To = input.To[0].Address
	// 	output.Value = input.To[0].Value
	case "EOS", "eos":
		output.CoinType = "EOS"
		// target the first point
		output.To = input.To[0].Address
		output.Value = input.To[0].Value
	default:
		return output, fmt.Errorf("unexpected type: %v", input.Type)
	}

	output.BizType = bizType
	output.Confirmations = input.Confirmations
	output.CreateAt = common.Timestamp()
	output.Data = input
	output.ExtraData = ""
	output.Fee = common.Float64string(input.Fee)
	output.Hash = input.Hash
	output.State = input.State
	output.UpdateAt = common.Timestamp()
	output.Index = common.Gdb.OrderID()
	return output, nil
}
