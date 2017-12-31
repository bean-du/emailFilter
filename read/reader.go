package read

import (
	"fmt"
)

type Files struct {
	file chan string
	result chan string
}

func GetNewFiles(maxNum int) *Files  {
	return &Files{
		file   : make(chan string , maxNum),
		result : make(chan string ,maxNum),
	}
}

func (f *Files)Read()  {
	go func(f *Files) {
		defer close(f.file)
		fileList , err := getAllFileName(getCurrentDir())
		if err !=nil {
			fmt.Println(err)
		}
		for _,file := range fileList{
			f.file <- file
		}
	}(f)
}

func (f *Files)Filter()  {
	defer close(f.result)
	for file := range f.file {
		strings , _ := filterEmail(file)

		for _ , str := range strings{
			f.result <- str
		}
	}
}

func (f *Files)Write()  {
	file := "Emails.txt"

	for str := range f.result{
		WriteFile(file , str + "\r\n")
	}
}


