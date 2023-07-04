package testutils

import "encoding/json"

func StringFromMap(m map[string]interface{}) string {
	b, _ := json.Marshal(m)
	return (string(b))
}
