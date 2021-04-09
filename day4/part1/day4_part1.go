package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const VALID = "byr,iyr,eyr,hgt,hcl,ecl,pid"

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := 0
	for _, s := range strings.Split(string(f), "\n\n") {
		if checkSubstrings(s) {
			r++
		}
	}
	fmt.Println(r)
}

func checkSubstrings(s string) bool {
	for _, n := range strings.Split(VALID, ",") {
		if !strings.Contains(s, n) {
			return false
		}
	}
	return true
}
