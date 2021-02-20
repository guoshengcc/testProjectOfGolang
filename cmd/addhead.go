/*Package cmd ...
 * 向文件头中添加指定信息
**/
package cmd

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
func AddHeadMsg(fe FileEntity, msg string) error {
	// TODO
	return nil
}