package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

/*自定义错误事件appHandler*/

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

//创建自定义类型，继承自http服务器文件中的UserErrorInterface
type testingUserError string
func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}
func errUserError(w http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}
func errUnknown(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func noError(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "no error")
	return nil
}



var tests = []struct{
	h appHandler
	code int
	msg string
} {
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}


func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapperV4(tt.h)
		//使用httptest包，httptest.ResponseRecorder实现了http.ResponseWriter
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",	//相当于浏览器输入的url，这里是随便写的
			nil,
			)

		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		//因为body会末尾有换行，把它去掉
		body := strings.Trim(string(b), "\n")

		if response.Code != tt.code || body != tt.msg {

			t.Errorf("期望获得 （%d， %s）， 但是却得到：（%d， %s）",
				tt.code, tt.msg, response.Code, body)
		}

	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapperV4(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))		//将f转为Server (interface)
		response, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")

		//response.StatusCode
		if response.StatusCode != tt.code || body != tt.msg {
			t.Errorf("期望获得 （%d， %s）， 但是却得到：（%d， %s）",
				tt.code, tt.msg, response.StatusCode, body)
		}
	}
}