package log
import (
	"testing"
)

/*
*测试Log函数
*/
func Test_Log(t *testing.T) {
	err:=InitLog("172.16.0.47:4150","test_log")
	if err!=nil{
		t.Fatal("init log err "+err.Error())
		return
	}
	SetLogLevel(LEVEL_INFO)
	SetLogModel(MODEL_INFO)
	Debug("go test debug")
	Info("go test info ")
	Warn("go test warn ")
	Error("go test error ")
}