package main

import (
	"Robot/robot_common"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
import "fmt"
var Bet = []string{"10","50","100","500","1000"}
func main(){
	//1 创建ws 连接
	robot_common.Ws()
	time.Sleep(2*time.Second)
	fmt.Println(robot_common.Client_id)
	//2 获取token
	resp_token := robot_common.Token(robot_common.Client_id)
	fmt.Println("token",resp_token)
	if resp_token.Code == 200 {
		Token := resp_token.Token
		//3 账户密码登录
		resp_login := robot_common.Login(robot_common.Client_id,Token)
		fmt.Println(resp_login)
		//4进入房间
		fmt.Println("进入房间")
		resp_room_index := robot_common.Index(robot_common.Client_id,Token,"1")
		fmt.Println(resp_room_index)
		for  {
			var zone = "1"
			n := rand.Intn(5)
			if n == 0 {
				zone = "1"
			}else{
				zone = strconv.Itoa(n)
			}
			resp_room := robot_common.OnBet(Bet[n],resp_login.Token,zone,robot_common.Client_id)
			fmt.Println(resp_room)
			if resp_room.Code == 200 {
				time.Sleep(150 * time.Second)
			}
			time.Sleep( 10 *time.Second)
		}

	}



	// 投注
	http.ListenAndServe(":8001", nil)
}
