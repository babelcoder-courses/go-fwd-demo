package engine_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"fwd-search-api/test/util"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

type FileEngineTestSuite struct {
	suite.Suite
}

func (suite *FileEngineTestSuite) TestDump() {
	content, _ := util.GenerateResources(5)
	server := util.StartServer(content)
	defer server.Close()

	file, _ := os.CreateTemp("", "*")
	fileName := file.Name()
	defer os.Remove(fileName)

	util.ExecuteCommand(
		"dump",
		"--url="+server.URL,
		"--file-path="+fileName,
	)

	bytes, _ := os.ReadFile(fileName)
	suite.Assert().Equal(content, bytes)
}

func (suite *FileEngineTestSuite) TestFilter() {
	term := faker.Word()
	found := util.GenerateFoundResources(term)
	resources, _ := util.GenerateResources(10, found...)
	file, _ := os.CreateTemp("", "*")
	fileName := file.Name()
	defer os.Remove(fileName)
	file.Write(resources)

	output := util.ExecuteCommand(
		"filter",
		term,
		"--file-path="+fileName,
	)

	var expected bytes.Buffer
	for _, r := range found {
		fmt.Fprintf(&expected, "%d: %s\n", r.ID, r.Title)
	}

	suite.Assert().ElementsMatch(
		strings.Split(expected.String(), "\n"),
		strings.Split(output, "\n"),
	)
}

func TestFileEngine(t *testing.T) {
	suite.Run(t, new(FileEngineTestSuite))
}
