package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	jsonData, err := json.MarshalIndent(*c, "", "    ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(jsonData))
}
