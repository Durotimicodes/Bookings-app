package main

import (
	"fmt"
	"strings"
)

func main() {

	str := []string{"ower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {

	var prefix string
	var result string

	first2str := strs[0]
	val := strings.Split(first2str, "")
	prefix = val[0] + val[1]
	fmt.Println(prefix)

	//loop over the array datastructure
	for i := 0; i < len(strs); i++ {

			/*pick the first string, get the prefix and check if all other strings have similar prefix*/
		
		if strings.HasPrefix(string(strs[i]), prefix) {
			//if similar, store value in "prefix " and return it
			result = prefix
		} else {
			result = ""
		}
	}
	return result

}
