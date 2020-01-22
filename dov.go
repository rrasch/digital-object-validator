package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)


func IsPreservationWorkflow(files []string) bool {
	reMaster := regexp.MustCompile(`_m\.tif$`)
	reDMaker := regexp.MustCompile(`_d\.tif$`)
	numMaster := 0
	numDMaker := 0
	for i :=1; i < len(files); i++ {
		if reMaster.MatchString(files[i]) {
			numMaster++
		} else if reDMaker.MatchString(files[i]) {
			numDMaker++
		}
	}
	return numMaster > 0 && numDMaker > 0
}


func GetWorkflow(files []string) string {
	workflow := ""

	reTif      := regexp.MustCompile(`\.tif$`)
	reMaster   := regexp.MustCompile(`_m\.tif$`)
	reDMaker   := regexp.MustCompile(`_d\.tif$`)
	reNumbered := regexp.MustCompile(`[0-9]{6}\.tif$`)
	reEOC      := regexp.MustCompile(`(?i)eoc\.csv$`)
	numTif      := 0
	numMaster   := 0
	numDMaker   := 0
	numNumbered := 0
	numEOC      := 0
	for i :=1; i < len(files); i++ {
		if reTif.MatchString(files[i]) {
			numTif++
			if reMaster.MatchString(files[i]) {
				numMaster++
			} else if reDMaker.MatchString(files[i]) {
				numDMaker++
			} else if reNumbered.MatchString(files[i]) {
				numNumbered++
			}
		} else if reEOC.MatchString(files[i]) {
			numEOC++
		}
	}
	if numTif > 0 && numTif == numNumbered {
		if numEOC == 1 {
			workflow = "bookeye"
		} else {
			workflow = "access"
		}
	} else if numMaster > 0 && numDMaker > 0 {
		workflow = "preservation"
	}
	return workflow
}


func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	sort.Strings(files)
	return files, nil
}


func Validate(dir string) {
	abs, err := filepath.Abs(dir)
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
	foo := filepath.Base(abs)
	fmt.Println(foo)

	files, err := OSReadDir(abs)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	fmt.Println("Workflow:", GetWorkflow(files))
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


