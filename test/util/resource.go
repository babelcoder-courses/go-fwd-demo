package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Resource struct {
	ID    uint   `json:"id"`
	Title string `json:"title" faker:"sentence"`
}

func GenerateResource(id uint) Resource {
	var resource Resource

	faker.FakeData(&resource)
	resource.ID = id

	return resource
}

func GenerateResources(count int, pre ...Resource) ([]byte, []Resource) {
	var resources []Resource

	for i := 0; i < count; i++ {
		resource := GenerateResource(uint(i + 1))
		resources = append(resources, resource)
	}

	resources = append(resources, pre...)
	bytes, _ := json.Marshal(resources)

	return bytes, resources
}

func GenerateFoundResources(term string) []Resource {
	var found []Resource
	rand.Seed(time.Now().Unix())

	for i := 0; i < rand.Intn(10); i++ {
		resource := GenerateResource(uint(i + 1))
		resource.Title = fmt.Sprintf("%s %s %s", faker.Word(), term, faker.Word())
		found = append(found, resource)
	}

	return found
}
