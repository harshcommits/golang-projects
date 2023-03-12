package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	s := "The quick brown fox jumps over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string, with some punctuations. Yes!"

	//split strings
	sub1 := strings.Split(s, " ")
	fmt.Printf("%q\n", sub1) //%q quotes every word; split on spaces in string
	sub2 := strings.Split(s, "the")
	fmt.Printf("%q\n", sub2) //split string in 2 parts from 'the' before lazy

	//join strings
	fmt.Println(strings.Join(s2, " - ")) // joins words using delimiter '-'

	//split around white space
	fmt.Printf("%q\n", strings.Fields(s))

	//FieldsFunc; foll. function divides the string based on punctuations
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	fmt.Printf("%q\n", strings.FieldsFunc(s3, f))

	//replacer to replace data in string
	rep := strings.NewReplacer(".", "|", ",", "|", "!", "|") //replace . , ! with |
	fmt.Println(rep.Replace(s3))

}
