package robot_common

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)
var ws *websocket.Conn
var Client_id string
type RespWs struct {
	Code int32
	Message string
	Data map[string]interface{}
}
//ws连接
func Ws(orgin string){
	//发起连接
	url := "ws://ss.wmy2.com:8282"
	fmt.Println(url)
	ws,err := websocket.Dial(url, "", "/" + orgin)
	if err != nil {
		fmt.Println(err)
	}
	go HeartbeatSend(ws)//发送心跳
	go ForRead(ws)


}
//心跳每30秒发一次
func HeartbeatSend(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second*10)
		//time.Sleep(time.Microsecond * 1000000)
		websocket.Message.Send(conn, "ping")
	}
}
type Clients struct {
	Origin string
	ClienId string
}
var Client = make(chan *Clients)

//var ChanClients chan *Clients
func ForRead(ws *websocket.Conn) {
	for {
		var msg1 = make([]byte,2048)
		var n int
		var err error
		if n, err = ws.Read(msg1); err != nil {
			fmt.Println("接收信息出错",nil,err,msg1)
			if ws.IsClientConn() {
				//log.Fatal("连接已断开",err.Error())
				fmt.Println(ws.IsClientConn(),"WS")
			}
		}
		//fmt.Printf("Received:", n ,msg1)
		log.Printf("Received:", n ,string(msg1[:n]))
		r := RespWs{}
		json.Unmarshal(msg1[:n],&r)
		fmt.Println(r,r.Data["client_id"])
		if r.Code == 9002 {
			fmt.Println("9002",r.Message)
			Client_id := r.Data["client_id"].(string)
			fmt.Println(ws.Config().Origin,Client_id)
			c :=&Clients{}
			c.Origin =  ws.Config().Origin.String()[1:]
			c.ClienId = Client_id
			Client <- c
		}else if r.Code == 9003 {
			fmt.Println("9005",r.Message)
		} else if r.Code == 9005 {
			start_countdown := r.Data["start_countdown"]
			lottery_id := r.Data["lottery_id"]
			fmt.Println("9005",lottery_id,start_countdown,r.Message)
		}else if r.Code == 9009 {
			//开奖更新下载金额限制
			lottery_id := r.Data["lottery_id"]
			if lottery_id == 1 {
				BetSum = 0
			}
			fmt.Println("9009",r.Message)
		}

	}
}