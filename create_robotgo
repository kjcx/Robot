package main

import (
	"Robot/robot_common"
	"ckp/ckp_common"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
	_"github.com/go-sql-driver/mysql"
)

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
}
var GormDb  *gorm.DB
var GormErr  error
func main(){
	//接收参数创建机器人数量
	sql := fmt.Sprintf("INSERT INTO `zhongcai`.`zc_user`(`user_id`, `user_name`, `nickname`, `phone`, `email`, `password`, `pay_password`, `balance`, `status`, `invite_code`, `add_time`) VALUES (null, %s, %s, '', '', 'e10adc3949ba59abbe56e057f20f883e', 'e10adc3949ba59abbe56e057f20f883e', 9991952.10, 1, '000000', 1547380054)","jq0021","jq0021")
	fmt.Println(sql)
	u := User{}
	u.AddTime = time.Now().Unix()
	u.Balance = 10000000
	u.Status = 1
	u.UserName = GetRandomString(6)
	u.Nickname = GetRandomString(6)
	md5 := "123456"
	u.Password = ckp_common.Md5(md5)
	u.PayPassword = ckp_common.Md5(md5)
	GormDb.Create(&u)

	content := u.UserName + "," + md5
	robot_common.AppendToFile("./robot_common/account.txt","\n" + content)
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
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