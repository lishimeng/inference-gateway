package users

import (
	"github.com/lishimeng/app-starter"
	model "github.com/lishimeng/inference-gateway/internal/db"
)

// CreateProfile 创建Profile
// @param code 账号编号
func CreateProfile(code string) {
	profile := model.UserProfile{
		UserCode:              code,
		IdCardVerified:        model.UnVerified,
		PhoneNumberVerified:   model.UnVerified,
		WechatUnionIdVerified: model.UnVerified,
	}
	app.GetOrm().Context.Insert(&profile)
}
