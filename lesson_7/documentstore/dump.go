package documentstore

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func NewStoreFromDump(dump []byte) (*Store, error) {
	var store Store
	if err := json.Unmarshal(dump, &store); err != nil {
		return nil, errors.New("failed to unmarshal dump")
	}
	return &store, nil
}

func (s *Store) Dump() ([]byte, error) {
	dump, err := json.Marshal(s)
	if err != nil {
		return nil, errors.New("failed to marshal store")
	}
	return dump, nil
}

func NewStoreFromFile(filename string) (*Store, error) {
	dump, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewStoreFromDump(dump)
}

func (s *Store) DumpToFile(filename string) error {
	dump, err := s.Dump()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, dump, os.ModePerm)
}
