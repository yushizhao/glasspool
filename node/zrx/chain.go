package zrx

import (
	"time"
)

const (
	HEIGHT    = 6614791
	TIMESTAMP = 1540921599000
	TIMESPAN  = 5432
	COINBASE  = "0xb1E52C1f985971fF2C00a33260925df5190E0153"
)

var (
	CurrentHeight int64
	CurrentTime   int64
	CurrentHash   string
	Fee           = 0.001281
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
