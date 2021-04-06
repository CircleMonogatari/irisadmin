package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMap(t *testing.T){
	m := map[string]interface{}{
		"a":1,
		"b":"qwe",
		"c":0.321,
	}

	var str string
	for k,v := range m{
		str = str + "&" + fmt.Sprintf("%v:%v", k, v)
	}


	fmt.Println(str[1:])
}

func TestWx(t *testing.T){
	//code2Session

	url := "https://api.weixin.qq.com/sns/jscode2session"
	// appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 一次性读取
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("错误")
		return
	}
	fmt.Println(string(bs))
}
