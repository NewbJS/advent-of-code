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

var passwordsFound int = 0

func doThing(slice []string) {
	for sIdx, el := range slice {
		if sIdx < len(slice)-1 {
			// Sorry for poor varible names btw
			splitAtDash := strings.Split(el, "-")
			f := strings.Split(splitAtDash[1], " ")
			lowestPoint, _ := strconv.ParseInt(splitAtDash[0], 10, 64)
			highestPoint, _ := strconv.ParseInt(f[0], 10, 64)

			z := strings.Split(f[1], "")
			letterToSearchFor := z[0]
			var amountFound int64 = 0

			x := strings.Split(f[2], "")
			for _, n := range x {
				if n == letterToSearchFor {
					amountFound++
				}
			}
			if amountFound >= lowestPoint && amountFound <= highestPoint {
				passwordsFound++
			}
		}
	}
	fmt.Println("Found", passwordsFound, "valid passwords!")
}
