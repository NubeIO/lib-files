package jparse

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
)

type Json struct {
}

func New() *Json {
	return &Json{}
}

// ParseToData parses the JSON-encoded data and stores the result
// in the value pointed to by "data"
func (e *Json) ParseToData(file string, data interface{}) error {
	value, err := fileutils.ReadJSON(file)
	if err != nil {
		return fmt.Errorf("read: %s", err.Error())
	}
	if err := json.Unmarshal(value, data); err != nil {
		return fmt.Errorf("json unmarshal: %s", err.Error())
	}
	return nil
}
