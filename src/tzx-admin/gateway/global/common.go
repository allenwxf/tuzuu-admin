package global

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

type Rules struct {
	key   string
	Value string
}

//获取body
func Unmarshal(c *gin.Context) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	msg := make(map[string]interface{})
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

//获取客户的Ip
func GetIp(c *gin.Context) string {
	req := c.Request
	ip := req.Header.Get("Remote_addr")
	if ip == "" {
		ip = req.RemoteAddr
	}
	return ip
}

//验证参数是否存在
func ValidateData(v_data map[string]interface{}, rules []Rules) (bool, string) {
	for _, value := range rules {
		param := value.key
		rule := value.Value
		if _, ok := v_data[param]; !ok {
			return false, rule
		}
		if v_data[param] == nil {
			return false, rule
		}
	}
	return true, ""
}

func Echo(c *gin.Context, code int, body interface{}) {
	defer RecoverHandle("Echo ...")
	c.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Body": body,
		"Time": time.Now().Unix(),
	})
}


func RecoverHandle(v ...interface{}) {
	if err := recover(); err != nil {
		if len(v) > 0 {
			s := fmt.Sprintln(v...)
			fmt.Println(err, "\n", s, "\n", string(debug.Stack()))
		} else {
			fmt.Println(err, "\n", "\n", string(debug.Stack()))
		}
	}
}