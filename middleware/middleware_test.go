package middleware

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T)  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)

		}
	}()
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()


	fmt.Println("3")


	panic("测试")

}