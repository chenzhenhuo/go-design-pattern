package strategy

import "fmt"

// 支付是采用微信还是支付宝
type IPay interface {
	Pay()
}

var payMethods = map[string]IPay{
	"wx":  &WxPay{},
	"ali": &AliPay{},
}

func NewPayCommon(payMethod string) IPay {
	pay, ok := payMethods[payMethod]
	if !ok {
		return nil
	}
	return pay
}

type WxPay struct {
}

func (w *WxPay) Pay() {
	fmt.Println("WxPay")
}

type AliPay struct {
}

func (a *AliPay) Pay() {
	fmt.Println("AliPay")
}
