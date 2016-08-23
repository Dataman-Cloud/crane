package model

import "encoding/json"

type UpdateOptions struct {
	Method  string          `json:"Method"`
	Options json.RawMessage `json:"Options"`
}
