package read

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)
// 获取当前文件路径
func getCurrentDir() string  {
	dir , err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir,"\\","/",-1)
}

// 获取指定路径下的所有文件
func getAllFileName(dirName string)([]string,error)  {
	fileName := []string{}
	list , err := ioutil.ReadDir(dirName)

	if err != nil {
		fmt.Println(err)
	}
	// 排除文件夹和非指定的文件
	for _ , file := range list{
		isTxt := strings.HasSuffix(file.Name(),"txt")
		if ok := file.IsDir(); !ok && isTxt{
			fileName = append(fileName,file.Name())
		}
	}
	return fileName , nil
}
