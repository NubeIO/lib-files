package jparse

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-dirs/fileutils"
)

type Json struct {
}

func New() *Json {
	return &Json{}
}

// ParseToData parses the JSON-encoded data and stores the result
// in the value pointed to by "data"
func (e *Json) ParseToData(file string, data interface{}) error {
	f := fileutils.New()
	value, err := f.ReadJSON(file)
	if err != nil {
		return fmt.Errorf("read: %s", err.Error())
	}
	if err := json.Unmarshal(value, data); err != nil {
		return fmt.Errorf("json unmarshal: %s", err.Error())
	}
	return nil
}
