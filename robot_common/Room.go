package robot_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
type RespRoom struct {
	Code int32
	Token string
	Message string
	Data *Data
}
type Data struct {
	CountdownInfo *CountdownInfo
}
type BetInfo struct {
	Money int
	Token string
	Zone string
	C Clients
}
type CountdownInfo struct {
	Status int64
	Countdown int64
	FpCountdown int64
	StartCountdown int64
	LotteryId string
}
var BetSumChan  = make(chan *BetInfo)
var BetSum = 0
var StartCountdownSleep bool
//进入房间
func Index(client_id,token,room_id string) *RespRoom{
	//urlindex := "http://ss.wmy2.com/Api/Index/config/client_id/"+client_id+"/token/"+token
	//fmt.Println(urlindex)
	urlindex := "http://ss.wmy2.com/Api/Room/index/client_id/"+client_id+"/token/"+token+"/room_id/" + room_id
	fmt.Printf(urlindex)
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
	if r.Data.CountdownInfo.Status == 3 && StartCountdownSleep == false{
		StartCountdownSleep = true
		go func() {
			time.Sleep(time.Duration(r.Data.CountdownInfo.StartCountdown)*time.Second)
			Game1OnBetStatus = 1
		}()
	}else if r.Data.CountdownInfo.Status ==1 || r.Data.CountdownInfo.Status == 2{
		Game1OnBetStatus = 1
		StartCountdownSleep = true
	}else{
		StartCountdownSleep = false
	}
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
