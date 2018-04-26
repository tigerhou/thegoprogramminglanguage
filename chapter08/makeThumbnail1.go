package main

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
	"tigerhou.com/thegoprogramminglanguage/chapter08/thumbnail"
)

func main() {
	dirname := "/Users/tigerhou/Documents/妈出国旅游照片"
	os.Chdir(dirname)
	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		println(err)
		return
	}
	filename := make(chan string)
	go makeThumbnail2(filename)
	for _, file := range dirs {
		f := file
		if f.IsDir() || f.Name() == ".DS_Store" {
			continue
		}
		filename <- f.Name()
	}
}

func makeThumbnail(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnail2(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumbfile, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
