package service

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	request3 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/response"
	"log"
	order2 "saas/biz/dal/db/ent/order"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/wechat"
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

func NewWXPayment(ctx context.Context, c *app.RequestContext) do.Payment {
	return &Payment{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (p Payment) UnifyPay(req payment.PayReq) (interface{}, error) {

	attach := "自定义数据说明"
	description := "Image形象店-深圳腾大-QQ公仔"

	options := &request2.RequestJSAPIPrepay{
		Amount: &request2.JSAPIAmount{
			Total:    int(req.Total * 100),
			Currency: "CNY",
		},
		Attach:      attach,
		Description: description,
		OutTradeNo:  req.OrderSn, // 这里是商户订单号，不能重复提交给微信
		Payer: &request2.JSAPIPayer{
			OpenID: req.OpenId, // 用户的openid， 记得也是动态的。
		},
	}

	// 如果需要覆盖掉全局的notify_url
	//options.SetNotifyUrl("https://pay.xxx.com/wx/notify")
	// 下单
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response2, err := wechat.PaymentWechatApp.Order.JSAPITransaction(ctx, options)

	if err != nil {
		hlog.Infof("error: %s", err)
		return response2, err
	}

	payConf, err := wechat.PaymentWechatApp.JSSDK.BridgeConfig(response2.PrepayID, false)
	if err != nil {
		return nil, err
	}

	return payConf, nil
}

// QRPay 支付二维码
func (p Payment) QRPay(req payment.PayReq) (interface{}, error) {
	attach := "自定义数据说明"
	description := "Image形象店-深圳腾大-QQ公仔"
	options := &request2.RequestNativePrepay{
		Amount: &request2.NativeAmount{
			Total:    int(req.Total * 100),
			Currency: "CNY",
		},
		Attach:      attach,
		Description: description,
		OutTradeNo:  req.OrderSn, // 这里是商户订单号，不能重复提交给微信
	}

	response2, err := wechat.PaymentWechatApp.Order.TransactionNative(p.ctx, options)

	if err != nil {
		hlog.Infof("error: %s", err)
		return response2, err
	}

	return response2, nil
}

// Notify 回调

func (p Payment) Notify(c interface{}) {
	//rs, err := PaymentApp.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
	//rs, err := PaymentApp.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
	//_, err := services.PaymentApp.HandlePaidNotify(c.Request, func(chat-bot *power.HashMap, content *power.HashMap, fail string) interface{} {
	//if content == nil || (*content)["out_trade_no"] == nil {
	//  return fail("no content notify")
	//}
	//return true
	//})

	req, err := adaptor.GetCompatRequest(&p.c.Request)
	if err != nil {
		hlog.Info(err)
		return
	}
	hlog.Info(req)
	res, err := wechat.PaymentWechatApp.HandlePaidNotify(req,
		func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {

			// 看下支付通知事件状态
			// 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
			if message.EventType != "TRANSACTION.SUCCESS" {
				fmt.Dump(transaction)
				return true
			}

			if transaction.OutTradeNo != "" {

				order, err := p.db.Order.Query().Where(order2.OrderSn(transaction.OutTradeNo)).First(p.ctx)
				if err != nil {
					hlog.Info(transaction.OutTradeNo + "订单不存在")
				}

				p.db.Order.UpdateOne(order).SetCompletionAt(time.Now()).Save(p.ctx)

				amount, _ := order.QueryAmount().First(p.ctx)

				p.db.OrderAmount.UpdateOne(amount).
					SetActual(0).
					Save(p.ctx)
				fee := transaction.Amount.Total * 100
				p.db.OrderPay.Create().
					SetOrder(order).
					SetPay(float64(fee)).
					SetPayAt(time.Now()).
					SetPayWay("wxPay").
					SetPaySn(transaction.OutTradeNo).
					SetPrepayID(transaction.TransactionID).
					//SetPayExtra(transaction).
					Save(p.ctx)

				// 这里对照自有数据库里面的订单做查询以及支付状态改变
				log.Printf("订单号：%s 支付成功", transaction.OutTradeNo)
			} else {
				// 因为微信这个回调不存在订单号，所以可以告诉微信我还没处理成功，等会它会重新发起通知
				// 如果不需要，直接返回true即可
				fail("payment fail")
				return nil
			}
			return true
		},
	)

	// 这里可能是因为不是微信官方调用的，无法正常解析出transaction和message，所以直接抛错。
	if err != nil {
		res.StatusCode = 502
		err = res.Write(p.c.Request.BodyWriter())
		return
		//panic(err)
	}

	// 这里根据之前返回的是true或者fail，框架这边自动会帮你回复微信
	err = res.Write(p.c.Request.BodyWriter())
	if err != nil {
		panic(err)
	}

}

// RefundOrder 退款
func (p Payment) RefundOrder(req payment.RefundOrderReq) (*response.ResponseRefund, error) {
	transactionID := req.TransactionId
	outRefundNo := req.OutRefundNo

	wechat.PaymentWechatApp.GetConfig().GetString("app_id", "")

	options := &request3.RequestRefund{
		TransactionID: transactionID,
		OutRefundNo:   outRefundNo,
		Reason:        "",
		//NotifyUrl:     "", // 异步接收微信支付退款结果通知的回调地址
		FundsAccount: "",
		Amount: &request3.RefundAmount{
			Refund: 1, // 退款金额，单位：分
			Total:  1, // 订单总金额，单位：分
			From: []*request3.RefundAmountFrom{
				&request3.RefundAmountFrom{
					Account: "AVAILABLE",
					Amount:  1,
				},
			}, // 退款出资账户及金额。不传仍然需要这个空数组防止微信报错
			Currency: "CNY",
		},
	}

	rs, err := wechat.PaymentWechatApp.Refund.Refund(p.ctx, options)
	if err != nil {
		hlog.Infof("error: %s", err)
		return nil, err
	}
	//c.JSON(http.StatusOK, rs)

	return rs, nil
}
