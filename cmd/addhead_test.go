/* Copyright (c) 2021 Digital China Group Co.,Ltd
 * Licensed under the GNU General Public License, Version 3.0 (the "License").
 * You may not use this file except in compliance with the license.
 * You may obtain a copy of the license at
 *     https://www.gnu.org/licenses/
 *
 * This program is free; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.0 of the License.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
**/

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
