package subsystemChatgptApi

import (
	"fmt"
	component_ktv_ai_lib "github.com/tangwenru/go-component-ktv-ai/lib"
	"github.com/tangwenru/go-component-ktv-ai/subsystem/chatgpt-api/type"
)

type Vip struct {
}

func init() {

}

// @VersionCategory ： 3.5 | 4.0 ...
func (this *Vip) PayedUpdate(userId int64, versionCategory string, dollarAmount float64) (error, *typeChatgptApi.VipPayedUpdateData) {
	query := typeChatgptApi.VipPayedUpdateQuery{
		VersionCategory: versionCategory,
		DollarAmount:    dollarAmount,
	}
	vipPayedUpdate := typeChatgptApi.VipPayedUpdateData{}

	bytesResult, err := component_ktv_ai_lib.SubsystemChatgptApi(userId, "vip/PayedUpdate", &query, &vipPayedUpdate)

	if err != nil {
		fmt.Println("SubsystemChatgptApi) Detail:", string(bytesResult))
	}

	return err, &vipPayedUpdate
}
