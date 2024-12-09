package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/payment"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
// https://powerwechat.artisan-cloud.com/zh/mini-program/user-info.html
type Payment struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewPayment(ctx context.Context, c *app.RequestContext) do.Payment {
	return &Payment{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (p Payment) WXPay(req payment.WXPayReq) error {
	//options := &requestV3POR.RequestJSAPIPrepay{
	//	Amount: &requestV3POR.JSAPIAmount{
	//		Total:    1,
	//		Currency: "CNY",
	//	},
	//	Attach:      "自定义数据说明",
	//	Description: "Image形象店-深圳腾大-QQ公仔",
	//	OutTradeNo:  "55197789397733956592224980201", // 这里是商户订单号，不能重复提交给微信
	//	Payer: &requestV3POR.JSAPIPayer{
	//		OpenID: "o4QEk5Mf1Do7utS7-SF5Go30s8i4", // 用户的openid， 记得也是动态的。
	//	},
	//}
	//
	//// 如果需要覆盖掉全局的notify_url
	////options.SetNotifyUrl("https://pay.xxx.com/wx/notify")
	//
	//// 下单
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//response, err := wechat.PaymentWechatApp.Order.JSAPITransaction(ctx, options)
	//
	//if err != nil {
	//	log.Printf("error: %s", err)
	//	c.JSON(400, response)
	//	return
	//}
	//
	//payConf, err := wechat.PaymentWechatApp.JSSDK.BridgeConfig(response.PrepayID, false)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//c.JSON(200, payConf)

	return nil
}

func (p Payment) WXNotify(req payment.WXNotifyReq) error {

	//res, err := wechat.PaymentWechatApp.HandlePaidNotify(
	//	c.Request,
	//	func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {
	//
	//		// 看下支付通知事件状态
	//		// 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
	//		if message.EventType != "TRANSACTION.SUCCESS" {
	//			fmt.Dump(transaction)
	//			return true
	//		}
	//
	//		if transaction.OutTradeNo != "" {
	//			// 这里对照自有数据库里面的订单做查询以及支付状态改变
	//			log.Printf("订单号：%s 支付成功", transaction.OutTradeNo)
	//		} else {
	//			// 因为微信这个回调不存在订单号，所以可以告诉微信我还没处理成功，等会它会重新发起通知
	//			// 如果不需要，直接返回true即可
	//			fail("payment fail")
	//			return nil
	//		}
	//		return true
	//	},
	//)
	//
	//// 这里可能是因为不是微信官方调用的，无法正常解析出transaction和message，所以直接抛错。
	//if err != nil {
	//	res.StatusCode = 502
	//	err = res.Write(c.Writer)
	//	return
	//	//panic(err)
	//}
	//
	//// 这里根据之前返回的是true或者fail，框架这边自动会帮你回复微信
	//err = res.Write(c.Writer)
	//
	//if err != nil {
	//	panic(err)
	//}

	return nil
}

// APIRefundOrder 退款
func APIRefundOrder() {
	//transactionID := c.DefaultQuery("transactionID", "")
	//outRefundNo := c.DefaultQuery("OutRefundNo", "")
	//
	//wechat.PaymentWechatApp.GetConfig().GetString("app_id", "")
	//
	//options := &requestV3PRR.RequestRefund{
	//	TransactionID: transactionID,
	//	OutRefundNo:   outRefundNo,
	//	Reason:        "",
	//	//NotifyUrl:     "", // 异步接收微信支付退款结果通知的回调地址
	//	FundsAccount: "",
	//	Amount: &requestV3PRR.RefundAmount{
	//		Refund: 1, // 退款金额，单位：分
	//		Total:  1, // 订单总金额，单位：分
	//		From: []*requestV3PRR.RefundAmountFrom{
	//			&requestV3PRR.RefundAmountFrom{
	//				Account: "AVAILABLE",
	//				Amount:  1,
	//			},
	//		}, // 退款出资账户及金额。不传仍然需要这个空数组防止微信报错
	//		Currency: "CNY",
	//	},
	//	GoodsDetail: []*requestV3PRR.RefundGoodDetail{
	//		&requestV3PRR.RefundGoodDetail{
	//			MerchantGoodsID:  "1217752501201407033233368018",
	//			WechatPayGoodsID: "1001",
	//			GoodsName:        "公仔",
	//			UnitPrice:        1,
	//			RefundAmount:     1,
	//			RefundQuantity:   1,
	//		},
	//	},
	//}
	//
	//rs, err := wechat.PaymentWechatApp.Refund.Refund(c, options)
	//if err != nil {
	//	panic(err)
	//}
	//c.JSON(http.StatusOK, rs)

	return
}
