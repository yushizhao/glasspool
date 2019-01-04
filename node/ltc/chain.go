package ltc

import (
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

const (
	HEIGHT    = 1555369
	TIMESTAMP = 1546585663000
	TIMESPAN  = 120000
	SEED      = "Glasspool LTC HD Master Key Seed"
	COINBASE  = "LZyzAuBPDxC1UdisvxSMTZHJ1Gc2o7ic4A"
)

var (
	COPY          = chaincfg.MainNetParams
	NET           = &COPY
	CurrentHeight int64
	CurrentTime   int64
	CurrentHash   string
	Fee           = 0.001788
)

func Init() {
	Master, _ = hdkeychain.NewMaster([]byte(SEED), NET)
	ChildID = 1 // 0 for coinbase

	incre := (time.Now().Unix()*1000 - TIMESTAMP) / TIMESPAN
	CurrentHeight = HEIGHT + incre
	CurrentTime = TIMESTAMP + incre*TIMESPAN
	CurrentHash = Hash(itob(CurrentHeight))

	// Hack
	NET.PubKeyHashAddrID = 0x30
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
