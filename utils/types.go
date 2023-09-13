package utils

import (
	fmt "fmt"

	"github.com/ethereum/go-ethereum/common"
)

// MemDB is an ethdb.KeyValueReader implementation which is not thread safe and
// assumes that all keys are common.Hash.
type MemDB struct {
	kvs map[common.Hash][]byte
}

// NewMemDB creates a new empty MemDB
func NewMemDB() *MemDB {
	return &MemDB{
		kvs: make(map[common.Hash][]byte),
	}
}

// Has returns true if the MemBD contains the key
func (m *MemDB) Has(key []byte) (bool, error) {
	h := common.BytesToHash(key)
	_, ok := m.kvs[h]
	return ok, nil
}

// Get returns the value of the key, or nil if it's not found
func (m *MemDB) Get(key []byte) ([]byte, error) {
	h := common.BytesToHash(key)
	value, ok := m.kvs[h]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return value, nil
}

// Put sets or updates the value at key
func (m *MemDB) Put(key []byte, value []byte) {
	h := common.BytesToHash(key)
	m.kvs[h] = value
}
