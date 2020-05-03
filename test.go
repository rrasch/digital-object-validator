package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


// // func FindMissing(regexStr string, files []string, seqLen int) []string {
// 
// func main() {
// 
// 	files := []
// 
// 	regexStr := fmt.Sprintf(`^(%s_)(\d{6})((?:_\d{2}){0,2})_[md].tif$`, id)
// 
// 	reNumbered := Regex.MustCopmile(regexStr)
// 
// // 	numbered = make(map[string]int)
// // 
// // 	var missing []string
// // 
// // 	var theRest []string
// // 
// // 	var maxIndex int
// 
// 	var matches []string
// 
// 	for _, file := range files {
// 		matches = reNumbered.FindStringSubmatch(file.Name())
// 		if matches != nil {
// 			fmt.Prinln(matches)
// // 			i, err := strconv.Atoi(matches[2])
// // 			if err != nil {
// // 				panic(err)
// // 			}
// // 			if i > max {
// // 				maxIndex = i
// // 			}
// // 			numbered[file.Name()] = i;
// 		}
// // 		else {
// // 			append(theRest, file.Name())
// // 		}
// 	}
// 
// // 	// Found no numbered files
// // 	if numbered == nil {
// // 		return NIL
// // 	}
// // 
// // 	format = "%s%0" + seqLen + "d_%s.tif"
// // 
// // 	imageTypes = ["m", "d"]
// // 
// // 	for i := 1; i < maxIndex; i++ {
// // 		for _, imgType := range imageTypes {
// // 			expected := fmt.Sprintf(format, matches[1], imgType)
// // 			if _, found := numbered[expected]; !found {
// // 				append(missing, expected)
// // 				fmt.Println("Missing", expected)
// // 			}
// // 		}
// // 	}
// // 
// // 	if missing != nill {
// // 		fmt.Println("Too many missing files.")
// // 	}
// // 
// // 	return theRest
// }
// 



func main () {
	str := "000006"
	i, err := strconv.Atoi(str)
	fmt.Println(i)
	if err != nil {
		panic(err)
	}

	regex := fmt.Sprintf(`^(%s_)(\d{6})((?:_\d{2}){0,2})_[md].tif$`, "nyu")
// 	regex := fmt.Sprintf(`^(%s_)(\d{6})((_\d{2}){0,2})_[md].tif$`, "nyu")
	fmt.Println(regex)
	re := regexp.MustCompile(regex)
// 	matches := re.FindStringSubmatch("nyu_000012_01_02_m.tif")
// 	matches := re.FindStringSubmatch("nyu_000012_01_02_m.tif")
// 	matches := re.FindAllStringSubmatch("nyu_000012_01_02_m.tif", -1)

	files := [2]string{"nyu_000012_01_02_m.tif", "nyu_000012_01_04_m.tif"}

	mymap := make(map[int][2]int)

	maxPageNum := 0

	// Find maximum page and insert values
	for _, file := range files {

		matches := re.FindStringSubmatch(file)

		if matches != nil {
			fmt.Println(matches)

			pageNum, _ := strconv.Atoi(matches[2])
			fmt.Println("page number:", pageNum)

			if pageNum > maxPageNum {
				maxPageNum = pageNum
			}

			foo, exists := mymap[pageNum]
			if !exists {
				foo = [2]int{0,0}
			}

			fmt.Println(mymap[pageNum][0])

			if matches[3] != "" {
				fmt.Println("matches[3]:", matches[3])

				for i, val := range strings.Split(matches[3], "_") {
					if i == 0 {
						continue
					}
					lastIndex, _ := strconv.Atoi(val)
					fmt.Println(lastIndex)
					if lastIndex >  foo[i - 1] {
						foo[i - 1] = lastIndex
					}
				}
			}

			mymap[pageNum] = foo
		}

		fmt.Println("mymap:", mymap)

	}


	format1 := "%s_%06d_%s.tif"
	format2 := "%s_%06d_%02d_%s.tif"
	format3 := "%s_%06d_%02d_%02d_%s.tif"

// 	imageTypes = ["m", "d"]

	fmt.Println("max page num:", maxPageNum)

	var expected []string

	id := "nyu"

	for i := 1; i <= maxPageNum; i++ {
		foo, exists := mymap[i]
		if !exists {
			expected = append(expected, fmt.Sprintf(format1, id, i, "d"))
		} else {
			for j := 1; j <= foo[0]; j++ {
				if foo[1] == 0 {
					expected = append(expected, fmt.Sprintf(format2, id, i, j, "d"))
				} else {
					for k:= 1; k <= foo[1]; k++ {
						expected = append(expected, fmt.Sprintf(format3, id, i, j, k, "d"))
					}
				}
			}
		}
	}

	for _, file := range expected {
		fmt.Println(file)
	}



// 		for _, imgType := range imageTypes {
// 			expected := fmt.Sprintf(format, "nyu", imgType)
// 			if _, found := numbered[expected]; !found {
// 				append(missing, expected)
// 				fmt.Println("Missing", expected)
// 			}
// 		}



}

