package main

import (
	"Robot/robot_common"
	"bytes"
	"ckp/ckp_common"
	"flag"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
	_"github.com/go-sql-driver/mysql"
)
import "fmt"
var Bet = []string{"10","50","100","50","500","10","1000","10","100","50","1000","10","100","50","10"}
var GormDb  *gorm.DB
var GormErr  error
type User struct {
	UserId int
	UserName string
	Nickname string
	Phone string
	Email string
	Password string
	PayPassword string
	Balance float64
	Status int64
	InviteCode string
	AddTime int64
	IsRobot int64
}

func init(){
	//导入配置文件
	configMap := make(map[string]string)
	configMap["user"] = "zhongcai"
	configMap["password"] = "zhongcai123"
	configMap["host"] = "127.0.0.1"
	configMap["port"] = "3306"
	configMap["database"] = "zhongcai"
	//获取配置里host属性的value
	fmt.Println(configMap["host"])
	//查看配置文件里所有键值对
	fmt.Println(configMap)

	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.31.232:3306)/chuangke?charset=utf8")
	mysqlurl := configMap["user"] + ":" + configMap["password"] + "@tcp(" + configMap["host"] + ":" + configMap["port"] + ")/" + configMap["database"] + "?charset=utf8"
	//gorm_model
	GormDb, GormErr = gorm.Open("mysql", mysqlurl)
	if GormErr != nil {
		fmt.Println(GormErr)
	}
	// 全局禁用表名复数
	GormDb.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	////设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "zc_" + defaultTableName
	}
	GormDb.Debug()
	//defer GormDb.Close()
}
var nums  = flag.Int("num", 50, "robot num")
func main(){
	flag.Parse()
	u := User{}
	u.IsRobot = 1
	data := []*User{}
	//定义一个int型字符
	fmt.Println(*nums,*nums)
	//time.Sleep(10*time.Second)
	GormDb.Where(&u).Limit(*nums).Find(&data)
	fmt.Println(data)
	//1读取账号
	//accounts := robot_common.File("./robot_common/account.txt")
	//fmt.Println(accounts)
	//1 创建ws 连接
	for _,v := range data{
		robot_common.Ws(v.UserName)
	}
	//for _,account := range accounts{
	//	robot_common.Ws(account.UserName)
	//}
	fmt.Println("ssss")
	go Select()
	go func() {
		for{
			fmt.Println("在线人数:", len(robot_common.ClientOnline),robot_common.Game1OnBetStatus)
			time.Sleep(10*time.Second)
			if len(robot_common.ClientOnline) > 2 {
				//robot_common.ClientOnline[0].Ws.Close()
				//robot_common.ClientOnline[0].Ws = nil
				//fmt.Println(robot_common.ClientOnline[0].Ws.IsClientConn())
				//robot_common.ClientOnline = robot_common.ClientOnline[1:]
			}

			//robot_common.ClientOnline[0].Ws.IsClientConn()
		}

	}()
	//执行chan读取下注信息
	go RunOnBet()
	go func() {
		for   {
			c := <- robot_common.Client
			t := robot_common.Clients{Origin: c.Origin, ClienId: c.ClienId}

			go Token(t)
			fmt.Println(t)
		}
	}()
	fmt.Println("ssss")
	time.Sleep(2000*time.Second)
	// 投注
	http.ListenAndServe(":8001", nil)
}
var RobotNum int64
var RobotMoney int64
func Select(){

	type Site struct{
		RobotNum int64
		RobotMoney int64
	}
	s := Site{}
	for  {
		fmt.Println("RobotNumRobotNumRobotNum")
		GormDb.Raw("select robot_num,robot_money from zc_site where site_id = 1").Find(&s)
		fmt.Println("RobotNum",s.RobotNum)
		if RobotNum == 0 {
			RobotNum = s.RobotNum
		}else if RobotNum > 0 && RobotNum != s.RobotNum {
			RobotNum = s.RobotNum
			cmd := fmt.Sprintf("/www/robot.sh %d",RobotNum)
			exec_shell(cmd)
		}
		RobotMoney = s.RobotMoney
		fmt.Println("RobotMoney:",RobotMoney)
		time.Sleep(10*time.Second)
	}
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
				if robot_common.Game1OnBetStatus == 1 {
					PreOnBet(resp_login.Token,c)
					time.Sleep(80*time.Second)
				}else{
					time.Sleep(10*time.Second)
				}
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
	if robot_common.BetSum <= int(RobotMoney) {
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
			if robot_common.BetSum <=int(RobotMoney) {
				resp_room := robot_common.OnBet(ckp_common.IntToString(BetInfo.Money),BetInfo.Token,BetInfo.Zone,BetInfo.C.ClienId)
				fmt.Println(resp_room)
				if resp_room.Code == 200 {
					fmt.Println("下注成功")
					robot_common.BetSum += BetInfo.Money
				}else if resp_room.Code == 3033 {
					//不可以下注
					robot_common.Game1OnBetStatus = 0
				}
			}else{
				fmt.Println("超过下载金额")
			}
			break
		}
	}
}
//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)
	return out.String(), err
}
//错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
