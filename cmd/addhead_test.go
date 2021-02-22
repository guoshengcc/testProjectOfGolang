package cmd

import "testing"

func TestAddHeadMsg(t *testing.T) {
	// TODO
	fe := FileEntity{
		filePath: "/tmp/test.go",
		fileType: GO,
	}

	AddHeadMsg(fe, "juse test")
}
