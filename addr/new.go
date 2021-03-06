package addr

import (
	"fmt"

	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/node/btc"
	"github.com/yushizhao/glasspool/node/cyb"
	"github.com/yushizhao/glasspool/node/eos"
	"github.com/yushizhao/glasspool/node/eth"
	"github.com/yushizhao/glasspool/node/gnt"
	"github.com/yushizhao/glasspool/node/ltc"
	"github.com/yushizhao/glasspool/node/usdc"
	"github.com/yushizhao/glasspool/node/usdt"
	"github.com/yushizhao/glasspool/node/vet"
	"github.com/yushizhao/glasspool/node/xrp"
	"github.com/yushizhao/glasspool/node/zrx"
	"github.com/yushizhao/glasspool/order"
)

// order key alphabetically for signing
// POST /api/v1/addresses/new
type POSTAddressesNew struct {
	Callback string `json:"callback"`
	Type     string `json:"type"`
}

type POSTAddressesNewResult struct {
	Address string `json:"address"`
	State   string `json:"state"`
	Type    string `json:"type"`
}

// type POSTAddressesNewError struct {
// 	Message string `json:"message"`
// 	Status  int    `json:"status"`
// }

func (input POSTAddressesNew) Process() (output POSTAddressesNewResult, err error) {
	output.Type = input.Type
	switch input.Type {
	case "ETH":
		secret, address, err := eth.GenerateAddress()
		common.Gdb.WALLET([]byte(input.Type+address), secret)
		output.Address = address
		output.State = "used"
		return output, err
	case "VET":
		secret, address, err := vet.GenerateAddress()
		common.Gdb.WALLET([]byte(input.Type+address), secret)
		output.Address = address
		output.State = "used"
		return output, err
	case "USDC":
		secret, address, err := usdc.GenerateAddress()
		common.Gdb.WALLET([]byte(input.Type+address), secret)
		output.Address = address
		output.State = "used"
		return output, err
	case "GNT":
		secret, address, err := gnt.GenerateAddress()
		common.Gdb.WALLET([]byte(input.Type+address), secret)
		output.Address = address
		output.State = "used"
		return output, err
	case "ZRX":
		secret, address, err := zrx.GenerateAddress()
		common.Gdb.WALLET([]byte(input.Type+address), secret)
		output.Address = address
		output.State = "used"
		return output, err
	case "BTC":
		// fix seed HD
		address, err := btc.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, err
	case "LTC":
		// fix seed HD
		address, err := ltc.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, err
	case "USDT":
		// fix seed HD
		address, err := usdt.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, err
	case "CYB":
		address := cyb.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, nil
	case "EOS":
		address := eos.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, nil
	case "XRP":
		address := xrp.GenerateAddress()
		output.Address = address
		output.State = "used"
		return output, nil
	default:
		return output, nil
	}
}

func (input POSTAddressesNewResult) AutoDeposit() (tx order.TxData, err error) {
	switch input.Type {
	case "ETH":
		eth.UpdateBlockNumber()
		tx.TimestampBegin = eth.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = eth.CurrentHeight
		tx.BlockHash = eth.CurrentHash
		tx.Hash = eth.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / eth.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = eth.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: eth.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "VET":
		vet.UpdateBlockNumber()
		tx.TimestampBegin = vet.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = vet.CurrentHeight
		tx.BlockHash = vet.CurrentHash
		tx.Hash = vet.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / vet.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = vet.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: vet.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "USDC":
		usdc.UpdateBlockNumber()
		tx.TimestampBegin = usdc.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = usdc.CurrentHeight
		tx.BlockHash = usdc.CurrentHash
		tx.Hash = usdc.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / usdc.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = usdc.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: usdc.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "GNT":
		gnt.UpdateBlockNumber()
		tx.TimestampBegin = gnt.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = gnt.CurrentHeight
		tx.BlockHash = gnt.CurrentHash
		tx.Hash = gnt.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / gnt.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = gnt.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: gnt.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "ZRX":
		zrx.UpdateBlockNumber()
		tx.TimestampBegin = zrx.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = zrx.CurrentHeight
		tx.BlockHash = zrx.CurrentHash
		tx.Hash = zrx.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / zrx.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = zrx.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: zrx.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "BTC":
		btc.UpdateBlockNumber()
		tx.TimestampBegin = btc.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = btc.CurrentHeight
		tx.BlockHash = btc.CurrentHash
		tx.Hash = btc.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / btc.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = btc.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: btc.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "LTC":
		ltc.UpdateBlockNumber()
		tx.TimestampBegin = ltc.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = ltc.CurrentHeight
		tx.BlockHash = ltc.CurrentHash
		tx.Hash = ltc.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / ltc.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = ltc.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: ltc.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "USDT":
		usdt.UpdateBlockNumber()
		tx.TimestampBegin = usdt.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = usdt.CurrentHeight
		tx.BlockHash = usdt.CurrentHash
		tx.Hash = usdt.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / usdt.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = usdt.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: usdt.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "CYB":
		cyb.UpdateBlockNumber()
		tx.TimestampBegin = cyb.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = cyb.CurrentHeight
		tx.BlockHash = cyb.CurrentHash
		tx.Hash = cyb.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / cyb.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = cyb.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: cyb.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "EOS":
		eos.UpdateBlockNumber()
		tx.TimestampBegin = eos.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = eos.CurrentHeight
		tx.BlockHash = eos.CurrentHash
		tx.Hash = eos.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / eos.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = eos.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: eos.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	case "XRP":
		xrp.UpdateBlockNumber()
		tx.TimestampBegin = xrp.CurrentTime
		tx.TimestampFinish = tx.TimestampBegin
		tx.BlockNumber = xrp.CurrentHeight
		tx.BlockHash = xrp.CurrentHash
		tx.Hash = xrp.HashInt(tx.TimestampBegin)
		tx.Confirmations = 0

		// pseudo random
		valueTo := float64(tx.TimestampBegin%tx.BlockNumber) / xrp.HEIGHT
		txpointTo := order.TxPoint{Address: input.Address, Value: common.Float64string(valueTo)}
		tx.To = []order.TxPoint{txpointTo}

		tx.Fee = xrp.Fee

		valueFrom := valueTo + tx.Fee
		txpointFrom := order.TxPoint{Address: xrp.COINBASE, Value: common.Float64string(valueFrom)}
		tx.From = []order.TxPoint{txpointFrom}

		tx.Type = input.Type

	default:
		return tx, fmt.Errorf("unexpected type: %v", input.Type)
	}

	// Callback?
	return tx, err
}
