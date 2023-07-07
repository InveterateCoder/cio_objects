package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveJson(item any, path string) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err = encoder.Encode(item); err != nil {
		panic(err)
	}
}

func ReadJson[T any](path string) T {
	var item T
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err = json.NewDecoder(file).Decode(&item); err != nil {
		panic(err)
	}
	return item
}
