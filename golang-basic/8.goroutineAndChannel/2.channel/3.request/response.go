package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 缓冲通道
// 发送操作会在队列的尾部插入一个元素，接收操作会从队头移除一个元素
// 当通道被填满时，进行发送操作的 goroutine 会阻塞，直到通道中有空余容量
// 当通道为空时，进行接收操作的 goroutine 会阻塞，直到通道中有数据
func main() {

	// 创建一个容量为 3 的通道
	response := make(chan string, 3)
	fmt.Printf("容量：%d\n", cap(response))
	go func() {response <- request("http://120.77.177.229/")}()
	go func() {response <- request("http://120.77.177.229/")}()
	go func() {response <- request("http://120.77.177.229/")}()
	doc := <-response

	fmt.Printf("通道内元素个数：%d\n", len(response))
	fmt.Printf("len: %d\nfirst 100 char: \n%s" ,len(doc), doc[:100])
}

func request(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println(fmt.Errorf("getting %s: %s", url, resp.Status))
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	resp.Body.Close()
	return string(body)
}