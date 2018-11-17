package robot_common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
type RespLogin struct {
	Code int32
	Token string
	Message string
}
//账户登录
func Login(client_id ,token string ) *RespLogin{
	urllogin := "http://ss.wmy2.com/Api/User/login/client_id/" +client_id  +"/token/" + token
	postValue := url.Values{
		"user_name": {"a8866"},
		"password": {"123456"},
	}
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
