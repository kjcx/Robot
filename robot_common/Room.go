package robot_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type RespRoom struct {
	Code int32
	Token string
	Message string
}
type BetInfo struct {
	Money int
	Token string
	Zone string
	C Clients
}
var BetSumChan  = make(chan *BetInfo)
var BetSum = 0
//进入房间
func Index(client_id,token,room_id string) *RespRoom{
	//urlindex := "http://ss.wmy2.com/Api/Index/config/client_id/"+client_id+"/token/"+token
	//fmt.Println(urlindex)
	urlindex := "http://ss.wmy2.com/Api/Room/index/client_id/"+client_id+"/token/"+token+"/room_id/" + room_id
	resp,err := http.Get(urlindex)
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := RespRoom{}
	json.Unmarshal(body,&r)
	return &r
}
//下注
func OnBet(money string,token,zone,client_id string) *RespRoom{
	//统计下注金额
	url := "http://ss.wmy2.com/Api/Room/onBet/bet_balance/" + money + "/token/" + token +"/zone/" + zone + "/client_id/" + client_id + "/level/2"
	fmt.Println(url)
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := RespRoom{}
	json.Unmarshal(body,&r)
	return &r
}
