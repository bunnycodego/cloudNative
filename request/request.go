package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常: ", ins.Error())
		}
	}()
	if err != nil {
		panic(err) //recover捕获panic的参数
	}
}

func testHttpNewRequest() {
	//1.创建一个客户端
	client := http.Client{}
	//2.创建一个请求
	request, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	CheckErr(err)
	//3.客户端发送请求
	cookieName := &http.Cookie{Name: "username", Value: "Marnesh"}
	//添加cookie
	request.AddCookie(cookieName)
	response, err := client.Do(request)
	CheckErr(err)
	//设置请求头
	request.Header.Set("Accept-Language", "zh-cn")
	defer response.Body.Close()
	//查看请求头的数据
	fmt.Printf("Header:%+v\n", request.Header)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	//4.操作数据
	if response.StatusCode == 200 {
		data, err := ioutil.ReadAll(response.Body)
		CheckErr(err)
		fmt.Println("网络请求成功")
		fmt.Println(string(data))
	} else {
		fmt.Println("网络请求失败", response.Status)
	}
}

func main() {
	testHttpNewRequest()
}
