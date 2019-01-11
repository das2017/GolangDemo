package models

import (
	"database/sql"
	"errors"

	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type PaymentRecord struct {
	Id           int64
	AccountID    int64
	PartnerID    string
	UserID       string
	CreateTime   string
	Amount       float64
	OuterTradeNo string
	Remark       string
	Status       int
	Msg          string
}
type PaymentRecordStr struct {
	AccountID    string
	PartnerID    string
	UserID       string
	CreateTime   string
	Amount       string
	OuterTradeNo string
	Remark       string
}

func init() {
	dbconn := beego.AppConfig.String("DBConn")
	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(0)
	db.Ping()
	Db = db
}

func Close() {
	if Db != nil {
		Db.Close()

	}

}

func AddPaymenRec(rec PaymentRecordStr) (PaymentRecord, error) {
	var isql = "INSERT pay_demo SET account_id=?,partner_id=?,user_id=?,amount=?,outer_tradeno=?,remark=?"
	AccountID, _ := strconv.ParseInt(rec.AccountID, 10, 64)
	Amount, _ := strconv.ParseFloat(rec.Amount, 64)
	response := PaymentRecord{0, AccountID, rec.PartnerID, rec.UserID, rec.CreateTime, Amount, rec.OuterTradeNo, rec.Remark, 0, ""}
	if Db == nil {
		return response, errors.New("AddPaymenRec connect mysql failed")
	}
	stmt, _ := Db.Prepare(isql)
	defer stmt.Close()
	beego.Informational("AddPaymenRec rec=%#v", rec)
	res, err := stmt.Exec(AccountID, rec.PartnerID, rec.UserID, Amount, rec.OuterTradeNo, rec.Remark)
	if err == nil {
		response.Id, _ = res.LastInsertId()
		response.Status = 1
		response.Msg = "已生效"
		return response, nil
	}

	return response, nil
}
func GetPaymenRec(AccountID int64) (PaymentRecord, error) {
	var qsql = "SELECT * FROM pay_demo WHERE  account_id=?"
	var response PaymentRecord
	response.Msg = "失败"
	if AccountID != 0 {
		if Db == nil {
			return response, errors.New("GetPaymenRec connect mysql failed")
		}
		stmt, _ := Db.Prepare(qsql)
		rows, err := stmt.Query(AccountID)
		defer rows.Close()
		if err != nil {
			return response, err
		}
		var timedate string
		for rows.Next() {
			err = rows.Scan(&response.Id, &response.AccountID, &response.PartnerID, &response.UserID, &timedate, &response.Amount, &response.OuterTradeNo, &response.Remark)
			if err != nil {
				return response, err
			}
			DefaultTimeLoc := time.Local
			loginTime, err := time.ParseInLocation("2006-01-02 15:04:05", timedate, DefaultTimeLoc)
			if err == nil {
				unix_time := loginTime.Unix() //time to int64
				response.CreateTime = time.Unix(unix_time, 0).Format("2006-01-02 15:04:05")
				response.Status = 2
				response.Msg = "成功"
				return response, err
			} else {
				return response, err
			}
		}
		return response, nil
	}
	return response, errors.New("GetPaymenRec Requset is non porinter")
}
