package main

import (
	"fmt"
	"os"
	"path/filepath"
// 	"regexp"
)


func Validate(dir string) {
	abs, err := filepath.Abs(dir)
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
	foo := filepath.Base(abs)
	fmt.Println(foo)

}


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "DIRECTORY...")
		os.Exit(1)
	}

	for i :=1; i < len(os.Args); i++ {
		dir :=  os.Args[i]
		fmt.Println("directory:", dir)
		if _, err := os.Stat(dir); err == nil {
			Validate(dir)
		} else if os.IsNotExist(err) {
			fmt.Println("Directory", dir, "does not exist.")
		} else {
			fmt.Println(err)
		}
	}

}


