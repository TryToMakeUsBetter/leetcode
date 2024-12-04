package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrCommitNoneNumber = errors.New("没有提交数据")
	ErrCommitPrime      = errors.New("提交的是质数")
	ErrCommitWrongType  = errors.New("提交了错误的数据类型")
)

var (
	PrimeSuccess = "数据处理成功"
)

var (
	GlobalSmallPrime = make(map[int]int, 0)
)

func Prime(c *gin.Context) {
	input := c.Query("input")

	if input == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ErrCommitNoneNumber.Error(),
		})
		return
	}

	inputNum, err := strconv.Atoi(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ErrCommitWrongType.Error(),
		})
		return
	}

	out, err := PrimeFactorization(inputNum, isPrime)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  err.Error(),
		})
		return
	}
	outStr := transferToStr(inputNum, out)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  PrimeSuccess,
		"out":  outStr,
	})
}

func Display(input int, res []int) {
	fmt.Printf(" %d = ", input)
	for i := 0; i < len(res); i++ {
		fmt.Printf(" %d ", res[i])
		if i != len(res)-1 {
			fmt.Printf("*")
		}
	}
	fmt.Println()
}

func transferToStr(input int, res []int) (str string) {
	s := fmt.Sprintf("%d=", input)
	for i := 0; i < len(res); i++ {
		s = fmt.Sprintf("%s%d", s, res[i])
		if i != len(res)-1 {
			s = fmt.Sprintf("%s*", s)
		}
	}
	return s
}

func PrimeFactorization(input int, isp isPrimeFunc) (out []int, err error) {
	output := make([]int, 0)
	res, flag := isp(input)

	if !flag {
		output = append(output, res)
	} else {
		return []int{}, ErrCommitPrime
	}

	for !flag {
		input = input / res
		res, flag = isp(input)
		output = append(output, res)
	}

	return output, nil

}

type isPrimeFunc func(int) (int, bool)

func isPrime(input int) (res int, flag bool) {
	if input <= 3 {
		return input, true
	}

	if input%2 == 0 {
		return 2, false
	}

	for i := 3; i <= int(math.Sqrt(float64(input)))+1; i += 2 {
		if input%i == 0 {
			return i, false
		}
	}

	return input, true
}

func main() {
	r := gin.New()

	gin.SetMode("debug")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/Prime", Prime)
	}

	s := &http.Server{
		Addr:    ":80",
		Handler: r,
	}

	s.ListenAndServe()
}
