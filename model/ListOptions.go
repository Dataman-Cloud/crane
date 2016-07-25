package model

type ListOptions struct {
	Offset uint64
	Limit  uint64

	Filter map[string]interface{}
}
