package main

import (
	"fmt"
	"go-mastering-lecture/08-test/httpServerTest/handler"
	"net/http"
	"os"
)

type userErrorInterface interface {
	error  //组合自error接口
	Message() string
}
func main() {
	// 127.0.0.1:8080/list/07-errors/fib.txt
	http.HandleFunc("/", errWrapperV4(handler.HandleListFileV4))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(http.ResponseWriter, *http.Request) error

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
				return  //新增。记得return，不然会继续输出http.Error(w, http.StatusText(code), code)这一段
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




