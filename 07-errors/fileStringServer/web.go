package main

import (
	"context"
	"github.com/golang/gddo/log"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 127.0.0.1:8080/list/07-errors/fib.txt
	//http.HandleFunc("/list/", handleListFile)
	http.HandleFunc("/list/", errWrapper(handleListFileV2))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleListFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/list/"):]	//也可以使用strings中的去除前缀方法
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(all)
	if err != nil {
		panic(err)
	}
}

type appHandler func(http.ResponseWriter, *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			//log是服务端终端打印的日志，给维护人员看的
			log.Warn(context.TODO(), "Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):	//请求的文件不存在
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			//http.StatusText(http.StatusNotFound隐藏内部错误，只告诉外部一个错误代码的简单信息
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func handleListFileV2(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]	//也可以使用strings中的去除前缀方法
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
	if err != nil {
		return err
	}
	return nil
}
