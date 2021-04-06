package wx

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"irisadmin/env"
	"irisadmin/util"
	"net/http"
)

var (
	APPID = ""
	SECRET = ""
)

func Init(){
	fmt.Println("加载 wx router")
	var app = env.GetApp()

	wxParty := app.Party("/wx")
	wxParty.Get("/code2Session", Code2Session)
}


type Code2SessionJson struct {
	Openid		string `json:"openid"`			/* 用户唯一标识 */
	Session_key		string `json:"session_key"`	/* 会话密钥 */
	Unionid		string `json:"unionid"`				/* 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。*/
	Errcode		int `json:"errcode"`				/*错误码*/
	Errmsg		string `json:"errmsg"`				/*错误信息*/
}



func Code2Session(ctx iris.Context){
	//GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := "https://api.weixin.qq.com/sns/jscode2session"

	JSCODE := ctx.URLParam("JSCODE")

	param := util.CreaterMapToParameter(map[string]interface{}{
		"appid":APPID,
		"secret":SECRET,
		"js_code":JSCODE,
		"grant_type":"authorization_code",
	})

	resp, err := http.Get(url+"?"+param)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var res Code2SessionJson
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	ctx.JSON(res)
}
