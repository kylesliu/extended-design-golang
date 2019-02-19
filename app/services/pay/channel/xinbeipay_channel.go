package pay_channel

import (
	"extended-design-golang/app/services/pay"
)

type PayChannelXinBeiPay struct {
	pay.PayChannelInterface
}

func (self PayChannelXinBeiPay) GetPayInstance() {
}

func (self PayChannelXinBeiPay) GetPayInfo() string {
	return "xinbei"
}
func (self PayChannelXinBeiPay) GetPayCallback() string {
	return "xinbei"
}
