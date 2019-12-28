package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go walk(file.Name(), &wg)
	}
	wg.Wait()
}

func walk(root string, wg *sync.WaitGroup) {
	start := time.Now()
	err := filepath.Walk(root, found)
	if err != nil {
		if wd, wderr := os.Getwd(); wderr == nil {
			fmt.Printf("error walking the path %q: %v\n", wd, err)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("%s: %d\n", root, elapsed.Milliseconds())
	wg.Done()
}

func found(dir string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
		return err
	}
	if info.IsDir() && info.Name() == ".git" {
		fmt.Println(path.Dir(dir))
	}
	return nil
}
