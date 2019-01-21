package robot_common

import (
	"encoding/json"
	"fmt"
	"goadmin/common"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"strconv"
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
	//url := "ws://127.0.0.1:12345/ws"
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
		if conn != nil {
			websocket.Message.Send(conn, "ping")
			//fmt.Println(conn.Config().Origin,"发送心跳")
		}

	}
}
type Clients struct {
	Origin  string
	ClienId string
	Ws      *websocket.Conn
}
var Client = make(chan *Clients)
var ClientOnline []*Clients
var Game1OnBetStatus = 0
//var ChanClients chan *Clients
func ForRead(ws *websocket.Conn) {
	for {
		var msg1 = make([]byte,2048)
		//var n int
		if ws != nil {
			n, err := ws.Read(msg1);
			if  err != nil {
				fmt.Println("接收信息出错",nil,err,msg1)
				if ws.IsClientConn() {
					break
					log.Fatal("连接已断开",err.Error())
					fmt.Println(ws.IsClientConn(),"WS",n,ws.Config().Origin,io.EOF)
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
				c.Ws = ws
				//保存所有ws连接
				ClientOnline = append(ClientOnline,c)
				Client <- c
			}else if r.Code == 9003 {
				fmt.Println("9003",r.Message)
			} else if r.Code == 9005 {
				lottery_id := r.Data["lottery_id"]
				//fmt.Println(lottery_id.(string) == "1")
				start_countdown := r.Data["start_countdown"]
				fmt.Printf("%T\n",start_countdown)
				if lottery_id == "1" {
					go func() {
						fmt.Println("aaaaaaaaaaaaaaaaaa")
						string := strconv.FormatFloat(start_countdown.(float64), 'f', -1, 64)
						start := common.StringToInt(string)
						time.Sleep(time.Duration(start)*time.Second)
						Game1OnBetStatus = 1
						time.Sleep(160*time.Second)
						Game1OnBetStatus = 0
					}()
				}
				fmt.Println("9005",lottery_id,start_countdown,r.Message)
			}else if r.Code == 9009 {
				//开奖更新下载金额限制
				lottery_id := r.Data["lottery_id"]
				//fmt.Println(lottery_id.(string) == "1")
				if lottery_id == 1.0 {
					fmt.Println("1111")
					BetSum = 0
				}
				fmt.Println("9009",r.Message)
			}
		}
	}
}