package model

import "github.com/lishimeng/app-starter"

// Driver 驾驶员, 不指定组织
type Driver struct {
	app.Pk
	Code  string `orm:"column(code);unique"`  // 编号
	Name  string `orm:"column(name);unique"`  // 姓名,保护
	Phone string `orm:"column(phone);unique"` // 电话,保护中间四位
	app.TableChangeInfo
}

// UserProfile 用户档案,每个用户必须有
type UserProfile struct {
	app.Pk
	UserCode              string     `orm:"column(user_code)"`                // 用户编号
	RealName              string     `orm:"column(real_name)"`                //  真实姓名
	IdCard                string     `orm:"column(id_card)"`                  //  身份证号
	IdCardVerified        VerifyFlag `orm:"column(id_card_verified)"`         // 身份证号验证标记
	PhoneNumber           string     `orm:"column(phone_number)"`             //  手机号
	PhoneNumberVerified   VerifyFlag `orm:"column(phone_number_verified)"`    // 手机号验证标记
	WechatUnionId         string     `orm:"column(wechat_union_id)"`          // 微信UnionId
	WechatUnionIdVerified VerifyFlag `orm:"column(wechat_union_id_verified)"` // 微信UnionId验证标记
	app.TableChangeInfo
}

type VerifyFlag int

const (
	Verified   = 1
	UnVerified = 0
)
