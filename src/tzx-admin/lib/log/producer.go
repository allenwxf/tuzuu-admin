package log

import (
	"os"
	"log"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

//初始化
func init(){
	producer=nil
}


/*
*初始化nsq生产者
*参数说明:
*@param:nsqdAddr		nsqd地址
*@param:topic			话题
*/
func InitProducer(nsqdAddr string)(error) {
	cfg := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer(nsqdAddr, cfg)
	producer.SetLogger(log.New(os.Stderr, "", log.Flags()),nsq.LogLevelError)
	if err != nil||producer==nil{
		log.Fatal(err)
		return err
	}
	return nil
}

/*
*生产者nsq
*参数说明:
*@param:nsqdAddr		nsqd地址
*@param:topic			话题
*/
func Publish(topic string,msg []byte)(error) {
	//producer=initProducer(ipip,conf)
	if err := producer.Publish(topic, []byte(msg)); err != nil {
		log.Fatal("publish error: " + err.Error())
		return err
	}
	return nil
}

func Close(){
	producer=nil
}