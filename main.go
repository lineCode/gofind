package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var regex string

var scanDir string

func init() {
	flag.StringVar(&regex, "f", "", "Find Regex in Filename")
	flag.Parse()
}

func main() {
	counter := 0
	found := 0
	t := time.Now()
	re := regexp.MustCompile(regex)
	setScanDir()
	filepath.Walk(
		scanDir,
		func(fpath string, info os.FileInfo, err error) error {
			counter++
			if re.MatchString(fpath) {
				fmt.Println(fpath)
				found++
			}

			return err
		})
	t1 := time.Now()
	fmt.Println("Files gezählt:    ", counter)
	fmt.Println("Files gefunden:   ", found)
	fmt.Println("Time:             ", t1.Sub(t))
}

func setScanDir() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	scanDir = dir
}
