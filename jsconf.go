package jsconf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"errors"
)

// ExistResult represent the result of the function Exist.
type ExistResult int

const (
	// Error represent an error in os.Stat().
	Error ExistResult = 1

	// IsDir if 'fileName' is a directory.
	IsDir ExistResult = 2

	// IsFile 'fileName' is a file.
	IsFile ExistResult = 3
)

// Exist return an ExistResult value depending on the 'fileName' Stat().
func Exist(fileName string) (ExistResult, error) {
	finfo, err := os.Stat(fileName)

	if err != nil {
		return Error, err
	}

	if finfo.IsDir() {
		return IsDir, nil
	}

	return IsFile, nil
}

// SaveToFile saves the 'data' received (an struct or map) in the file 'fileName'
func SaveToFile(fileName string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, bytes, 0644)
}

// LoadFromFile read a json file 'fileName' and saves its content in 'data'. 'data' must be
// a pointer to the destination struct (or map)
func LoadFromFile(fileName string, data interface{}) error {
	// If data is not a pointer, return an error
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return errors.New("pointer expected")
	}

	// Leemos el objeto json de un archivo.
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Transformamos el objeto json a una estructura y la guardamos en la variable 'data'.
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
}
