package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)



const PREFIX  = "/list/"

func HandleListFileV4(w http.ResponseWriter, r *http.Request) error {
	//增加对URL Path的检查,检查Path中这个"list"是否是出现在最前面
	if strings.Index(r.URL.Path, PREFIX) != 0 {
		//自定义错误
		return userError("path必须以 "+ PREFIX + " 开头")
	}
	path := r.URL.Path[len(PREFIX):]	//也可以使用strings中的去除前缀方法
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = w.Write(all)
	return err
}

type userError string
func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}
