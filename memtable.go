package lsm

import "errors"

const MEMTABLE_LIMIT uint32 = 4096 * 1024 // 4MB 

type MemTable struct {
	Size uint32
	Entries []KeyValue
}

func NewMemTable() *MemTable {
	return &MemTable{
		Size: 0,
		Entries: make([]KeyValue, 0)
	}
}

func (mt *MemTable) clear() error {
	mt.Entries = []KeyValue{}
	mt.Size = 0;
	return nil
}

func (mt *MemTable) flush() error {
	return nil
}

func (mt *MemTable) isFull() bool {
	return mt.Size >= MEMTABLE_LIMIT
}

func (mt *MemTable) Get(key string) (string, error) {
	// scan from newest to oldest
	for i:= len(mt.Entries) - 1; i >= 0; i-- {
		if mt.Entries[i].Key == key {
			if mt.Entries[i].Deletedi == true {
				return "", nil
			} else {
				return  mt.Entries[i].Value, nil
			}
		}	
	}
	return "", errors.New("Key not found")
}

func (mt *MemTable) Put(key, val string, del bool) error {
	var err error
	kv := NewKeyValue(key, val, false)
	mt.Entries = append(mt.Entries, kv)
	if mt.isFull() {
		err = mt.flush()
		if err != nil {
			return err
		}
	}
	return nil
}

func (mt *MemTable) Delete(key, val string, del bool) error {
	err := mt.Put(key, val, true)
	if err != nil
		return err
	return nil
}
