package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main () {
	str := "000006"
	i, err := strconv.Atoi(str)
	fmt.Println(i)
	if err != nil {
		panic(err)
	}

	regex := fmt.Sprintf(`^(%s_)(\d{6})((?:_\d{2}){0,2})_[md].tif$`, "nyu")
	fmt.Println(regex)
	re := regexp.MustCompile(regex)
// 	matches := re.FindStringSubmatch("nyu_000012_01_02_m.tif")
	matches := re.FindStringSubmatch("nyu_000012_01_02_m.tif")
// 	matches := re.FindAllStringSubmatch("nyu_000012_01_02_m.tif", -1)


	mymap := make(map[string][]int)

	if matches != nil {
		fmt.Println(matches)

		var insertNum []int
		if matches[3] != "" {
			fmt.Println(matches[3])

			for i, val := range strings.Split(matches[3], "_") {
				if i == 0 {
					continue
				}
				lastIndex, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				insertNum = append(insertNum, lastIndex)
				fmt.Println(lastIndex)
			}
			mymap["foo"] = insertNum
		}

		fmt.Println(mymap)

	}


}

