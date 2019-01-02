package common

import (
	"encoding/binary"
	"encoding/json"
	"sort"
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

// generate message to sign
func MapMessage(m map[string]interface{}) (msg string) {
	// To store the keys in slice in sorted order
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// To perform the opertion you want
	for _, k := range keys {
		msg += k

		e := m[k]

		if str, ok := e.(string); ok {
			msg += str
			continue
		}

		if i, ok := e.(int64); ok {
			msg += strconv.FormatInt(i, 10)
			continue
		}

		if i, ok := e.(int); ok {
			msg += strconv.Itoa(i)
			continue
		}

		if f, ok := e.(float64); ok {
			msg += Float64string(f)
			continue
		}

		if subm, ok := e.(map[string]interface{}); ok {
			msg += MapMessage(subm)
			continue
		}

		if a, ok := e.([]string); ok {
			for _, str := range a {
				msg += str
			}
			continue
		}

		if a, ok := e.([]int); ok {
			for _, i := range a {
				msg += strconv.Itoa(i)
			}
			continue
		}

		if a, ok := e.([]int64); ok {
			for _, i := range a {
				msg += strconv.FormatInt(i, 10)
			}
			continue
		}

		if a, ok := e.([]float64); ok {
			for _, f := range a {
				msg += Float64string(f)
			}
			continue
		}

		if a, ok := e.([]map[string]interface{}); ok {
			for _, subm := range a {
				msg += MapMessage(subm)
			}
			continue
		}

		if a, ok := e.([]interface{}); ok {
			msg += ArrayMessage(a)
			continue
		}

	}
	return msg
}

// generate message to sign
func ArrayMessage(a []interface{}) (msg string) {
	for _, e := range a {

		if str, ok := e.(string); ok {
			msg += str
			continue
		}

		if i, ok := e.(int64); ok {
			msg += strconv.FormatInt(i, 10)
			continue
		}

		if i, ok := e.(int); ok {
			msg += strconv.Itoa(i)
			continue
		}

		if f, ok := e.(float64); ok {
			msg += Float64string(f)
			continue
		}

		if subm, ok := e.(map[string]interface{}); ok {
			msg += MapMessage(subm)
			continue
		}

		if a, ok := e.([]string); ok {
			for _, str := range a {
				msg += str
			}
			continue
		}

		if a, ok := e.([]int); ok {
			for _, i := range a {
				msg += strconv.Itoa(i)
			}
			continue
		}

		if a, ok := e.([]int64); ok {
			for _, i := range a {
				msg += strconv.FormatInt(i, 10)
			}
			continue
		}

		if a, ok := e.([]float64); ok {
			for _, f := range a {
				msg += Float64string(f)
			}
			continue
		}

		if a, ok := e.([]map[string]interface{}); ok {
			for _, subm := range a {
				msg += MapMessage(subm)
			}
			continue
		}

		if a, ok := e.([]interface{}); ok {
			msg += ArrayMessage(a)
			continue
		}

	}
	return msg
}
