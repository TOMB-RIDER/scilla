/*

=======================
Scilla - Information Gathering Tool
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/scilla

	@Author:      edoardottt, https://www.edoardoottavianelli.it

*/

package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//CheckIgnore checks the inputted status code to be ignored
func CheckIgnore(input string) []string {
	result := []string{}
	temp := strings.Split(input, ",")
	temp = RemoveDuplicateValues(temp)
	for _, elem := range temp {
		elem := strings.TrimSpace(elem)
		if len(elem) != 3 {
			fmt.Println("The status code you entered is invalid (It should consist of three characters).")
			os.Exit(1)
		}
		if ignoreInt, err := strconv.Atoi(elem); err == nil {
			// if it is a valid status code without * (e.g. 404)
			if 100 <= ignoreInt && ignoreInt <= 599 {
				result = append(result, elem)
			} else {
				fmt.Println("The status code you entered is invalid (100 <= code <= 599).")
				os.Exit(1)
			}
		} else if strings.Contains(elem, "*") {
			// if it is a valid status code without * (e.g. 4**)
			if IgnoreClassOk(elem) {
				result = append(result, elem)
			} else {
				fmt.Println("The status code you entered is invalid. You can enter * only as 1**,2**,3**,4**,5**.")
				os.Exit(1)
			}
		}
	}
	result = RemoveDuplicateValues(result)
	result = DeleteUnusefulIgnoreresponses(result)
	return result
}

//DeleteUnusefulIgnoreresponses removes from to-be-ignored arrays
//the responses included yet with * as classes
func DeleteUnusefulIgnoreresponses(input []string) []string {
	var result []string
	toberemoved := []string{}
	classes := []string{}
	for _, elem := range input {
		if strings.Contains(elem, "*") {
			classes = append(classes, elem)
		}
	}
	for _, class := range classes {
		for _, elem := range input {
			if class[0] == elem[0] && elem[1] != '*' {
				toberemoved = append(toberemoved, elem)
			}
		}
	}
	result = Difference(input, toberemoved)
	return result
}

//IgnoreClassOk states if the class of ignored status codes
//is correct or not (4**,2**...)
func IgnoreClassOk(input string) bool {
	if strings.Contains(input, "*") {
		if _, err := strconv.Atoi(string(input[0])); err == nil {
			i, err := strconv.Atoi(string(input[0]))
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			if i >= 1 && i <= 5 {
				if input[1] == byte('*') && input[2] == byte('*') {
					return true
				}
			}
		}
	}
	return false
}

//IgnoreResponse returns a boolean if the response
//should be ignored or not.
func IgnoreResponse(response int, ignore []string) bool {
	responseString := strconv.Itoa(response)
	// if I don't have to ignore responses, just return true
	if len(ignore) == 0 {
		return false
	}
	for _, ignorePort := range ignore {
		if strings.Contains(ignorePort, "*") {
			if responseString[0] == ignorePort[0] {
				return true
			}
		}
		if responseString == ignorePort {
			return true
		}
	}
	return false
}
