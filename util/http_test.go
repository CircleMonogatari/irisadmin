package util

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://www.baidu.com"

	data,err := Get(url, nil)
	if err != nil{
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(data)
}

func Test_CreaterMapToParameter(t *testing.T) {
	data := map[string]interface{}{
		"a" : 1,
		"b":"dasdsa",
		"c":0.5,
	}

	str := CreaterMapToParameter(data)

	if str != "a=1&b=dasdsa&c=0.5"{
		t.Fatal("CreaterMapToParameter Error: ", str)
	}
}
