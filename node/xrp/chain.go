package xrp

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

const (
	HEIGHT    = 44078639
	TIMESTAMP = 1546235700000
	TIMESPAN  = 250
	SEED      = "Glasspool XRP HD Master Key Seed"
	COINBASE  = "rj8YaPZfYkRtx4xZdws26afjAiXmjBWf3B"
)

var (
	NET           = &chaincfg.MainNetParams
	CurrentHeight int64
	CurrentTime   int64
	CurrentHash   string
	Fee           = 0.0005
)

func Init() {
	Master, _ = hdkeychain.NewMaster([]byte(SEED), NET)
	ChildID = 1 // 0 for coinbase

	incre := (time.Now().Unix()*1000 - TIMESTAMP) / TIMESPAN
	CurrentHeight = HEIGHT + incre
	CurrentTime = TIMESTAMP + incre*TIMESPAN
	CurrentHash = Hash(itob(CurrentHeight))

	// Hack
	NET.PubKeyHashAddrID = 0x7b
}

// linear estimate
func UpdateBlockNumber() {
	span := time.Now().Unix()*1000 - CurrentTime
	if span > TIMESPAN {
		incre := span / TIMESPAN
		CurrentHeight = CurrentHeight + incre
		CurrentTime = CurrentTime + incre*TIMESPAN
		CurrentHash = HashInt(CurrentHeight)
	}
}
