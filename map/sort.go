package main

import (
	"encoding/json"
	"os"
	"sort"

	"github.com/Sirupsen/logrus"
)

type Location struct {
	Location map[int]uint16
}

type SortLocation struct {
	Location []int
}

func unmarshalFile(file string, obj interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	return dec.Decode(obj)
}

func encodeToFile(obj interface{}, file string) error {
	f, err := os.Create(file + ".tmp")
	if err != nil {
		logrus.Errorf("failed to create temp file: %s while encoding the data to file", file)
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(&obj); err != nil {
		logrus.Errorf("failed to encode the data to file: %s", f.Name())
		return err
	}

	if err := f.Close(); err != nil {
		logrus.Errorf("failed to close file after encoding to file: %s", f.Name())
		return err
	}
	return nil
}

func main() {
	var loc Location
	var sortLocation SortLocation
	if err := unmarshalFile("info3", &loc); err != nil {
		panic(err)
	}
	var keys []int
	for k := range loc.Location {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sortLocation.Location = keys
	if err := encodeToFile(&sortLocation, "info3"); err != nil {
		panic(err)
	}

}
