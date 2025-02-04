package main

import (
	"sync"
	"time"
)

var (
	dataStore = make(map[string]string)
	expiry    = make(map[string]time.Time)
	mu        sync.RWMutex
)

func SetKey(key, value string, expiryTime time.Time) {
	mu.Lock()
	defer mu.Unlock()
	dataStore[key] = value
	if !expiryTime.IsZero() {
		expiry[key] = expiryTime
	} else {
		delete(expiry, key)
	}
}

func GetKey(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	value, exists := dataStore[key]
	if !exists {
		return "", false
	}

	expiryTime, hasExpiry := expiry[key]
	if hasExpiry && time.Now().After(expiryTime) {
		mu.RUnlock()
		mu.Lock()
		delete(dataStore, key)
		delete(expiry, key)
		mu.Unlock()
		mu.RLock()
		return "", false
	}
	return value, true
}
