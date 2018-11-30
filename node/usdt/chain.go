package usdt

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

const (
	HEIGHT    = 549253
	TIMESTAMP = 1541671496000
	TIMESPAN  = 600321
	SEED      = "Glasspool USDT HD Master Key Seed"
	COINBASE  = "15rky5uaJM8etj6UZ22wpjRAS9pGM6CC78"
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
