package subsystemChatgpt

import (
	component_ktv_ai_lib "component_ktv_ai/lib"
	typeChatgpt "component_ktv_ai/subsystem/chatgpt/type"
	"fmt"
)

type Vip struct {
}

type VipDetail struct {
	Id                int64
	UserId            int64
	ProductType       string
	RatioDownPercent1 float64
	RatioDownPercent2 float64
	VipLevel          int64
	IsVip             bool
	Expired           int64
	Created           int64
	Updated           int64
}

func init() {

}

// 支付后 最后的收尾更新
func (this *Vip) PayedUpdate(userId int64, productType string, dollarAmount float64) (error, *typeChatgpt.VipPayedUpdateData) {
	query := typeChatgpt.VipPayedUpdateQuery{
		ProductType:  productType,
		DollarAmount: dollarAmount,
	}
	vipPayedUpdate := typeChatgpt.VipPayedUpdateData{}

	bytesResult, err := component_ktv_ai_lib.SubsystemChatgpt(userId, "vip/PayedUpdate", &query, &vipPayedUpdate)

	if err != nil {
		fmt.Println("UserVip) Detail:", string(bytesResult))
		//json.Unmarshal(bytesResult, &userVipDetail)
	}

	return err, &vipPayedUpdate
}