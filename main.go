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

package main

import (
	"flag"
	"fmt"
	"testproject/cmd"
)

func main() {
	var targetDirPath, configFilePath string
	flag.StringVar(&targetDirPath, "t", "", "target dir path")
	flag.StringVar(&configFilePath, "c", "./HeadContent.xml", "the haed content config XML file path,Default is './HeadContent.xml'")
	flag.Parse()
	fmt.Println("接收的外部参数信息：")
	fmt.Printf("targetDirPath:\t%s\nconfigFilePath:\t%s\n", targetDirPath, configFilePath)
	err := cmd.Run(targetDirPath, configFilePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ok")
	}
}
