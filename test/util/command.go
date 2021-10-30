package util

import (
	"bytes"
	"fwd-search-api/cmd"
	"io"
)

func ExecuteCommand(args ...string) string {
	b := new(bytes.Buffer)
	root := cmd.RootCmd

	root.SetOut(b)
	root.SetArgs(args)
	cmd.Execute()
	out, _ := io.ReadAll(b)

	return string(out)
}
