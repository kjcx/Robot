package robot_common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)
type Resp struct {
	Code int32
	Token string
	Msg string

}
//获取token
func Token(client_id string) *Resp{
	url := "http://ss.wmy2.com/Api/Public/getToken/client_id/" + client_id
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := Resp{}
	json.Unmarshal(body,&r)
	return &r
}
