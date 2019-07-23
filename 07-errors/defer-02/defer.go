package main

//生成一个斐波那契数列的文本，用以在defer-01中使用

import (
	"bufio"
	"fmt"
	"go-mastering-lecture/07-errors/defer-02/fib"
	"io"
	"strings"
)

type intGen func() int

//实现Reader接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for i:=0; i<15 && scanner.Scan(); i++ {  //直到scanner停止扫描且最多打印15行
		fmt.Println(scanner.Text())
	}
}


func main() {
	var f intGen = fib.Fibonacci()
	printFileContent(f)
}

