package wechat

type UserInfo struct {
	Openid string
}

type AppInfo struct {
	AppId  string `json:"appid,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type JsSessionReq struct {
	AppInfo
	Code      string `json:"js_code,omitempty"`
	GrantType string `json:"grant_type,omitempty"`
}

// JsSession 微信小程序js登录返回的会话信息
type JsSession struct {
	SessionKey string `json:"session_key,omitempty"`
	OpenInfo
	Err
}

type OpenInfo struct {
	Openid  string `json:"openid,omitempty"`
	UnionId string `json:"unionid,omitempty"`
}

type Err struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
