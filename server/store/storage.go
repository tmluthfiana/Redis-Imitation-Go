package store

import (
	"strconv"
	"sync"
)

// StorageMap maps a string key to a string value
type StorageMap struct {
	sync.RWMutex
	temporary map[string]string
}

// NewStorageMap creates a new string map
func NewStorageMap() *StorageMap {
	gm := StorageMap{
		temporary: make(map[string]string),
	}
	return &gm
}

// This command returns the value stored at the specified key
func (sm *StorageMap) Dump(key string) (value string, ok bool) {
	sm.RLock()
	defer sm.RUnlock()
	result, ok := sm.temporary[key]
	return result, ok
}

//Set a specified “key” with a “value”
func (sm *StorageMap) Set(key string, value string) bool {
	sm.Lock()
	defer sm.Unlock()
	sm.temporary[key] = value
	return true
}

// Increments the number stored at “key” by one
// If the key does not exist, it is set to 0 before performing the increment operation
// Return error if the value is not an integer
func (sm *StorageMap) Incr(key string) (result string, ok bool) {
	sm.Lock()
	defer sm.Unlock()

	val, ok := sm.temporary[key]
	if !ok {
		sm.temporary[key] = "0"
		val, _ = sm.temporary[key]
	}

	if res, err := strconv.Atoi(val); err == nil {
		newVal := strconv.Itoa(res + 1)
		sm.temporary[key] = newVal
		result = newVal
	} else {
		return "", false
	}

	return result, true
}

// Rename a given key with new key
func (sm *StorageMap) Rename(key string, newKey string) bool {

	if key == newKey {
		return true
	}
	sm.Lock()
	defer sm.Unlock()
	result, ok := sm.temporary[key]
	if ok {
		sm.temporary[newKey] = result
		delete(sm.temporary, key)
		return true
	}
	return false
}
