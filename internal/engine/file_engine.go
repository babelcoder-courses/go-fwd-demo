package engine

import (
	"encoding/json"
	"os"
	"runtime"
	"strings"

	"github.com/go-resty/resty/v2"
)

type FileEngine struct {
	URL      string
	FilePath string
}

func NewFileEngine(url, filePath string) *FileEngine {
	return &FileEngine{
		URL:      url,
		FilePath: filePath,
	}
}

func (e FileEngine) Dump() error {
	client := resty.New()
	_, err := client.R().SetOutput(e.FilePath).Get(e.URL)

	return err
}

func (e FileEngine) partialSearch(ch chan []Resource, term string, resources []Resource) {
	var result []Resource

	for _, resource := range resources {
		if strings.Contains(resource.Title, term) {
			result = append(result, resource)
		}
	}

	ch <- result
}

func (e FileEngine) Filter(term string) ([]Resource, error) {
	gorountineCount := runtime.NumCPU()
	ch := make(chan []Resource)
	var resources []Resource
	var result []Resource

	contentBytes, err := os.ReadFile(e.FilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contentBytes, &resources); err != nil {
		return nil, err
	}

	chunkLength := len(resources) / gorountineCount

	for i := 0; i < gorountineCount; i++ {
		if i == gorountineCount-1 {
			go e.partialSearch(ch, term, resources[i*chunkLength:])
		} else {
			go e.partialSearch(ch, term, resources[i*chunkLength:chunkLength*(i+1)])
		}
	}

	for i := 0; i < gorountineCount; i++ {
		result = append(result, <-ch...)
	}

	return result, nil
}
