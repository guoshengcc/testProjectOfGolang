/*Package cmd ...
 * 向文件头中添加指定信息
**/
package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 指定配置文件的相对路径。
// 该配置文件中记录了各种不同类型文件对应的不同前缀内容(headContent)
var headContentConfigXMLPath = "../HeadContent.xml"

// HeadContent 各种类型文件对应的不同前缀内容（head content）
type HeadContent struct {
	XMLName                           xml.Name                                `xml:"headcontent"`
	Ignore                            []string                                `xml:"ignore"`
	ContentCategorizationBySuffixname []HeadContentCategorizationBySuffixname `xml:"categorizationbysuffixname"`
	ContentCategorizationByFilename   []HeadContentCategorizationByFilename   `xml:"categorizationbyfilename"`
}

// HeadContentCategorizationBySuffixname 根据文件后缀区分的前缀内容
type HeadContentCategorizationBySuffixname struct {
	Suffixname string `xml:"suffixname"`
	Content    string `xml:"content"`
}

// HeadContentCategorizationByFilename 根据文件名区分的前缀内容
type HeadContentCategorizationByFilename struct {
	Filename string `xml:"filename"`
	Content  string `xml:"content"`
}

// FileBaseInfo 文件的基本信息
type FileBaseInfo struct {
	FilePathStr string // 文件绝对路径
	FileNameStr string // 文件类型
}

// IsIgnore 根据文件/目录名称判断其是否应该忽略
// 返回值中true-应该忽略，false-不该忽略
func (headContent *HeadContent) IsIgnore(name string) bool {
	if headContent.Ignore == nil ||
		len(headContent.Ignore) < 1 {
		return false
	}

	for _, ig := range headContent.Ignore {
		if ig == name {
			return true
		}
	}

	return false
}

// loadHeadContentFormConfigXML 从配置文件中加载head content的配置信息
func loadHeadContentFormConfigXML() (*HeadContent, error) {
	finfo, err := os.Stat(headContentConfigXMLPath)
	if err != nil {
		return nil, err
	}
	if finfo.IsDir() {
		return nil, fmt.Errorf("错误，变量['headContentConfigXMLPath':%s]应该为一个xml文件路径", headContentConfigXMLPath)
	}

	xmlFile, err := os.Open(headContentConfigXMLPath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}
	headContent := HeadContent{}
	err = xml.Unmarshal(xmlData, &headContent)
	if err != nil {
		return nil, err
	}
	return &headContent, nil
}

// GetHeadContentStr 根据文件类型获取其对应的前缀内容
// 这个前缀内容存储在headContent中
// 优先使用文件名进行匹配，文件名匹配不到再使用文件名后缀匹配，如果都匹配不到则返回空字符串
// 文件名匹配区分大小写，文件名后缀匹配不区分大小写
func GetHeadContentStr(fe FileBaseInfo, headContent *HeadContent) (string, error) {
	// 优先匹配文件名
	if headContent != nil &&
		headContent.ContentCategorizationByFilename != nil &&
		len(headContent.ContentCategorizationByFilename) > 0 {
		for _, fn := range headContent.ContentCategorizationByFilename {
			if fe.FileNameStr == fn.Filename {
				return fn.Content, nil
			}
		}
	}

	// 文件名匹配不到，再通过文件后缀匹配
	if headContent != nil &&
		headContent.ContentCategorizationBySuffixname != nil &&
		len(headContent.ContentCategorizationBySuffixname) > 0 {
		fileNameSplitArray := strings.Split(fe.FileNameStr, ".")
		if len(fileNameSplitArray) < 1 {
			return "", nil
		}
		fileNameSuffix := fileNameSplitArray[len(fileNameSplitArray)-1]
		for _, fsn := range headContent.ContentCategorizationBySuffixname {
			if strings.EqualFold(fileNameSuffix, fsn.Suffixname) {
				return fsn.Content, nil
			}
		}
	}

	return "", nil
}

// Run 按照配置内容，将前缀内容添加到指定目录中的所有文件中
func Run(dirPath string) error {
	// 读取配置
	allHeadContent, err := loadHeadContentFormConfigXML()
	if err != nil {
		return err
	}
	// 读取所有的文件
	fileBaseInfos, err := GetAllFilePath(dirPath, allHeadContent.Ignore)
	if err != nil {
		return err
	}

	if len(fileBaseInfos) < 1 {
		return nil
	}

	for _, fbi := range fileBaseInfos {
		if allHeadContent.IsIgnore(fbi.FileNameStr) {
			continue
		}
		headContentStr, err := GetHeadContentStr(fbi, allHeadContent)
		if err != nil {
			return err
		}
		if len(headContentStr) < 1 {
			continue
		}
		// 向各个文件中添加前缀内容(head  Content)
		err1 := AddHeadMsg(fbi, headContentStr)
		if err1 != nil {
			return err1
		}
	}

	return nil
}

// GetAllFilePath 获取指定目录下面的所有文件路径
func GetAllFilePath(dirPath string, ignores []string) ([]FileBaseInfo, error) {
	fileBaseInfos := make([]FileBaseInfo, 0, 3)

	finfo, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}

	if !finfo.IsDir() {
		return nil, fmt.Errorf("错误，变量['dirPath':%s]应该为一个目录路径", headContentConfigXMLPath)
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// 跳过需要忽略的文件/目录
		isIgnore := false
		if ignores != nil && len(ignores) > 0 {
			for _, ig := range ignores {
				if file.Name() == ig {
					isIgnore = true
					break
				}
			}
		}
		if isIgnore {
			continue
		}

		if file.IsDir() {
			// 递归获取子目录中的文件信息
			subDirFileBaseInfos, err := GetAllFilePath(dirPath+string(os.PathSeparator)+file.Name(), ignores)
			if err != nil {
				return fileBaseInfos, err
			}
			fileBaseInfos = append(fileBaseInfos, subDirFileBaseInfos...)
		} else {
			fbi := FileBaseInfo{
				FilePathStr: dirPath + string(os.PathSeparator) + file.Name(),
				FileNameStr: file.Name(),
			}
			fileBaseInfos = append(fileBaseInfos, fbi)
		}
	}

	return fileBaseInfos, nil
}

// AddHeadMsg 向文件头添加信息
// 只针对UTF-8项目
func AddHeadMsg(fe FileBaseInfo, msg string) error {
	// 读取目标文件
	fileObj, err := os.Open(fe.FilePathStr) // ./hello.txt
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
			err := os.Rename("./temp.temp", fe.FilePathStr)
			if err != nil {
				fmt.Printf("rename file filed,err:%v", err)
				return err
			}
			fmt.Println("finished!!")
			return nil
		}
	}
}
