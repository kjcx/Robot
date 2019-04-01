package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"time"
)
type Resp struct {
	Data Datas
}
type Datas struct {
	Product ProductData
	PropertyNames []*Name
}
type Name struct {
	Id int32
	Name string
	Values []*Value
}
type Value struct {
	Id int32
	Value string
	Weight int32
}
type ProductData struct {
	Id int32
	Name string
	TopPrice int32
}
func main(){
	url := "https://m.aihuishou.com/portal-api/product/inquiry-detail-new/25676?quickInquiryValue="
	resp,_ := http.Get(url)
	data,_ := ioutil.ReadAll(resp.Body)
	rp := &Resp{}
	json.Unmarshal(data,rp)
	fmt.Println(rp)
	fmt.Printf("%+v\n", rp)
	rp.Insert()
	time.Sleep(1000*time.Second)
}
func (r *Resp)Insert(){
	r.Db().Insert(r)
}
func (r *Resp)Db()*mgo.Collection{
	c := Session.DB("ahs").C("Product")
	return c
}
var  Session *mgo.Session
var err error
var GoRedis *redis.Client
func init()  {
	//go ckp_common.SubPush()
	//go Timer()
	//go ckp_common.SubPush()
	//configs := config.GetConfig()
	//fmt.Println(configs)
	Session, err = mgo.Dial("localhost:27017")
	//
	//dialInfo := &mgo.DialInfo{
	//	Addrs:     []string{configs["mongohost"] + ":"+ configs["mongoport"] },
	//	Direct:    false,
	//	Timeout:   time.Second * 10,
	//	Database:  configs["mongoudb"],
	//	Source:    "",
	//	Username:  configs["mongouser"],
	//	Password:  configs["mongopassword"],
	//	PoolLimit: 4096, // Session.SetPoolLimit
	//}
	//Session,err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("mongo连接错误")
	}

}