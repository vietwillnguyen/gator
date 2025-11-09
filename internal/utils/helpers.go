// internal/utils/helpers.go
package utils

import "encoding/json"

func ToJSON(v any) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data) + "\n"
}
