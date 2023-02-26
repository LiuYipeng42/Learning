package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"thumbnail"
)

func main() {

	filenameChan := make(chan string, 4)
	filenameChan <- "./test1.jpg"
	filenameChan <- "./test2.jpg"
	filenameChan <- "./test3.jpg"
	close(filenameChan)
	// makeThumbnails(filenameChan)

	filenameSlice := []string{"./test1.jpg", "./test2.jpg", "./test3.jpg"}

	makeThumbnailsSlice(filenameSlice)

}

// 从通道获取文件地址，返回生成的略缩图的字节数
func makeThumbnails(filenames <-chan string) int64 {

	sizes := make(chan int64)
	var wg sync.WaitGroup // 正在工作的 goroutine 的数量

	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done() // 减 1
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		// 等待变量的值变为 0
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		log.Println(size)
		total += size
	}
	return total
}

func makeThumbnailsSlice(filenames []string) (thumbfiles []string) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for i := 0; i < len(filenames); i++ {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(filenames[i])
	}

	for i := 0; i < len(filenames); i++ {
		it := <-ch
		if it.err != nil {
			fmt.Println(fmt.Errorf("filename: %s\nerror: %e", it.thumbfile, it.err))
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return
}
