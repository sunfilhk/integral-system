package models

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"

	"github.com/tealeg/xlsx"
)

type Investment struct {
	ID           string  `gorm:"column:id" json:"id"`
	Userid       string  `gorm:"column:userid" json:"userid"`
	Amount       float64 `gorm:"column:amount" json:"amount"`         //金额
	Customerid   string  `gorm:"column:customerid" json:"customerid"` //客户id
	Remark       string  `gorm:"column:remark" json:"remark"`         //备注
	IsDel        int     `gorm:"column:is_del" json:"is_del"`         //是否删除
	Status       int     `gorm:"column:status" json:"status"`
	Customername string  `gorm:"column:customername" json:"customername"`
	NickName     string  `gorm:"column:nick_name" json:"nick_name"`
	//Customerprofit float64    `gorm:"column:customerprofit" json:"customerprofit"`
	Salesman   float64   `gorm:"column:salesman" json:"salesman"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` //创建时间
	Alterbool  int       `gorm:"column:alterbool" json:"alterbool"`
	Profit     float64   `gorm:"column:profit" json:"profit"`
	Flag1      int       `gorm:"column:flag1" json:"flag1"`
	TeamId     int64     `gorm:"column:teamid" json:"teamid"`           //新加字段
	TeamName   string    `gorm:"column:teamname" json:"teamname"`       //新加字段
	Phone      string    `gorm:"column:phone" json:"phone"`             //新加字段
	Creation   time.Time `gorm:"column:creation" json:"creation"`       //新加字段
	InvestTime time.Time `gorm:"column:invest_time" json:"invest_time"` //新加字段
}
type Investlist struct {
	Id             string    `gorm:"column:id" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	IdCar          string    `gorm:"column:idcar" json:"id_car"`
	BankNum        string    `gorm:"column:banknum" json:"bank_num"`
	Amount         string    `gorm:"column:amount" json:"amount"`
	ProFit         string    `gorm:"column:profit" json:"profit"`
	ProFits        string    `gorm:"column:profits" json:"profits"`
	TeamName       string    `gorm:"column:teamname" json:"teamname"`
	InvestTime     time.Time `gorm:"column:investtime" json:"invest_time"`
	ExpirationDate time.Time `gorm:"column:expirationdate" json:"expiration_date"`
	Remark         string    `gorm:"column:remark" json:"remark"`
	NickName       string    `gorm:"column:nickname" json:"nick_name"`
	Status         string    `gorm:"column:status" json:"status"`
	OrderTime      string    `gorm:"column:ordertime" json:"order_time"`
	Form           int       `gorm:"column:form" json:"form"`
	Customerid     string    `gorm:"column:customerid" json:"customerid"`
	Alterbool      int       `gorm:"column:alterbool" json:"alterbool"`
	Flag1          int       `gorm:"column:flag1" json:"flag1"`
	UserId         int       `gorm:"column:user_id" json:"userid"`
}

type InvestmentExport struct {
	ID             string `gorm:"column:id" json:"id"`
	Userid         string `gorm:"column:userid" json:"userid"`
	Amount         string `gorm:"column:amount" json:"amount"`         //金额
	Customerid     string `gorm:"column:customerid" json:"customerid"` //客户id
	Remark         string `gorm:"column:remark" json:"remark"`         //备注
	IsDel          string `gorm:"column:is_del" json:"is_del"`         //是否删除
	Status         string `gorm:"column:status" json:"status"`
	CreateTime     string `gorm:"column:create_time" json:"create_time"`         //创建时间
	UpdateTime     string `gorm:"column:update_time" json:"update_time"`         //修改时间
	InvestTime     string `gorm:"column:invest_time" json:"invest_time"`         //创建时间
	ExpirationDate string `gorm:"column:expiration_date" json:"expiration_date"` //创建时间
	MonthlyTime    string `gorm:"column:monthly_time" json:"monthly_time"`       //创建时间
	Profit         string `gorm:"column:profit" json:"profit"`
	Oncestatus     string `gorm:"column:oncestatus" json:"oncestatus"`
}

type InvestmentprofitExport struct {
	ID           string `gorm:"column:id" json:"id"`
	Investmentid string `gorm:"column:investmentid" json:"investmentid"`
	Userid       string `gorm:"column:userid" json:"userid"`   //金额
	Profits      string `gorm:"column:profits" json:"profits"` //客户id
}
type InvestmentView struct {
	Investment
	Summaryid string `json:"summaryid"`
}

func NewInvestment(param map[string]string) error {
	sql2 := `select c.*,i.amount,i.remark from customer c
						left join investment i on c.id = i.customerid and i.id = '` + param["investmentid"] + `'
						where c.id = '` + param["customerid"] + "'"
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error; err != nil {
		fmt.Errorf("获取客户信息失败：", err)
		return err
	}
	param["name"] = c.Name
	param["phone"] = c.Phone
	param["bank"] = c.Bank
	param["banknum"] = c.Banknum
	param["identity"] = c.Identity
	param["sex"] = c.Sex
	param["identity"] = c.Identity
	param["remark"] = c.Remark
	param["amount"] = utils.Float64ToString(c.Amount)
	return nil
}
func GetInvestmentById(id string) Investment {
	sql2 := ` select * from investment where id = '` + id + "'"
	var in Investment
	if err := orm.Eloquent.Raw(sql2).Scan(&in).Error; err != nil {
		return in
	}
	return in
}

func (e *Investment) InvestmentList(param map[string]string) (result interface{}, err error) {
	con := ",0 flag1"
	if _, ok := AUDIT[param["roleKey"]]; ok {
		con = ",1 flag1"
	}
	//拼凑筛选条件sql
	sql := ` select i.*,i.create_time creation,TIMESTAMPDIFF(day,invest_time,now())>=3 alterbool` + con + ` from investment i
					where i.is_del=0 `
	status := param["customerid"]
	if param["customerid"] != "" {
		sql += ` and customerid = ` + status
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Investment, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Investlist) InvestLists(param map[string]string) (result interface{}, err error) {
	con := ",0 flag1"
	if _, ok := AUDIT[param["roleKey"]]; ok {
		con = ",1 flag1"
	}
	//g:=param["start"]

	//fmt.Println("shuhu",g)

	//拼凑筛选条件sql
	sql := ` SELECT f.*,ROUND(amount*profit,2)profits FROM (select i.id,i.invest_time investtime,u.user_id,i.expiration_date expirationdate,i.amount,i.customerid,
				i.is_del,i.profit,i.status,u.nick_name nickname,u.team_name teamname,u.team_id,c.bank banknum,
				c.name,i.create_time ordertime,c.phone,c.identity idcar,i.remark, TIMESTAMPDIFF(day,invest_time,now())>=3 alterbool	` + con + ` from investment i 
				LEFT JOIN customer c on i.customerid = c.id LEFT JOIN sys_user u on i.userid=u.user_id WHERE i.status in (0,1,3,4,5,6)`

	// 投资时间查询
	if param["start"] != "" && param["end"] != " 23:59:59" {
		sql += ` and  i.invest_time <= '` + param["end"] + `' and '` + param["start"] + `' <= i.invest_time`
	}

	// 创建时间查询
	if param["createStart"] != "" && param["createEnd"] != " 23:59:59" {
		sql += ` and i.create_time <= '` + param["createEnd"] + `' and '` + param["createStart"] + `' <= i.create_time`
	}

	sql += ` GROUP BY i.id)f where f.is_del=0 `

	//团队类型
	deptid := param["deptid"]
	if param["deptid"] != "" {
		sql += ` and  f.team_id  = ` + deptid + ``
	}
	//投资id
	id := param["id"]
	if param["id"] != "" {
		sql += ` and f.id  = '` + id + `'`
	}

	sb := param["userid"]
	if param["userid"] != "1" {
		sql += ` and f.user_id= '` + sb + `'`
	}
	//状态
	status := param["status"]
	if param["status"] != "" {
		sql += ` and f.status  = '` + status + `'`
	}
	//业务员
	nickname := param["nick_name"]
	if param["nick_name"] != "" {
		sql += ` and f.nickname  =  '` + nickname + `'`
	}

	keywords := param["keyword"]
	if param["keyword"] != "" {
		sql += ` and (  name like '%` + keywords + `%' or phone like '%` + keywords + `%')`
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Investlist, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Investlist) Search(param map[string]string) (result interface{}, err error) {
	status := param["status"]
	fifa := param["days"]
	//sql:=`SELECT *FROM investment WHERE status= '`+status+`' and  DATEDIFF(expiration_date,now()) >=0 and DATEDIFF(expiration_date,now())< `+fifa+` `
	sql := `SELECT *from (SELECT i.id,i.amount,sum(amount*profit)profits,i.customerid,i.status,i.invest_time investtime,i.expiration_date expirationdate,i.create_time as ordertime,i.remark,i.profit,c.name,c.phone,c.identity id_car,
			c.bank banknum ,s.nick_name nickname,s.team_name teamname FROM investment i LEFT JOIN customer c  on i.customerid=c.id 
			LEFT JOIN sys_user s on i.userid=s.user_id WHERE i.status= 0 GROUP BY i.id)d WHERE d.status= ` + status + ` and  DATEDIFF(expirationdate,now()) >=0 and DATEDIFF(expirationdate,now())< ` + fifa + ``
	//if param["status"] != ""{
	//	sql+= `and status=`+status
	//}
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "asc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Investlist, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Investment) InvestmentAdd(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	var c Customer
	c, err = c.NewCustomer(param["customerid"])
	if err != nil {
		return err
	}
	if c.Status != 0 {
		return errors.New("该客户审核未通过！请审核通过再来")
	}
	param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	param["userid"] = c.Userid
	//sql := ` insert into investment(id,userid,amount,customerid,remark,status,expiration_date,profit)value(:id,:userid,:amount,:customerid,:remark,1,DATE_FORMAT(date_add(now(),interval 1 year),'%Y-%m-%d 00:00:00'),:profit)`
	sql := ` insert into investment(id,userid,amount,customerid,remark,status,invest_time,monthly_time,expiration_date,profit,cus_name)value(:id,:userid,:amount,:customerid,:remark,1,:date,date_add(:date,interval 1 month),DATE_FORMAT(date_add(:date,interval 1 year),'%Y-%m-%d 00:00:00'),:profit,:name)`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}
	//var ul UserLevel
	//ul.GetSetUserLevel()

	if err = NewCustomer1(param); err != nil {
		return err
	}

	sql1 := `insert into customer_log(customerid,salesmanid,name,identity,sex,phone,bank,banknum,investmentid,amount,iremark,edittype)
									value(:customerid,:userid,:name,:identity,:sex,:phone,:bank,:banknum,:id,:amount,:remark,2)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}

func (e *Investment) NewInvestmentAdd(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	var c Customer
	c, err = c.NewCustomer(param["customerid"])
	if err != nil {
		return err
	}
	if c.Status != 0 {
		return errors.New("该客户审核未通过！请审核通过再来")
	}
	param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	param["userid"] = c.Userid
	//sql := ` insert into investment(id,userid,amount,customerid,remark,status,expiration_date,profit)value(:id,:userid,:amount,:customerid,:remark,1,DATE_FORMAT(date_add(now(),interval 1 year),'%Y-%m-%d 00:00:00'),:profit)`
	sql := ` insert into investment(id,userid,amount,customerid,remark,status,invest_time,monthly_time,expiration_date,profit,cus_name)value(:id,:userid,:amount,:customerid,:remark,1,:date,date_add(:date,interval 1 month),DATE_FORMAT(date_add(:date,interval 1 year),'%Y-%m-%d 00:00:00'),:profit,:name)`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}
	//var ul UserLevel
	//ul.GetSetUserLevel()

	if err = NewCustomer1(param); err != nil {
		return err
	}

	sql1 := `insert into customer_log(customerid,salesmanid,name,identity,sex,phone,bank,banknum,investmentid,amount,iremark,edittype)
									value(:customerid,:userid,:name,:identity,:sex,:phone,:bank,:banknum,:id,:amount,:remark,2)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}

func (e *Investment) InvestmentEdit(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	var in []Investment
	sql2 := `select *,TIMESTAMPDIFF(day,invest_time,now())>=3 alterbool from investment where id = :investmentid`
	sql2 = utils.SqlReplaceParames(sql2, param)
	if err := orm.Eloquent.Raw(sql2).Scan(&in).Error; err != nil {
		return err
	}
	if len(in) == 0 {
		return errors.New("找不到该信息")
	}
	if in[0].Alterbool != 0 {
		return errors.New("三天已过，无法修改")
	}
	param["userid"] = in[0].Userid
	sql := ` update investment set status = 1   where id = :investmentid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}

	if err = NewInvestment(param); err != nil {
		return err
	}
	param["edittype"] = "3"
	if in[0].Status == 3 {
		param["edittype"] = "5"
	}
	//修改时间
	sql1 := `insert into customer_log(customerid,salesmanid,name,identity,phone,bank,banknum,sex,investmentid,amount,amountnew,iremark,iremarknew,edittype)
							    value(:customerid,:userid,:name,:identity,:phone,:bank,:banknum,:sex,:investmentid,:amount,:amountnew,:remark,:iremarknew,:edittype)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
}

func (e *Investment) InvestmentBreak(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	//22.9.05修改可终止的时间段
	var ins []Investment
	sql12 := `select *,TIMESTAMPDIFF(day,invest_time,now())<=1 alterbool from investment where id = :investmentid`
	sql12 = utils.SqlReplaceParames(sql12, param)
	if err := orm.Eloquent.Raw(sql12).Scan(&ins).Error; err != nil {
		return err
	}
	if len(ins) == 0 {
		return errors.New("找不到该信息")

	}
	if ins[0].Alterbool == 1 {
		return errors.New("小于一天，无法终止")

	}
	param["userid"] = ins[0].Userid

	sql := ` update investment set status =if(status=0,4,status) where id = :investmentid`
	sql = utils.SqlReplaceParames(sql, param)
	if 0 == orm1.Exec(sql).RowsAffected {
		return errors.New("状态异常，无法终止")
	}

	if err = NewInvestment(param); err != nil {
		return err
	}

	sql1 := `insert into customer_log(customerid,salesmanid,name,identity,phone,bank,banknum,sex,investmentid,amount,iremark,edittype,sexnew)
							    value(:customerid,:userid,:name,:identity,:phone,:bank,:banknum,:sex,:investmentid,:amount,:remark,4,:breaktype)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
}

func (e *Investment) InvestmentRevoke(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	sql := ` DELETE FROM investment WHERE id = :investmentid`
	sql = utils.SqlReplaceParames(sql, param)
	//原来判断
	err = orm1.Exec(sql).Error

	if err = NewInvestment(param); err != nil {
		return err
	}

	//上级所有人都减业绩
	in := GetInvestmentById(param["investmentid"])
	var ref Referrer
	ref = ref.GetReferrer(in.Userid)
	ref.Referrers += "," + in.Userid
	ref.Referrers = strings.Trim(ref.Referrers, ",")
	// if strings.Trim(ref.Referrers, " ") != "" {
	sql2 := `update sys_user set accumulative = accumulative + real_value -` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + ref.Referrers + `) and level_lock=0`
	if err = orm1.Exec(sql2).Error; err != nil {
		return
	}
	sql4 := `update sys_user set real_value = real_value -` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) and level_lock=1`
	if err = orm1.Exec(sql4).Error; err != nil {
		return
	}
	// }

	// sql3 := `update sys_user set accumulative = accumulative + real_value -` + utils.Float64ToString(in.Amount) + `,real_value=0 where user_id in (` + in.Userid + `) `
	// if err = orm1.Exec(sql3).Error; err != nil {
	// 	return
	// }

	sql1 := ` DELETE FROM customer_log WHERE investmentid = :investmentid`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error

	sql11 := ` DELETE FROM investmentprofit WHERE investmentid = :investmentid`
	sql11 = utils.SqlReplaceParames(sql11, param)
	err = orm1.Exec(sql11).Error
	return
}

func (e *Investment) InvestmentById(id string) (res interface{}, err error) {
	sql := `select * from investment where id = ?`
	var in Investment
	err = orm.Eloquent.Raw(sql, id).Scan(&in).Error
	return in, err
}

func (e *Investment) ExportInvestmentList(param map[string]string) (interface{}, error) {
	sql := `select id,userid,amount,customerid,remark,is_del,status
			,DATE_FORMAT(create_time,'%Y-%m-%d %H:%i:%s') create_time
			,DATE_FORMAT(update_time,'%Y-%m-%d %H:%i:%s') update_time
			,DATE_FORMAT(invest_time,'%Y-%m-%d %H:%i:%s') invest_time
			,DATE_FORMAT(expiration_date,'%Y-%m-%d %H:%i:%s') expiration_date
			,DATE_FORMAT(monthly_time,'%Y-%m-%d %H:%i:%s') monthly_time,profit,oncestatus from investment order by invest_time asc `
	var ie []InvestmentExport
	err := orm.Eloquent.Raw(sql).Scan(&ie).Error

	return ie, err
}

// 导出一次性分润报表
func (e *Investment) ExportInvestment(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "id,amount,customerid,create_time,update_time,is_del,remark,status,invest_time,expiration_date,userid,monthly_time,profit,oncestatus"
	//param["filename"] = "ID,金额,客户ID,创建时间,修改时间,是否已删除,备注,状态,最初创建时间,合同到期时间,业务员ID,客户分润时间,客户利润点,一次性分润状态"
	param["filename"] = "id,amount,customerid,create_time,update_time,is_del,remark,status,invest_time,expiration_date,userid,monthly_time,profit,oncestatus"
	param["title"] = "投资列表导出" + param["date1"]
	URL, err = GetExcelURL(e.ExportInvestmentList, param)
	return
}

// 删除订单
func (e *Investment) DeleteInvestment() (err error) {
	//先备份
	sees := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			sees.Rollback()
			return
		}
		sees.Commit()
	}()
	//sql := `select * from investment `
	//var ie []InvestmentExport
	//if err = orm.Eloquent.Raw(sql).Scan(&ie).Error;err!=nil{
	//	return err
	//}
	inv, err := e.ExportInvestmentList(nil)
	if err != nil {
		return nil
	}
	ie := inv.([]InvestmentExport)
	//获取备份批次
	sql0 := `select max(batch)+1 batch from investment_backup `
	type Batchs struct {
		Batch string
	}
	var b Batchs
	if err = orm.Eloquent.Raw(sql0).Scan(&b).Error; err != nil {
		return err
	}
	if b.Batch == "" {
		b.Batch = "0"
	}
	sql1 := `insert into investment_backup(investment_id,amount,customerid,create_time,update_time,is_del,remark,status
						,invest_time,expiration_date,userid,monthly_time,profit,oncestatus,batch)values`
	for k, v := range ie {
		sql1 += `('` + v.ID + `','` + v.Amount + `','` + v.Customerid + `','` + v.CreateTime + `','` + v.UpdateTime + `','` + v.IsDel + `','` + v.Remark + `','` + v.Status + `','` + v.InvestTime + `
					','` + v.ExpirationDate + `','` + v.Userid + `','` + v.MonthlyTime + `','` + v.Profit + `','` + v.Oncestatus + `','` + b.Batch + `')`
		if k < len(ie)-1 {
			sql1 += ","
		}
	}
	if len(ie) > 0 {
		if err = sees.Exec(sql1).Error; err != nil {
			return err
		}
	}
	//备份分润信息
	sql2 := `select * from investmentprofit  `
	var ipe []InvestmentprofitExport
	if err = orm.Eloquent.Raw(sql2).Scan(&ipe).Error; err != nil {
		return err
	}
	sql3 := `insert into investmentprofit_backup(investmentprofit_id,investmentid,userid,profits,batch)values`
	for k, v := range ipe {
		sql3 += `('` + v.ID + `','` + v.Investmentid + `','` + v.Userid + `','` + v.Profits + `','` + b.Batch + `')`
		if k < len(ipe)-1 {
			sql3 += ","
		}
	}
	if len(ie) > 0 {
		if err = sees.Exec(sql3).Error; err != nil {
			return err
		}
	}

	sql4 := `delete from investment `
	if err = sees.Exec(sql4).Error; err != nil {
		return err
	}

	sql5 := `delete from investmentprofit `
	if err = sees.Exec(sql5).Error; err != nil {
		return err
	}
	sql6 := `update sys_user set accumulative = 0 `
	if err = sees.Exec(sql6).Error; err != nil {
		return err
	}

	return
}

// 导入订单
func (e *Investment) InvestmentImport(file multipart.File, Size int64) error {
	buf := make([]byte, Size)
	n, _ := file.Read(buf)
	//sees := orm.Eloquent.Begin()
	//var err error
	//defer func() {
	//	if err!=nil{
	//		sees.Rollback()
	//		return
	//	}
	//	sees.Commit()
	//}()

	xf, _ := xlsx.OpenBinary(buf[:n])
	for _, sheet := range xf.Sheets {
		if len(sheet.Rows) == 0 {
			continue
		}

		//sql1 = `insert into investment(id,amount,customerid,create_time,update_time,is_del,remark,status
		//				,invest_time,expiration_date,userid,monthly_time,profit,oncestatus)values`
		//记录字段位置
		filedvalue := make(map[string]int)
		sql1 := `insert into investment(`
		for j, row := range sheet.Rows {
			//保存参数
			sql2 := ""
			param := make(map[string]string)
			for i, cell := range row.Cells {
				if j == 0 {
					sql1 += cell.String()
					if i < len(row.Cells)-1 {
						sql1 += ","
					} else {
						sql1 += ")values"
					}
					filedvalue[cell.String()] = i
				} else {
					if i == 0 {
						//if j > 1{
						//	sql1 += ","
						//}
						sql2 += "("
					}
					sql2 += "'" + cell.String() + "'"
					if i < len(row.Cells)-1 {
						sql2 += ","
					} else {
						sql2 += ")"
					}
					//保存字段
					switch i {
					case filedvalue["amount"]:
						param["amount"] = cell.String()
					case filedvalue["userid"]:
						param["userid"] = cell.String()
					case filedvalue["id"]:
						param["investmentid"] = cell.String()
					case filedvalue["status"]:
						param["status"] = cell.String()
					}
				}

			}
			if j > 0 {
				var i InvestmentShareProfit
				//err = i.AddInvestmentShareProfitImport(sees,param)
				if param["status"] != "1" {
					err := i.AddInvestmentShareProfitImport(param)
					if err != nil {
						return err
					}
					if err = orm.Eloquent.Exec(sql1 + sql2).Error; err != nil {
						return err
					}
				}

				time.Sleep(time.Millisecond * 100)

			}

		}
		//if err = sees.Exec(sql1).Error;err!=nil{
		//if err = orm.Eloquent.Exec(sql1).Error;err!=nil{
		//	return err
		//}
	}
	return nil
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func (e *Investment) GetAccumulative(id string) (string, error) {
	var ref Referrer
	ref = ref.GetReferrer(id)
	als := strings.Trim(ref.Referrals, ",")
	sql := `select sum(amount)accumulative from investment where userid in(` + als + `) `
	type Acc struct {
		Accumulative string
	}
	var acc Acc
	if err := orm.Eloquent.Raw(sql).Scan(&acc).Error; err != nil {
		return "", nil
	}
	return acc.Accumulative, nil
}

func (e *Investlist) InvestExpirationLists() (result []Investlist, err error) {
	list := make([]Investlist, 0)
	// 查询到期时间小于当前时间的投资列表
	sql := `SELECT * FROM investment WHERE status= '0' and  DATEDIFF(expiration_date,now()) <=0`
	orm.Eloquent.Raw(sql).Scan(&list)
	result = list
	return
}
