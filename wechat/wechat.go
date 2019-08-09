package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WechatWork struct {
	client      *http.Client
	CorpId      string
	CorpSecret  string
	AccessToken string
}

func NewWechatWork() (wechat *WechatWork) {
	ID := os.Getenv("WECHAT_WORK_ID")
	SECRET := os.Getenv("WECHAT_WORK_SECRET")
	if ID == "" || SECRET == "" {
		log.Fatalf("Please set environments.")
	}
	return &WechatWork{
		client:     &http.Client{},
		CorpId:     ID,
		CorpSecret: SECRET,
	}
}

func (wechat *WechatWork) request(endpoint string) (data map[string]interface{}, errcode int) {
	resp, err := wechat.client.Get(endpoint)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("%v", err)
	}

	// log.Printf("%v", result)
	errcode = int(result["errcode"].(float64))

	return result, errcode
}

func (wechat *WechatWork) Gettoken() (ret bool) {
	endpoint := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", wechat.CorpId, wechat.CorpSecret)
	result, _ := wechat.request(endpoint)

	if a, ok := result["access_token"]; ok {
		wechat.AccessToken = a.(string)
		ret = true
	} else {
		wechat.AccessToken = ""
		ret = false
	}

	return ret
}

func (wechat *WechatWork) Getuserinfo(code string) (UserId string, errcode int) {
	endpoint := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s", wechat.AccessToken, code)
	result, errcode := wechat.request(endpoint)

	UserId = ""
	if errcode != 0 {
		log.Printf("%v", result)
	} else {
		if uid, ok := result["UserId"]; ok {
			UserId = uid.(string)
			errcode = 0
		} else {
			UserId = ""
			errcode = 1
		}
	}

	return UserId, errcode
}

func (wechat *WechatWork) Getuser(UserId string) (user map[string]interface{}, errcode int) {
	endpoint := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s", wechat.AccessToken, UserId)
	result, errcode := wechat.request(endpoint)

	user = nil
	if errcode != 0 {
		log.Printf("%v", result)
	} else {
		user = result
	}

	return result, errcode
}
