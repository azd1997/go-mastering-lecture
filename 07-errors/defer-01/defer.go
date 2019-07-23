package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-mastering-lecture/07-errors/defer-02/fib"
	"os"
)

func main() {
	//tryDefer2()
	//writeFile("07-errors/fib.txt")
	//writeFileV2("07-errors/fib.txt")
	writeFileV3("07-errors/fib.txt")
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

func tryDefer2() {
	for i:=0;i<100;i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)  //写入缓存
	defer writer.Flush()  //写入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFileV2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println("The file exists.")
		//fmt.Println("Error:", err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)  //写入缓存
	defer writer.Flush()  //写入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFileV3(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error.")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)  //写入缓存
	defer writer.Flush()  //写入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

type ErrFileAlreadyExists struct {
	ErrMsg string  //错误信息
	File string  //文件路径及名称
}

func (err *ErrFileAlreadyExists) Error() string {
	return fmt.Sprintf("出错：%s, %s", err.File, err.ErrMsg)
}