package controllers

import (
	"encoding/json"

	"github.com/GolangDemo/webApiDemo/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type PayController struct {
	beego.Controller
}

// @Title PayQuery
// @Description PayQuery
// @Param	body		body 	models.PaymentQueryRequest	true		"body for user content"
// @Success 200 {paymentRecord} models.PaymentRecord
// @router /PayQuery [post]
func (u *PayController) PayQuery() {
	var request models.PaymentQueryRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &request)
	payResponse, _ := models.GetPaymenRec(request.AccountID)
	u.Data["json"] = payResponse
	u.ServeJSON()
}

// @Title PayAdd
// @Description PayAdd
// @Param	body		body 	models.PaymentAddRequset	true		"body for user content"
// @Success 200 {paymentRecord} models.PaymentRecord
// @Failure 403 body is empty
// @router /PayAdd [post]
func (u *PayController) PayAdd() {
	var payRec models.PaymentAddRequset
	json.Unmarshal(u.Ctx.Input.RequestBody, &payRec)
	payResponse, _ := models.AddPaymenRec(payRec)
	u.Data["json"] = payResponse
	u.ServeJSON()
}
