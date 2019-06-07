package util

import (
	"regexp"
	"strconv"
)

func getSanitizedString(userInput string) string {
	reg := regexp.MustCompile(`[^0-9\+\-\*\/]+`)
	sanitizedString := reg.ReplaceAllString(userInput, "")
	return sanitizedString
}

func getUserNumbersFromInput(operandString1 string, operandString2 string) (int, int) {
	userInputFirstNumber, _ := strconv.Atoi(operandString1)
	userInputSecondNumber, _ := strconv.Atoi(operandString2)
	return userInputFirstNumber, userInputSecondNumber
}

func getProcessedString(userInput string, finalRes int) string {
	tempStr := strconv.Itoa(finalRes)
	finalString := userInput + "= " + tempStr
	return finalString
}

func GetEvaluatedString(userInput string) string {
	var finalRes int
	var finalString string

	sanitizedString := getSanitizedString(userInput)
	// Now perform an operation based on user input

	//if Operator is addition
	operator := regexp.MustCompile(`\+`)
	arr := operator.Split(sanitizedString, 2)
	if len(arr) > 1 {
		userInputFirstNumber, userInputSecondNumber := getUserNumbersFromInput(arr[0], arr[1])
		finalRes = userInputFirstNumber + userInputSecondNumber
		finalString = getProcessedString(userInput, finalRes)
		return finalString
	}
	//if Operator is subtraction
	operator = regexp.MustCompile(`-`)
	arr = operator.Split(sanitizedString, 2)
	if len(arr) > 1 {
		userInputFirstNumber, userInputSecondNumber := getUserNumbersFromInput(arr[0], arr[1])
		finalRes = userInputFirstNumber - userInputSecondNumber
		finalString = getProcessedString(userInput, finalRes)
		return finalString
	}
	//if Operator is multiplication
	operator = regexp.MustCompile(`\*`)
	arr = operator.Split(sanitizedString, 2)
	if len(arr) > 1 {
		userInputFirstNumber, userInputSecondNumber := getUserNumbersFromInput(arr[0], arr[1])
		finalRes = userInputFirstNumber * userInputSecondNumber
		finalString = getProcessedString(userInput, finalRes)
		return finalString
	}
	//if Operator is division
	operator = regexp.MustCompile(`\/`)
	arr = operator.Split(sanitizedString, 2)
	if len(arr) > 1 {
		userInputFirstNumber, userInputSecondNumber := getUserNumbersFromInput(arr[0], arr[1])
		finalRes = userInputFirstNumber / userInputSecondNumber
		finalString = getProcessedString(userInput, finalRes)
		return finalString
	}

	//else return error message
	return "Oops!!! Is something wrong with the input provided? " + userInput

}
