package component_ktv_ai

import (
	"fmt"
	//component_ktv_ai_lib "component_ktv_ai/lib"
	component_ktv_ai_lib "github.com/tangwenru/go-component-ktv-ai/lib"
)

type User struct {
}

type UserDetail struct {
	Id          int64  `json:"id"`
	AccountName string `json:"accountName"`
	Nickname    string `json:"nickname"`
	AvatarUrl   string `json:"avatarUrl"`
	Enabled     bool   `json:"enabled"`
	Expired     int64  `json:"expired"`
	WebSid      string `json:"webSid"`
	Role        string `json:"role"`
	AppSid      string `json:"appSid"`
	ClientSid   string `json:"clientSid"`
	Created     int64  `json:"created"`
}

func init() {

}

func (this *User) Detail(userId int64) (error, *UserDetail) {
	userDetail := UserDetail{}
	query := map[string]string{}
	bytesResult, err := component_ktv_ai_lib.MainSystem(userId, "user/info", &query, &userDetail)

	if err != nil {
		fmt.Println("User info:", string(bytesResult))
	}

	return err, &userDetail
}
