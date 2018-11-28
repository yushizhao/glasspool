package common

import (
	"encoding/binary"
	"encoding/json"
	"strconv"
	"time"
)

func JSONstring(a interface{}) string {
	res, err := json.Marshal(a)
	if err != nil {
		return "error"
	}
	return string(res)
}

// func JSONstringUnquote(a interface{}) string {
// 	res, err := strconv.Unquote(JSONstring(a))
// 	if err != nil {
// 		return "error"
// 	}
// 	return res
// }

// to millisecond
func Timestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Btoi returns v from an 8-byte big endian representation.
func Btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func Float64string(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}
