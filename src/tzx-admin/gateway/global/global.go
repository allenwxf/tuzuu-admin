package global

import (
	"tzx-admin/lib/config"
	"tzx-admin/lib/log"

	"google.golang.org/grpc"
	"fmt"
)

var (
	G_Port        string
	G_Runmode     string
	G_Server      map[string]string
	G_SoftVersion string
	G_AccountSId  string
	G_BaseUrl     string
	G_AppId       string
	G_Token       string
	g_Server      map[string]*grpc.ClientConn
)

func init() {
	jsonParse := config.NewJsonStruct()
	v := make(map[string]interface{})
	jsonParse.Load("./config/gateway.json", &v)
	fmt.Print(v)
	G_Port = v["port"].(string)
	G_Runmode = v["runmode"].(string)
	G_SoftVersion = v["softVersion"].(string)
	G_AccountSId = v["accountSId"].(string)
	G_Token = v["token"].(string)
	G_BaseUrl = v["baseUrl"].(string)
	G_AppId = v["appId"].(string)
	servers := v["server"].([]interface{})
	G_Server = make(map[string]string)
	g_Server = make(map[string]*grpc.ClientConn)
	for _, v := range servers {
		//G_Server = append(G_Server, )
		value := v.(map[string]interface{})
		ip := value["ip"].(string)
		G_Server[value["name"].(string)] = ip
	}
	log.Debug("%v", G_Server)
	funcInitLog()
}

//初始化日志
func funcInitLog() {
	// err := log.InitLog(g_Logaddr, g_Name)
	// if err != nil {
	// 	panic("init log err " + err.Error())
	// }
	if G_Runmode == "dev" {
		log.SetLogLevel(log.LEVEL_DEBUG)
		log.SetLogModel(log.MODEL_DEV)
	} else if G_Runmode == "info" {
		log.SetLogLevel(log.LEVEL_INFO)
		log.SetLogModel(log.MODEL_INFO)
	} else {
		log.SetLogLevel(log.LEVEL_WARNING)
		log.SetLogModel(log.MODEL_PRO)
	}
}

//获取微服务
func GetSingle(name string) (*grpc.ClientConn, error) {
	if conn, ok := g_Server[name]; ok {
		//存在
		return conn, nil
	} else {
		conn, err := grpc.Dial(G_Server[name], grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		if g_Server == nil{
			g_Server = make(map[string]*grpc.ClientConn)
		}
		g_Server[name] = conn
		return conn, nil
	}

}

//删除微服务
func DelServer(name string) {
	delete(g_Server, name)
}

func ExitServer() {
	for k, v := range g_Server {
		v.Close()
		log.Debug("exit %s", k)
	}
}
