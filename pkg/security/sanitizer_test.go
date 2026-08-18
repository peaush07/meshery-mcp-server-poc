package security

import (
	"encoding/json"
	"testing"
)

func TestSanitizeMap_SensitiveKeys(t *testing.T) {
	input := map[string]interface{}{
		"username": "admin",
		"password": "supersecretpassword123",
		"token":    "bearer-token-xyz-789",
		"cluster":  "k8s-prod-cluster",
	}

	sanitized := SanitizeMap(input)

	if sanitized["password"] != "[REDACTED_SECRET]" {
		t.Errorf("expected password to be redacted, got %v", sanitized["password"])
	}
	if sanitized["token"] != "[REDACTED_SECRET]" {
		t.Errorf("expected token to be redacted, got %v", sanitized["token"])
	}
	if sanitized["username"] != "admin" {
		t.Errorf("expected username to be admin, got %v", sanitized["username"])
	}
}

func TestSanitizeMap_NestedStructures(t *testing.T) {
	input := map[string]interface{}{
		"metadata": map[string]interface{}{
			"name": "meshery-design",
			"credentials": map[string]interface{}{
				"kubeconfig": "apiVersion: v1...",
				"api_key":    "key-12345",
			},
		},
		"endpoints": []interface{}{
			map[string]interface{}{
				"url":   "https://mesh.local",
				"token": "secret-auth-token",
			},
		},
	}

	sanitized := SanitizeMap(input)

	metadata := sanitized["metadata"].(map[string]interface{})
	creds := metadata["credentials"].(map[string]interface{})

	if creds["kubeconfig"] != "[REDACTED_SECRET]" {
		t.Errorf("expected kubeconfig to be redacted, got %v", creds["kubeconfig"])
	}
	if creds["api_key"] != "[REDACTED_SECRET]" {
		t.Errorf("expected api_key to be redacted, got %v", creds["api_key"])
	}

	endpoints := sanitized["endpoints"].([]interface{})
	ep0 := endpoints[0].(map[string]interface{})
	if ep0["token"] != "[REDACTED_SECRET]" {
		t.Errorf("expected endpoint token to be redacted, got %v", ep0["token"])
	}
}

func TestSanitizeMap_CaseInsensitive(t *testing.T) {
	input := map[string]interface{}{
		"AuthToken":  "secret-123",
		"KubeConfig": "cluster-config",
		"PASSWORD":   "pass-456",
		"publicField": "public-val",
	}

	sanitized := SanitizeMap(input)

	if sanitized["AuthToken"] != "[REDACTED_SECRET]" {
		t.Errorf("expected AuthToken to be redacted, got %v", sanitized["AuthToken"])
	}
	if sanitized["KubeConfig"] != "[REDACTED_SECRET]" {
		t.Errorf("expected KubeConfig to be redacted, got %v", sanitized["KubeConfig"])
	}
	if sanitized["PASSWORD"] != "[REDACTED_SECRET]" {
		t.Errorf("expected PASSWORD to be redacted, got %v", sanitized["PASSWORD"])
	}
	if sanitized["publicField"] != "public-val" {
		t.Errorf("expected publicField to remain untouched, got %v", sanitized["publicField"])
	}
}

func TestSanitizeJSON_ValidJSON(t *testing.T) {
	rawJSON := []byte(`{"service":"meshery","secret_key":"topsecret123","status":"active"}`)

	sanitizedBytes, err := SanitizeJSON(rawJSON)
	if err != nil {
		t.Fatalf("unexpected error sanitizing JSON: %v", err)
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(sanitizedBytes, &resultMap); err != nil {
		t.Fatalf("failed to unmarshal sanitized JSON: %v", err)
	}

	if resultMap["secret_key"] != "[REDACTED_SECRET]" {
		t.Errorf("expected secret_key to be redacted in JSON output, got %v", resultMap["secret_key"])
	}
	if resultMap["status"] != "active" {
		t.Errorf("expected status to remain active, got %v", resultMap["status"])
	}
}
