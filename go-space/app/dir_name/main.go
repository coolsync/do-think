package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// open dir
	dir_path := "/home/dart/Documents/doc/run2/"
	// files, err := os.OpenFile(dir_path, os.O_RDWR, os.ModeDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer files.Close()

	dirs, err := os.ReadDir(dir_path)
	if err != nil {
		log.Fatal(err)
	}

	// traveser dirs, get dir
	for _, dir := range dirs {
		fileinfo, _ := dir.Info()
		if !fileinfo.IsDir() {
			continue
		}
		if strings.HasPrefix(fileinfo.Name(), "learn") {
			tmp := strings.Split(fileinfo.Name(), "_")[1]
			tmp = "open" + "_" + tmp
			// fmt.Println(tmp)
			os.Rename(dir_path+dir.Name(), dir_path+tmp)
		}
	}
}
