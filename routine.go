package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/yushizhao/glasspool/common"
	"github.com/yushizhao/glasspool/order"
)

var state = []string{"holding", "online", "pending"}
var stateIndicator = 0

// orders may skip certain states, but will not revert after done/fail.
// tx state is not covered.
func orderStateUpdate() {
	ticker := time.NewTicker(common.UPDATETICKER * time.Second)
	defer ticker.Stop()
	for {
		select {

		case <-ticker.C:
			id := common.Gdb.ReadOrderStateUpdate()

			switch stateIndicator {
			case 3:
				for {
					idStr := strconv.FormatUint(id, 10)
					orderBytes := common.Gdb.QueryOrder(idStr)

					if orderBytes == nil {
						common.Gdb.WriteOrderStateUpdate(id)
						break
					}
					order := new(order.POSTTransactionsResult)
					json.Unmarshal(orderBytes, &order)
					// fail rate
					if id%common.FAILRATE == 0 {
						order.State = "failed"
					} else {
						// update more information?
						order.State = "done"
					}
					common.Gdb.ORDER([]byte(idStr), []byte(common.JSONstring(order)))
					id++

					urlBytes := common.Gdb.Callback(order.BizType, order.CoinType+order.To)
					returnBody := new(common.ReturnBody)
					returnBody.Normalize(order)
					log.Println(common.JSONstring(returnBody))
					callback(string(urlBytes), *returnBody)
				}

				stateIndicator = 0

			case 0, 1, 2:
				for {
					idStr := strconv.FormatUint(id, 10)
					orderBytes := common.Gdb.QueryOrder(idStr)

					if orderBytes == nil {

						break
					}
					order := new(order.POSTTransactionsResult)
					json.Unmarshal(orderBytes, &order)

					order.State = state[stateIndicator]

					common.Gdb.ORDER([]byte(idStr), []byte(common.JSONstring(order)))
					id++

					urlBytes := common.Gdb.Callback(order.BizType, order.CoinType+order.To)
					returnBody := new(common.ReturnBody)
					returnBody.Normalize(order)
					log.Println(common.JSONstring(returnBody))
					callback(string(urlBytes), *returnBody)
				}

				stateIndicator++

			}
		}
	}
}
