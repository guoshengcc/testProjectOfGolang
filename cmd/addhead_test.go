package cmd

import (
	"fmt"
	"testing"
)

func TestGetAllFilePath(t *testing.T) {
	// dirPath := "C:\\GitHub\\gomock" //"C:\\GitHub\\gomock"
	// filePathArray, err := GetAllFilePath(dirPath, []string{".git", ".github"})

	// if err != nil {
	// 	t.Error(err)
	// }

	// // 4+3+2
	// if len(filePathArray) < 1 {
	// 	t.Error("错误")
	// } else {
	// 	t.Log(filePathArray)
	// }
}

func TestLoadHeadContentFormConfigXML(t *testing.T) {
	headContent, err := loadHeadContentFormConfigXML("../HeadContent.xml")
	if err != nil {
		t.Error(err)
	}
	str := headContent.ContentCategorizationBySuffixname[0].Content
	fmt.Print(str)
	if len(str) < 10 {
		t.Error("错误")
	}
}
