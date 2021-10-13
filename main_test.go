package main

import (
	"os"
	"testing"
)

func TestFullFlowWithMock(t *testing.T) {
	os.Setenv("mode_mock", "true")
	main()
}
