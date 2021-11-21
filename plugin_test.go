package main

import (
	"os"
	"testing"

	"github.com/docker/docker/daemon/graphdriver/copy"
)

func TestPlugin(t *testing.T) {
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create("test_dir/file")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	_, err = file.WriteString("This is a test file")
	if err != nil {
		t.Error(err)
	}

	err = copy.DirCopy("test_dir", "dst_dir", copy.Content, false)
	if err != nil {
		t.Error(err)
	}
}
