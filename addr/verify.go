package addr

import "github.com/yushizhao/glasspool/common"

type POSTAddressesVerify struct {
	Type string `json:"type"`
}

type POSTAddressesVerifyResult struct {
	Address string `json:"address"`
	// Callback   string `json:"callback"`
	Valid bool `json:"valid"`
}

func (input POSTAddressesVerify) Process(a string) (output POSTAddressesVerifyResult) {
	output.Address = a
	url := common.Gdb.Callback("DEPOSIT", input.Type+a)
	if url == nil {
		output.Valid = false
	} else {
		output.Valid = true
	}
	return output
}
