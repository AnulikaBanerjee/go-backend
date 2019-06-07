package util

import (
	"regexp"
	"strconv"
)

func GetEvaluatedString(g string) string {
	var finalRes int
	var finalString string
	reg := regexp.MustCompile(`[^0-9\+\-\*\/]+`)
	processedString := reg.ReplaceAllString(g, "")
	// Now define an operation
	a := regexp.MustCompile(`\+`)
	arr := a.Split(processedString, 2)
	if len(arr) > 1 {
		finalRes1, _ := strconv.Atoi(arr[0])
		finalRes2, _ := strconv.Atoi(arr[1])
		finalRes = finalRes1 + finalRes2
		tempStr := strconv.Itoa(finalRes)
		finalString = g + "= " + tempStr
		return finalString
	}
	a = regexp.MustCompile(`-`)
	arr = a.Split(processedString, 2)
	if len(arr) > 1 {
		finalRes1, _ := strconv.Atoi(arr[0])
		finalRes2, _ := strconv.Atoi(arr[1])
		finalRes = finalRes1 - finalRes2
		tempStr := strconv.Itoa(finalRes)
		finalString = g + "= " + tempStr
		return finalString
	}
	a = regexp.MustCompile(`\*`)
	arr = a.Split(processedString, 2)
	if len(arr) > 1 {
		finalRes1, _ := strconv.Atoi(arr[0])
		finalRes2, _ := strconv.Atoi(arr[1])
		finalRes = finalRes1 * finalRes2
		tempStr := strconv.Itoa(finalRes)
		finalString = g + "= " + tempStr
		return finalString
	}
	a = regexp.MustCompile(`\/`)
	arr = a.Split(processedString, 2)
	if len(arr) > 1 {
		finalRes1, _ := strconv.Atoi(arr[0])
		finalRes2, _ := strconv.Atoi(arr[1])
		finalRes = finalRes1 / finalRes2
		tempStr := strconv.Itoa(finalRes)
		finalString = g + "= " + tempStr
		return finalString
	}

	return "Oops no operator found in" + g

}
