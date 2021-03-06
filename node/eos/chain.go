package eos

import "time"

const (
	ACCOUNT = "GLASSPOOL-EOS" // also used as coinbase

	HEIGHT    = 25999024
	TIMESTAMP = 1541719459000
	TIMESPAN  = 505
	COINBASE  = "GLASSPOOL-EOS"
)

var (
	CurrentHeight int64
	CurrentTime   int64
	CurrentHash   string
	Fee           = 0.0
)

func Init() {
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
