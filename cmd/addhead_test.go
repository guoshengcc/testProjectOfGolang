package cmd

import (
	"fmt"
	"testing"
)

func TestGetAllFilePath(t *testing.T) {
	dirPath := "C:\\tmp"
	filePathArray, err := GetAllFilePath(dirPath, make([]FileBaseInfo, 0, 3))

	if err != nil {
		t.Error(err)
	}

	// 4+3+2
	if len(filePathArray) != 9 {
		t.Error("错误")
	}
}

func TestLoadHeadContentFormConfigXML(t *testing.T) {
	headContent, err := loadHeadContentFormConfigXML()
	if err != nil {
		t.Error(err)
	}
	str := headContent.ContentCategorizationBySuffixname[0].Content
	fmt.Print(str)
	if len(str) < 10 {
		t.Error("错误")
	}
}
