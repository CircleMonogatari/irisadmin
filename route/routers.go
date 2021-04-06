package route

import (
	"irisadmin/route/api"
	"irisadmin/route/wx"
)

func Init()  {
	wx.Init()
	api.Init()
}
