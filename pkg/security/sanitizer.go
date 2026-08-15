package security

import (
	"encoding/json"
	"strings"
)

var sensitiveKeys = []string{
	"token", "password", "secret", "private_key", "kubeconfig", "auth",
}

// SanitizeJSON inspects JSON structures and redacts sensitive field values.
func SanitizeJSON(rawJSON []byte) ([]byte, error) {
	var genericMap map[string]interface{}
	if err := json.Unmarshal(rawJSON, &genericMap); err != nil {
		return rawJSON, nil // If not a JSON object, return as-is
	}

	redactMap(genericMap)
	return json.Marshal(genericMap)
}

func redactMap(m map[string]interface{}) {
	for k, v := range m {
		keyLower := strings.ToLower(k)
		for _, sensitive := range sensitiveKeys {
			if strings.Contains(keyLower, sensitive) {
				m[k] = "[REDACTED_SECRET]"
				break
			}
		}

		if childMap, ok := v.(map[string]interface{}); ok {
			redactMap(childMap)
		}
	}
}
