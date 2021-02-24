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
