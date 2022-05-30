package jparse

import (
	"log"
	"testing"
)

type Product struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func TestJSON(t *testing.T) {

	p := &Product{}
	j := New()

	if err := j.ParseToData("/data/product.json", p); err != nil {
		log.Println(err)

	}
	log.Println(p.Name)
}
