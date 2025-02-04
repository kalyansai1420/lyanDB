package main

import (
	"encoding/json"
	"os"
)

func LoadDatabase(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil 
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&dataStore)
}

func main() {
	LoadDatabase("lyanDB.rdb") 
	StartServer("6379")
}