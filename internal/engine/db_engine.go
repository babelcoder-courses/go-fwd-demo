package engine

import "fmt"

type DBEngine struct {
	URL string
}

func NewDBEngine(url string) *DBEngine {
	return &DBEngine{URL: url}
}

func (e DBEngine) Dump() error {
	fmt.Println("DB Dump")
	return nil
}

func (e DBEngine) Filter(term string) ([]Resource, error) {
	return nil, nil
}
