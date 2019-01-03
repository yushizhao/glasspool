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

		if tf, ok := e.(bool); ok {
			msg += strconv.FormatBool(tf)
		}

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

		if a, ok := e.([]bool); ok {
			for id, tf := range a {
				msg += strconv.Itoa(id)
				msg += strconv.FormatBool(tf)
			}
			continue
		}

		if a, ok := e.([]string); ok {
			for id, str := range a {
				msg += strconv.Itoa(id)
				msg += str
			}
			continue
		}

		if a, ok := e.([]int); ok {
			for id, i := range a {
				msg += strconv.Itoa(id)
				msg += strconv.Itoa(i)
			}
			continue
		}

		if a, ok := e.([]int64); ok {
			for id, i := range a {
				msg += strconv.Itoa(id)
				msg += strconv.FormatInt(i, 10)
			}
			continue
		}

		if a, ok := e.([]float64); ok {
			for id, f := range a {
				msg += strconv.Itoa(id)
				msg += Float64string(f)
			}
			continue
		}

		if a, ok := e.([]map[string]interface{}); ok {
			for id, subm := range a {
				msg += strconv.Itoa(id)
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
	for id, e := range a {

		msg += strconv.Itoa(id)

		if tf, ok := e.(bool); ok {
			msg += strconv.FormatBool(tf)
		}

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

		if a, ok := e.([]bool); ok {
			for id, tf := range a {
				msg += strconv.Itoa(id)
				msg += strconv.FormatBool(tf)
			}
			continue
		}

		if a, ok := e.([]string); ok {
			for id, str := range a {
				msg += strconv.Itoa(id)
				msg += str
			}
			continue
		}

		if a, ok := e.([]int); ok {
			for id, i := range a {
				msg += strconv.Itoa(id)
				msg += strconv.Itoa(i)
			}
			continue
		}

		if a, ok := e.([]int64); ok {
			for id, i := range a {
				msg += strconv.Itoa(id)
				msg += strconv.FormatInt(i, 10)
			}
			continue
		}

		if a, ok := e.([]float64); ok {
			for id, f := range a {
				msg += strconv.Itoa(id)
				msg += Float64string(f)
			}
			continue
		}

		if a, ok := e.([]map[string]interface{}); ok {
			for id, subm := range a {
				msg += strconv.Itoa(id)
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
