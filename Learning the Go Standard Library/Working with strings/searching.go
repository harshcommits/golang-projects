package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "The quick brown fox jumps over the lazy dog"
	vowels := "aeiou"
	file1 := "filename.txt"
	file2 := "temp_picfile.jpg"

	//simple search functions
	fmt.Println(strings.Contains(s, "jump"))

	//search for specific chars
	fmt.Println(strings.ContainsAny(s, "abcd"))

	//use index to find first occurence of substring
	fmt.Println(strings.Index(s, "fox")) //prints 16, since f in fox is 16th index
	fmt.Println(strings.Index(s, "cat")) //prints -1; meaning doesn't exist

	//IndexAny
	fmt.Println(strings.IndexAny("jdkksr", vowels)) //prints -1, since doesn't exist
	fmt.Println(strings.IndexAny("Golang", vowels)) //prints 1, since o in index 1

	//check for prefix and suffix
	fmt.Println(strings.HasPrefix(file2, "temp"))
	fmt.Println(strings.HasPrefix(file1, "txt"))

	//count substrings
	fmt.Println(strings.Count(s, "the")) //prints 1, since only one lowercase 'the'
	fmt.Println(strings.Count(s, "he"))  //prints 2, since 2 instances for 'he'

}
