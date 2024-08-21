package filesoperate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 文件常规操作
func FileOperatemMain() {
	filepath := "./testDir"
	toFile(filepath)

}

func toFile(dirPath string) {
	content := "Hello, this is a new file content." // 新文件的内容

	// 检查路径是否存在，并创建目录（如果不存在）
	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 路径不存在，创建文件夹
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				fmt.Println("创建目录失败： ", err)
				//panic(err) // 如果创建目录失败，则panic
			}
		} else {
			fmt.Println("目录已存在")
		}
	} else if !info.IsDir() {
		fmt.Println("目录已存在")
		//panic("指定的路径不是一个目录")
	}

	// 路径存在且是目录，删除所有文件
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("读取目录失败")
		//panic(err) // 如果读取目录失败，则panic
	}

	for _, entry := range entries {
		// 忽略目录，只删除文件
		if !entry.IsDir() {
			filePath := filepath.Join(dirPath, entry.Name())
			if err := os.Remove(filePath); err != nil {
				//panic(err) // 如果删除文件失败，则panic
				fmt.Println("删除文件目录失败")
			}
		}
	}

	// 创建新文件并写入内容
	nowTime := time.Now() // 更正变量名为 nowTime
	timeStr := nowTime.Format("20060102150405")
	newFileName := "newFile_" + timeStr + ".txt" // 移除等号两侧的空格
	fmt.Println("文件名称: ", newFileName)
	newFilePath := filepath.Join(dirPath, newFileName)
	if err := os.WriteFile(newFilePath, []byte(content), 0644); err != nil {
		fmt.Println("文件写入失败")
		//panic(err) // 如果写入文件失败，则panic
	}
}
