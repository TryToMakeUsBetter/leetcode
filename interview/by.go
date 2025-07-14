package interview

func ReverseString(input string) (res string) {
	var l = -1
	for index := 0; index < len(input)-1; index++ {
		lb := 0
		rb := len(input) - 1
		sl, tmpRes, slb, srb := single(input, index)
		if sl > l {
			l = sl
			res = tmpRes
			lb = slb
			rb = srb
		}
		if input[index+1] == input[index] {
			ml, tmpRes, mlb, mrb := multi(input, index, index+1)
			if ml > l {
				l = ml
				res = tmpRes
				lb = mlb
				rb = mrb
			}
		}

		if (lb == 0 && (lb) > len(input)-1-rb) || (rb == len(input)-1 && (lb) < len(input)-1-rb) {
			break
		}
	}
	return res
}

func single(input string, startIndex int) (length int, res string, leftBoundary, rightBoundary int) {
	leftIndex := startIndex
	rightIndex := startIndex
	for leftIndex >= 0 && rightIndex < len(input) && input[leftIndex] == input[rightIndex] {
		leftIndex--
		rightIndex++
	}

	if leftIndex >= 0 && rightIndex < len(input) {
		if input[leftIndex] != input[rightIndex] {
			leftIndex++
			rightIndex--
		}
	} else {
		leftIndex++
		rightIndex--
	}
	return rightIndex - leftIndex + 1, input[leftIndex : rightIndex+1], leftIndex, rightIndex
}

func multi(input string, leftStartIndex, rightStartIndex int) (length int, res string, leftBoundary, rightBoundary int) {
	for leftStartIndex >= 0 && rightStartIndex < len(input) && input[leftStartIndex] == input[rightStartIndex] {
		leftStartIndex--
		rightStartIndex++
	}
	if leftStartIndex >= 0 && rightStartIndex < len(input) {
		if input[leftStartIndex] != input[rightStartIndex] {
			leftStartIndex++
			rightStartIndex--
		}
	} else {
		leftStartIndex++
		rightStartIndex--
	}
	return leftStartIndex - rightStartIndex + 1, input[leftStartIndex : rightStartIndex+1], leftStartIndex, rightStartIndex
}
