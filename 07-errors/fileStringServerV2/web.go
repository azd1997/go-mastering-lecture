package main

import (
	"context"
	"fmt"
	"github.com/golang/gddo/log"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userErrorInterface interface {
	error  //组合自error接口
	Message() string
}
func main() {
	// 127.0.0.1:8080/list/07-errors/fib.txt
	//http.HandleFunc("/list/", handleListFile)
	http.HandleFunc("/", errWrapperV4(handleListFileV4))
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

func errWrapperV2(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//对handleFunc的panic作recover保护
		defer func() {
			r := recover()  //获取panic值
			fmt.Printf("panic: %v\n", r)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}()

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

func errWrapperV3(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//对handleFunc的panic作recover保护
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic: %v\n", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

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

func errWrapperV4(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//对handleFunc的panic作recover保护
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic: %v\n", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		err := handler(w, r)
		if err != nil {
			fmt.Println("出错：", err.Error())
			//增加对自定义错误的区别处理
			if userErr, ok := err.(userErrorInterface); ok {
				//直接把自定义的错误返回给用户，因为是需要他们知道这些信息
				http.Error(w, userErr.Message(), http.StatusBadRequest)
			}

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
	return err
}

const PREFIX  = "/list/"
func handleListFileV3(w http.ResponseWriter, r *http.Request) error {
	//增加对URL Path的检查,检查Path中这个"list"是否是出现在最前面
	if strings.Index(r.URL.Path, PREFIX) != 0 {
		return errors.New("path必须以 "+ PREFIX + " 开头")
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
func handleListFileV4(w http.ResponseWriter, r *http.Request) error {
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
