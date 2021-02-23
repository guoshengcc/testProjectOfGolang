/*Package cmd ...
 * 向文件头中添加指定信息
**/
package cmd

import (
	"fmt"
	"os"
)

// FileType 文件类型
type FileType uint32

// 枚举 各种文件类型
const (
	GO FileType = iota
	CS
	JS
	CSS
	YML
	XML
)

// FileEntity 文件的描述
type FileEntity struct {
	filePath string   // 文件绝对路径
	fileType FileType // 文件类型
}

// AddHeadMsg 向文件头添加信息
// 只针对UTF-8项目
func AddHeadMsg(fe FileEntity, msg string) error {
	// 读取目标文件
	fileObj, err := os.Open(fe.filePath) // ./hello.txt
	if err != nil {
		fmt.Printf("open file filed ,err :%v", err)
		return err
	}

	// 创建或打开缓存文件
	tempObj, err := os.OpenFile("./temp.temp", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file filed,err:%v", err)
		return err
	}
	_, err2 := tempObj.Write([]byte(msg)) // 向文件末尾写入，此时的文件末尾实为文件开头
	if err2 != nil {
		fmt.Printf("write file filed ,err:%v", err)
		return err2
	}
	var temp [128]byte
	for {
		n, err := fileObj.Read(temp[:])
		if err != nil {
			fmt.Printf("read file filed ,err :%v", err)
			return err
		}
		_, err1 := tempObj.Write(temp[:n])
		if err1 != nil {
			fmt.Printf("write file filed ,err:%v", err)
			return err1
		}
		if n < 128 {
			fileObj.Close()
			tempObj.Close()
			err := os.Rename("./temp.temp", fe.filePath)
			if err != nil {
				fmt.Printf("rename file filed,err:%v", err)
				return err
			}
			fmt.Println("finished!!")
			return nil
		}
	}
}
