package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golangDemo/webDemo/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) PayQuery() {
	AccountID, _ := c.GetInt64("AccountID1")
	payment, _ := models.GetPaymenRec(AccountID)
	c.Data["AccountID"] = payment.AccountID
	c.Data["PartnerID"] = payment.PartnerID
	c.Data["UserID"] = payment.UserID
	c.Data["CreateTime"] = payment.CreateTime
	c.Data["Amount"] = payment.Amount
	c.Data["OuterTradeNo"] = payment.OuterTradeNo
	c.Data["Remark"] = payment.Remark
	c.Data["Status"] = payment.Status
	c.Data["Msg"] = payment.Msg
	c.TplName = "query.html"
}
func (c *MainController) PayAdd() {
	var payment models.PaymentRecordStr
	c.ParseForm(&payment)
	pay, _ := models.AddPaymenRec(payment)
	c.Data["AccountID"] = pay.AccountID
	c.Data["PartnerID"] = pay.PartnerID
	c.Data["UserID"] = pay.UserID
	c.Data["CreateTime"] = pay.CreateTime
	c.Data["Amount"] = pay.Amount
	c.Data["OuterTradeNo"] = pay.OuterTradeNo
	c.Data["Remark"] = pay.Remark
	c.TplName = "query.html"
}

func (c *MainController) Home() {
	c.TplName = "index.tpl"
}
