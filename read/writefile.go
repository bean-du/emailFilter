package read

import (
	"os"
	"io"
)

// 检查文件是否存在
func CheckFileExist(fileName string)(exist bool)  {
	exist = true
	if _ , err := os.Stat(fileName); os.IsNotExist(err){
		exist = false
	}
	return
}

func WriteFile(filename, word string) error  {
	var f *os.File
	var err error
	//检查文件是否存在，如果存在就打开此文件，如果不存在就创建一个
	if CheckFileExist(filename) {
		f , err = os.OpenFile(filename , os.O_APPEND , 0666)
	}else {
		f , err = os.Create(filename)
	}
	// 将传入的字符串写入到文件中
	_ , err = io.WriteString(f , word)
	if err != nil {
		return err
	}
	return nil
}