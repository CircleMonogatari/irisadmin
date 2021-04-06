package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreaterMapToParameter(m map[string]interface{}) string  {
	var str string
	for k,v := range m{
		str = str + "&" + fmt.Sprintf("%v=%v", k, v)
	}
	return str[1:]
}

func Get(url string, parameter map[string]interface{}) (string, error) {
	var path string

	if (parameter == nil) || (len(parameter) == 0){
		path = url
	}else {
		path = url + "?" + CreaterMapToParameter(parameter)
	}

	resp, err := http.Get(path)
	if err != nil {
		// handle error
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}