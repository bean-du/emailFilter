package read

import (
	"os"
	"io/ioutil"
	"regexp"
)

func filterEmail(fileName string) (str []string , err error)  {
	f , err := os.Open(fileName)
	if err != nil {
		return nil , err
	}

	s , err := ioutil.ReadAll(f)
	if err != nil {
		return  nil ,err
	}

	reg := regexp.MustCompile(`[\d\w]+@[\d\w]+.[\w]{2,5}`)

	res := reg.FindAllString(string(s),-1)

	return res ,nil
}
