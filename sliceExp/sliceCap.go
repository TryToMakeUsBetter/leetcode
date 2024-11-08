package sliceexp

func MySlice() (res1, res2 []int) {
	slice1 := make([]int, 5, 10)
	slice2 := slice1[0:2]

	for i := 0; i < 5; i++ {
		slice1[i] = i
	}

	return slice1, slice2
}

/*
Enlarge Slice,but the original slice address doesn't change
*/
func EnlargeSlice1(slice1 *[]int) {
	for i := 0; i < 2; i++ {
		*slice1 = append(*slice1, i+5)
	}
	(*slice1)[0] = -1
}

/*
Enlarge Slice,but the original slice address changes
*/
func EnlargeSliceAgain(slice1 *[]int) {
	for i := 0; i < 4; i++ {
		*slice1 = append(*slice1, i+len(*slice1)-1)
	}
	(*slice1)[0] = -100
}
