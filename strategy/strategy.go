package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name string, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardID,
			Money:  money,
		},
		strategy: strategy,
	}
}

type PaymentContext struct {
	Name   string
	CardID string
	Money  int
}

type PaymentStrategy interface {
	Pay(context *PaymentContext)
}

type AliPay struct{}

func (a *AliPay) Pay(context *PaymentContext) {
	fmt.Println("Pay $%d to %s by AliPay with cardID %s", context.Money, context.Name, context.CardID)
}

type WechatPay struct{}

func (w *WechatPay) Pay(context *PaymentContext) {
	fmt.Println("Pay $%d to %s by WechatPay with cardID %s", context.Money, context.Name, context.CardID)
}
