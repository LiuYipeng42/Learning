package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	// "io/ioutil"
)

// http://httpbin.org/get
func main() {

	// http.Get 在运行前会读取环境变量，会自动使用代理，所以要清除代理对应环境变量
	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("http_proxy")
	os.Unsetenv("HTTPS_PROXY")
	os.Unsetenv("https_proxy")

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		byteNum, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\nbyteNum: %d\n", byteNum)

		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
	}
}
