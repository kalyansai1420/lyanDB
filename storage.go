package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	dataStore = make(map[string]string)
	
	expiry    = make(map[string]time.Time)
	mu        sync.RWMutex
	dbMutex   sync.Mutex
)


func SaveDatabase(filename string) error {
	dbMutex.Lock() 
	defer dbMutex.Unlock()

	fmt.Println("Saving Database:", dataStore) 

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(dataStore)
}

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

func checkIfKeyExists(key string) bool {
	_, exists := dataStore[key]
	return exists
}

func deleteKey(key string) bool {
	if _, exists := dataStore[key]; exists {
		delete(dataStore,key) 
		return true
	}
	return false
}
