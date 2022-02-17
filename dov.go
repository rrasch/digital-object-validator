package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	//"strings"
)


func IsPreservationWorkflow(files []string) bool {
	reMaster := regexp.MustCompile(`_m\.tif$`)
	reDMaker := regexp.MustCompile(`_d\.tif$`)
	numMaster := 0
	numDMaker := 0
	for i := 0; i < len(files); i++ {
		if reMaster.MatchString(files[i]) {
			numMaster++
		} else if reDMaker.MatchString(files[i]) {
			numDMaker++
		}
	}
	return numMaster > 0 && numDMaker > 0
}


// 		expected := fmt.Sprintf("%s_%06d.tif", id, i)
//
// 		if file.Name() != expected {
// 			fmt.Println("Found", file.Name(), "but was expecting", expected)
// 		}

func ValidateAccess(id string, files []string) bool {

	reTif      := regexp.MustCompile(`\.tif$`)
	reNumbered := regexp.MustCompile("^" + id + `_(\d{6})\.tif$`)

	numErrors := 0

	numbered := make(map[string]int)

	var mislabeled []string

	var missing []string

	for _, file := range files {
		if reNumbered.MatchString(file) {
			numbered[file] = 1;
		} else if reTif.MatchString(file) {
			mislabeled = append(mislabeled, file)
			numErrors++
			fmt.Println(file, "found but not named properly.")
		} else {
			fmt.Println("Unknown file found:", file)
		}
	}

	found := reNumbered.FindStringSubmatch(files[len(files)-1])
	lastIndex, err := strconv.Atoi(found[1])
	if err != nil {
		panic(err)
	}
	for i := 1; i < lastIndex; i++ {
		expected := fmt.Sprintf("%s_%06d.tif", id, i)
		if _, found := numbered[expected]; !found {
			missing = append(missing, expected)
			numErrors++
			fmt.Println("Missing", expected)
		}
	}
	if numErrors > 0 {
		return true;
	} else {
		return false;
	}
}


func FindMissingRegex(regexStr string, files []string, seqLen int) []string {
	reNumbered := regexp.MustCompile(regexStr)

	numbered := make(map[string]int)

	var missing []string

	var theRest []string

	var maxIndex int

	var matches []string

	for _, file := range files {
		matches = reNumbered.FindStringSubmatch(file)
		if matches != nil {
			i, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			if i > maxIndex {
				maxIndex = i
			}
			numbered[file] = i;
		} else {
			theRest = append(theRest, file)
		}
	}

	// Found no numbered files
	if numbered == nil {
		return nil
	}

	format := "%s%0" + strconv.Itoa(seqLen) + "d_%s.tif"

	imageTypes := [2]string{"m", "d"}

	for i := 1; i < maxIndex; i++ {
		for _, imgType := range imageTypes {
			expected := fmt.Sprintf(format, matches[1], imgType)
			if _, found := numbered[expected]; !found {
				missing = append(missing, expected)
				fmt.Println("Missing", expected)
			}
		}
	}

	if missing != nil {
		fmt.Println("Too many missing files.")
	}

	return theRest
}





func FindMissing(id string, files []string, seqLen int) []string {


	regexStr := fmt.Sprintf(`^(%s_)(\d{6})((?:_\d{2}){0,2})_[md].tif$`, id)

	reNumbered := regexp.MustCompile(regexStr)

	numbered := make(map[string]int)

	var missing []string

	var theRest []string

	var maxIndex int

	var matches []string

	for _, file := range files {
		matches = reNumbered.FindStringSubmatch(file)
		if matches != nil {
			i, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			if i > maxIndex {
				maxIndex = i
			}
			numbered[file] = i;
		} else {
			theRest = append(theRest, file)
		}
	}

	// Found no numbered files
	if numbered == nil {
		return nil
	}

	format := "%s%0" + strconv.Itoa(seqLen) + "d_%s.tif"

	imageTypes := [2]string{"m", "d"}

	for i := 1; i < maxIndex; i++ {
		for _, imgType := range imageTypes {
			expected := fmt.Sprintf(format, matches[1], imgType)
			if _, found := numbered[expected]; !found {
				missing = append(missing, expected)
				fmt.Println("Missing", expected)
			}
		}
	}

	if missing != nil {
		fmt.Println("Too many missing files.")
	}

	return theRest
}











func ValidateFrontMatter(id string, files []string) {
	regex := fmt.Sprintf(`^(%s_fr)(\d{2,3})_[md].tif$`, id)
	FindMissingRegex(regex, files, 2)
}

func ValidateMiddle(id string, files []string) {
	regex := fmt.Sprintf(`^(%s_)(\d{6})((?:_\d{2}){0,2})_[md].tif$`, id)
	FindMissingRegex(regex, files, 6)
}

func ValidateBackMatter(id string, files []string) {
	regex := fmt.Sprintf(`^(%s_bk)(\d{2,3})_[md].tif$`, id)
	FindMissingRegex(regex, files, 2)
}





func ValidateBookEye(id string, files []string) {

}


func ValidatePreservation(id string, files []string) {

}


func IsValidID (id string) bool {
	reID := regexp.MustCompile(`^[0-9A-Za-z]+(_[0-9A-Za-z]+)*$`)
	return reID.MatchString(id)
}


func GetWorkflow(files []string) string {
	workflow := ""

	reTif      := regexp.MustCompile(`\.tif$`)
	reMaster   := regexp.MustCompile(`_m\.tif$`)
	reDMaker   := regexp.MustCompile(`_d\.tif$`)
	reNumbered := regexp.MustCompile(`([0-9]{6})\.tif$`)
	reEOC      := regexp.MustCompile(`(?i)eoc\.csv$`)
	numTif      := 0
	numMaster   := 0
	numDMaker   := 0
	numNumbered := 0
	numEOC      := 0
	for i := 0; i < len(files); i++ {
		fmt.Println("files", i, files[i])
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
	fmt.Println("tifs:", numTif)
	fmt.Println("masters:", numMaster)
	fmt.Println("dmakers:", numDMaker)
	fmt.Println("numbered tifs:", numNumbered)
	fmt.Println("eoc files:", numEOC)
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
// 	reTif := regexp.MustCompile(`\.tif$`)
	for _, file := range fileInfo {
// 		if reTif.MatchString(file.Name()) {
			files = append(files, file.Name())
// 		}
	}
	sort.Strings(files)
	return files, nil
}


func Validate(dir string) {
	abs, err := filepath.Abs(dir)
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
	id := filepath.Base(abs)
	fmt.Println(id)

	if !IsValidID(id) {
		fmt.Println(id, "is not a valid identifier.")
	}

	files, err := OSReadDir(abs)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	workflow := GetWorkflow(files)
	fmt.Println("Workflow:", workflow)
	log.Print("Workflow:", workflow)

	switch workflow {
	case "access":
		ValidateAccess(id, files)
	case "bookeye":
		ValidateBookEye(id, files)
	case "preservation":
		ValidatePreservation(id, files)
	}

}


func IsTif(file string) {
	out, err := exec.Command("file", file).CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "DIRECTORY...")
		os.Exit(1)
	}

// 	for i := 1; i < len(os.Args); i++ {
	for _, dir := range os.Args[1:]  {
// 		dir :=  os.Args[i]
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


