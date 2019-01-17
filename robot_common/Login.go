package robot_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)
type RespLogin struct {
	Code int32
	Token string
	Message string
}
//账户登录
func Login(UserName string,client_id ,token string ) *RespLogin{
	time.Sleep(100 * time.Microsecond)
	urllogin := "http://ss.wmy2.com/Api/User/login/client_id/" +client_id  +"/token/" + token
	postValue := url.Values{
		"user_name": {UserName},
		"password": {"123456"},
	}
	fmt.Println("login:",urllogin,postValue)
	resp, err := http.PostForm(urllogin, postValue)
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := RespLogin{}
	json.Unmarshal(body,&r)
	return &r
}
