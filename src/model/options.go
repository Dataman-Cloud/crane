package model

import "encoding/json"

type ListOptions struct {
	Offset uint64
	Limit  uint64

	Filter map[string]interface{}
}

type UpdateOptions struct {
	Method  string          `json:"Method"`
	Options json.RawMessage `json:"Options"`
}
