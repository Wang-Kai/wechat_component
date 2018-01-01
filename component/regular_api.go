// regular tasks api
// these api will be run on time
package component

import (
	"fmt"

	"github.com/goushuyun/log"

	"github.com/franela/goreq"
)

type RegularApi interface {
	GetAccessToken(ticket string) (string, float64)
	GetPreAuthCode(accessToken string) (string, float64)
}

type regularApi struct {
	wt *WechatThird
}

// 获取第三方平台令牌
func (ra *regularApi) GetAccessToken(ticket string) (string, float64) {

	postData := struct {
		Component_appid         string `json:"component_appid"`
		Component_appsecret     string `json:"component_appsecret"`
		Component_verify_ticket string `json:"component_verify_ticket"`
	}{
		Component_appid:         ra.wt.appId,
		Component_appsecret:     ra.wt.appSecret,
		Component_verify_ticket: ticket,
	}

	log.Debug("==================================")
	log.JSON(postData)
	log.Debug("==================================")

	res, err := goreq.Request{
		Method: "POST",
		Uri:    apiComponentToken,
		Body:   postData,
	}.Do()
	if err != nil {
		log.Error("getAccessToken api failed: ", err.Error())
		return "", 0
	}

	result := &struct {
		CAT       string  `json:"component_access_token"`
		ExpiresIn float64 `json:"expires_in"`
	}{}
	err = unmarshalResponseToJson(res, result)
	if err != nil {
		log.Error("Parse access token failed: ", err)

	}
	return result.CAT, result.ExpiresIn

}

// 获取预授权码，用于公众号oauth
func (ra *regularApi) GetPreAuthCode(accessToken string) (string, float64) {
	postData := struct {
		Component_appid string `json:"component_appid"`
	}{
		Component_appid: ra.wt.appId,
	}

	res, err := goreq.Request{
		Method: "POST",
		Uri:    fmt.Sprintf(apiCreatePreAuthCode, accessToken),
		Body:   postData,
	}.Do()
	if err != nil {
		log.Error("getPreAuthCode api failed: ", err.Error())
		return "", 0
	}
	result := &struct {
		PAC       string  `json:"pre_auth_code"`
		ExpiresIn float64 `json:"expires_in"`
	}{}

	err = unmarshalResponseToJson(res, result)
	if err != nil {
		log.Error("Parse pre auth token failed: ", err)
		return "", 0
	}
	return result.PAC, result.ExpiresIn

}
