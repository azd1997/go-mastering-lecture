package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

//panic传入参数类型为interface{}；recover输出类型为interface{}
func tryRecover() {
	defer func() {
		r := recover()  //获取到了panic的传入值
		if err, ok := r.(error); ok {
			fmt.Println("出错了：", err)
		} else {
			panic(fmt.Sprintln("我不知道该干嘛：", r))
		}
	}()

	panic("随便一句话")
	//panic(errors.New("出错panic"))
}
