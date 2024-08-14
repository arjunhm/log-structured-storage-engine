package lsm

import "fmt"

type KeyValue struct {
	Key     string
	Value   string
	Size    uint32
	Deleted bool
}

func NewKeyValue(key, val string, del bool) *KeyValue {
	kv := &KeyValue{
		Key:     key,
		Value:   val,
		Deleted: del,
	}
	kv.Size = uint32(len(key)) + uint32(len(val))

	return kv
}


func (kv *KeyValue) Display() {
	fmt.Println(kv.Key + "-" + kv.Value);
}
