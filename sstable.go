package lsm

import "errors"

type SSTable struct {
	ID    uint32
	Data  []KeyValue
	Index map[string]uint32
}

func NewSSTable(id uint32) *SSTable {
	return &SSTable{
		ID:    id,
		Data:  make([]KeyValue, 0),
		Index: make(map[string]uint32),
	}
}
