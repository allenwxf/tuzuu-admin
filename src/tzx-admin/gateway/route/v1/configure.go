package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	//"encoding/json"
	//"context"
	"tzx-admin/gateway/global"
	//"tzx-admin/lib/log"
)

func AddNewTourRoute(c *gin.Context) {
	_, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		global.Echo(c, 10001, "body不能为空")
		return
	}
	//route := &msg.AddNewTourRouteRequest{}
	//err = json.Unmarshal(body, &route)
	//if err != nil {
	//	global.Echo(c, 10002, "body格式不对")
	//	return
	//}
	//conn, err := global.GetSingle("route")
	//if err != nil {
	//	global.Echo(c, 10003, "不存在此服务")
	//	return
	//}
	//client := msg.NewRouteControllerClient(conn)
	//resp, err := client.AddNewTourRoute(context.Background(), route)
	//if err != nil {
	//	log.Error(err.Error())
	//	global.Echo(c, 10004, err.Error())
	//	return
	//}
	//global.Echo(c, 10000, resp)
	return
}