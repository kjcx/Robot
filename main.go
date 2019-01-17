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
	//1读取账号
	accounts := robot_common.File("./robot_common/account.txt")
	fmt.Println(accounts)
	//1 创建ws 连接
	for _,account := range accounts{
		robot_common.Ws(account.UserName)
	}
	go func() {
		for   {
			c := <- robot_common.Client
			fmt.Errorf("ccc",c)
			t := robot_common.Clients{c.Origin,c.ClienId}
			go Token(t)
			fmt.Println(t)
		}
	}()
	fmt.Println("ssss")
	time.Sleep(2000*time.Second)
	fmt.Println(robot_common.Client_id)
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
		for  {
			var zone = "2"
			rand.Seed(time.Now().Unix())
			n := rand.Intn(5)
			fmt.Println("n",n)
			if n == 0 {
				zone = "1"
			}else{
				nn :=rand.Intn(5)
				zone = strconv.Itoa(nn)
			}
			time.Sleep(time.Second *1)
			resp_room := robot_common.OnBet(Bet[n],resp_login.Token,zone,c.ClienId)
			fmt.Println(resp_room)
			if resp_room.Code == 200 {
				time.Sleep(150 * time.Second)
			}
			time.Sleep( 10 *time.Second)
		}

	}
	// 投注
}