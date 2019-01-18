package main

import (
	"Robot/robot_common"
	"ckp/ckp_common"
	"math/rand"
	"net/http"
	"time"
)
import "fmt"
var Bet = []string{"10","50","100","50","500","10","1000","10","100","50","1000","10","100","50","10"}
func main(){
	//1读取账号
	accounts := robot_common.File("./robot_common/account.txt")
	fmt.Println(accounts)
	//1 创建ws 连接
	for _,account := range accounts{
		robot_common.Ws(account.UserName)
	}
	//执行chan读取下注信息
	go RunOnBet()
	go func() {
		for   {
			c := <- robot_common.Client
			t := robot_common.Clients{c.Origin,c.ClienId}
			go Token(t)
			fmt.Println(t)
		}
	}()
	fmt.Println("ssss")
	time.Sleep(2000*time.Second)
	// 投注
	http.ListenAndServe(":8001", nil)
}

func Token(c robot_common.Clients){
	//2 获取token
	resp_token := robot_common.Token(c.ClienId)
	fmt.Println("token",resp_token)
	if resp_token.Code == 200 {
		Token := resp_token.Token
		//3 账户密码登录
		resp_login := robot_common.Login(c.Origin,c.ClienId,Token)
		fmt.Println(resp_login)
		//4进入房间
		fmt.Println("进入房间",c.Origin,c.ClienId)
		resp_room_index := robot_common.Index(c.ClienId,Token,"1")
		fmt.Println(resp_room_index)

		go func() {
			for {
				PreOnBet(resp_login.Token,c)
				time.Sleep(30*time.Second)
			}
		}()

	}
	// 投注
}
/**
	队列下单
 */
func PreOnBet(Token string,c robot_common.Clients){
	fmt.Println("总金额:",robot_common.BetSum)
	if robot_common.BetSum <=50000 {
		//区域随机
		Zone := rand.Intn(6)
		if Zone != 0{
			//下注金额随机
			BetRand := rand.Intn(14)
			bf := &robot_common.BetInfo{}
			bf.Token = Token
			bf.Zone = ckp_common.IntToString(Zone)
			bf.C = c
			bf.Money = ckp_common.StringToInt(Bet[BetRand])
			nn := rand.Intn(160)
			time.Sleep(time.Duration(nn) *time.Second)
			robot_common.BetSumChan <- bf
		}

	}


}
/**
	执行下注
 */
func RunOnBet(){
	for   {
		select {
		case BetInfo := <-robot_common.BetSumChan:
			fmt.Println("总金额:",robot_common.BetSum,BetInfo)
			if robot_common.BetSum <=50000 {
				resp_room := robot_common.OnBet(ckp_common.IntToString(BetInfo.Money),BetInfo.Token,BetInfo.Zone,BetInfo.C.ClienId)
				fmt.Println(resp_room)
				if resp_room.Code == 200 {
					fmt.Println("下注成功")
					robot_common.BetSum += BetInfo.Money
				}
			}else{
				fmt.Println("超过下载金额")
			}
			break
		}
	}
}