package security

import (
	"encoding/json"
	"strings"
)

var sensitiveKeys = []string{
	"token", "password", "secret", "private_key", "kubeconfig", "auth", "api_key",
}

// SanitizeJSON inspects JSON structures and redacts sensitive field values.
func SanitizeJSON(rawJSON []byte) ([]byte, error) {
	var genericMap map[string]interface{}
	if err := json.Unmarshal(rawJSON, &genericMap); err != nil {
		return rawJSON, nil // If not a JSON object, return as-is
	}

	sanitized := SanitizeMap(genericMap)
	return json.Marshal(sanitized)
}

// SanitizeMap recursively redacts sensitive key/value pairs in nested maps and slices.
func SanitizeMap(m map[string]interface{}) map[string]interface{} {
	redactMap(m)
	return m
}

func redactMap(m map[string]interface{}) {
	for k, v := range m {
		keyLower := strings.ToLower(k)
		isSensitive := false
		for _, sensitive := range sensitiveKeys {
			if strings.Contains(keyLower, sensitive) {
				m[k] = "[REDACTED_SECRET]"
				isSensitive = true
				break
			}
		}

		if isSensitive {
			continue
		}

		switch val := v.(type) {
		case map[string]interface{}:
			redactMap(val)
		case []interface{}:
			redactSlice(val)
		}
	}
}

func redactSlice(s []interface{}) {
	for _, item := range s {
		if childMap, ok := item.(map[string]interface{}); ok {
			redactMap(childMap)
		} else if childSlice, ok := item.([]interface{}); ok {
			redactSlice(childSlice)
		}
	}
}
