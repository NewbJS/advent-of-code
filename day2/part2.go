package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("c:\\Users\\solar\\downloads\\passwords.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	doThing(strings.Split(string(data), "\n"))

}

func doThing(slice []string) {
	var passwordsFound int32

	for sIdx, el := range slice {
		if sIdx < len(slice)-1 {
			splitAtDash := strings.Split(el, "-")
			f := strings.Split(splitAtDash[1], " ")
			lowestPoint, _ := strconv.ParseInt(splitAtDash[0], 10, 64)
			highestPoint, _ := strconv.ParseInt(f[0], 10, 64)

			z := strings.Split(f[1], "")
			letterToSearchFor := z[0]

			x := strings.Split(f[2], "")
			amountFound := 0

			for idx, n := range x {
				idx64 := int64(idx)
				if n == letterToSearchFor && (idx64+1 == lowestPoint || idx64+1 == highestPoint) {
					amountFound++

				}
			}

			if amountFound == 1 {
				passwordsFound++
			}

		}
	}
	fmt.Println("Found", passwordsFound, "valid passwords!")
}
