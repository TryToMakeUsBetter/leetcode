package interview

import "fmt"

/*
字符串操作(把输入中数字后面的字符串重复数字对应的遍数，最开始没有数字默认为1)

input: a12b3cd
output: abbbbbbbbbbbbcdcdcd
*/

func PrintString(input []byte) {
	charList := make([]byte, len(input))
	charRear := -1
	numList := make([]byte, len(input))
	numRear := -1
	var GetNumFlag bool
	var GetStrFlag bool

	if (input[0] >= 'a' && input[0] <= 'z') || (input[0] >= 'A' && input[0] <= 'Z') {
		GetNumFlag = true
		GetStrFlag = false
		numRear++
		numList[numRear] = '1'
	}

	for i := 0; i < len(input); i++ {
		if judgeByteInAlpha(input[i]) {
			GetNumFlag = true
			charRear++
			charList[charRear] = input[i]
			GetStrFlag = false

		} else {
			GetStrFlag = true
			if GetNumFlag && GetStrFlag {
				num, newRear := calPrintNum(numList, numRear)
				numRear = newRear
				charRear = printChar(charList, charRear, num)
			}
			numRear++
			numList[numRear] = input[i]
			GetNumFlag = false
		}

		if i == len(input)-1 && judgeByteInAlpha(input[len(input)-1]) {
			GetStrFlag = true
			if GetNumFlag && GetStrFlag {
				num, newRear := calPrintNum(numList, numRear)
				numRear = newRear
				charRear = printChar(charList, charRear, num)
			}
			numRear++
			numList[numRear] = input[i]
			GetNumFlag = false
		}
	}

}
func judgeByteInAlpha(input byte) bool {
	if (input >= 'a' && input <= 'z') || (input >= 'A' && input <= 'Z') {
		return true
	}
	return false
}
func calPrintNum(numList []byte, rear int) (int, int) {
	var res = 0
	for i := 0; i <= rear; i++ {
		res = int(numList[i]) - int('0') + res*10
	}
	return res, -1
}

func printChar(charList []byte, rear int, num int) int {
	printStr := string(charList[0 : rear+1])
	for i := 0; i < num; i++ {
		fmt.Print(printStr)
	}
	charList = make([]byte, 0)
	return -1
}
