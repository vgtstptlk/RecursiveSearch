package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

var (
	waitGroup sync.WaitGroup
	chann = make(chan struct{}, runtime.GOMAXPROCS(runtime.NumCPU()))
)


func main() {
	waitGroup.Add(1)
	walk("/Users/vgtstptlk/Desktop")
	waitGroup.Wait()
}


func walk(dir string) {

	chann <- struct{}{}
	defer func() {
		<-chann
		waitGroup.Done()
	}()

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("error... (open)")
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("error... (reading)")
	}

	for _, f := range files {
		if f.IsDir() {
			waitGroup.Add(1)
			go walk(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}

	defer waitGroup.Done()
}
