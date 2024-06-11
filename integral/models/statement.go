package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Statement struct {
	Percent       float64 `gorm:"column:percent" json:"percent"`
	Profits       float64 `gorm:"column:profits" json:"profits"`
	Amount        float64 `gorm:"column:amount" json:"amount"`
	Userid        string  `gorm:"column:userid" json:"userid"`
	Username      string  `gorm:"column:username" json:"username"`
	NickName      string  `gorm:"column:nick_name" json:"nick_name"`
	Name          string  `gorm:"column:customername" json:"customername"`
	Investmentid  string  `gorm:"column:investmentid" json:"investmentid"`
	TeamName      string  `gorm:"column:teamname" json:"teamname"`
	TotalInvest   string  `gorm:"column:totalinvest" json:"totalinvest"`
	TotalCustomer string  `gorm:"column:totalcustomer" json:"totalcustomer"`
	//Referrals	string 	`gorm:"column:referrals" json:"referrals"`
	ReftTotal    string    `gorm:"column:reftotal" json:"reftotal"`
	Investtol    string    `gorm:"column:investtol" json:"invest_tol"`
	IsDel        string    `gorm:"column:is_del" json:"isdel"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createtime"`
	InestTime    time.Time `gorm:"column:invest_time" json:"invest_time"`
	Scount       string    `gorm:"column:scount" json:"scount"`
	Status       string    `gorm:"column:status" json:"status"`
	Cuscount     string    `gorm:"column:cuscount" json:"cuscount"`
	CountProfits string    `gorm:"column:countprofits" json:"countprofits"`
	AccumulaTive string    `gorm:"column:accumulative" json:"accumulative"`
}
type Total struct {
	Profits float64 `json:"profits"`
	Total   string  `json:"total"`
	Amount  float64 `json:"amount"`
}
type TotalList struct {
	Profits float64 `json:"profits"`
	Userid  string  `json:"userid"`
	Amount  float64 `json:"amount"`
}

// 业务员报表
func (e *Statement) StatementSalesmanNew(param map[string]string) (result interface{}, total interface{}, err error) {
	//获取投资列表
	//sql := `select c.userid,u.nick_name,c.name customername,u.username,i.amount,ip.profits,ip.investmentid from investment i
	//	left join investmentprofit ip on i.id = ip.investmentid
	//	left join customer c on i.customerid = c.id
	//	left join sys_user u on ip.userid = u.user_id
	//	where i.status in (0,4,5) and i.is_del = 0 and i.create_time <= :end and
	//										:start <= i.create_time `
	//	//where i.status in (0,4,5) and i.is_del = 0 and i.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//	//									date_add(:date,interval -day(:date)+1 day) <= i.create_time `
	//sql:=`SELECT a.userid,a.nick_name,a.username,a.teamname,(a.accumulative+a.real_value)accumulative,a.reftotal,a.investtol,a.is_del,a.create_time,ifnull(b.scount,0) scount,cuscount,countprofits from(
	//					SELECT  u.user_id userid,u.nick_name,u.username,u.team_name teamname,u.accumulative,u.real_value,CHAR_LENGTH(r.referrals)-CHAR_LENGTH(REPLACE(r.referrals,',','')) as reftotal,
	//						count(i.id)investtol, i.is_del,i.create_time   FROM
	//						sys_user u
	//						LEFT JOIN referrer r on u.user_id= r.userid
	//						LEFT JOIN investment i on  u.user_id=i.userid
	//					  WHERE u.status in (0,4,5,6) GROUP BY r.userid
	//
	//						)a
	//						left join (
	//							select count(1) scount,referrer user_id from sys_user GROUP BY ifnull( referrer,0)
	//						) b on a.userid = b.user_id LEFT JOIN (
	//						  SELECT count(1)cuscount ,userid FROM customer GROUP BY userid
	//						)c on a.userid = c.userid LEFT JOIN (
	//							SELECT sum(profits) countprofits,userid FROM investmentprofit GROUP BY userid
	//						)d on a.userid =d.userid
	//						WHERE  a.is_del = 0 and a.create_time <= :end and  :start <= a.create_time `
	//sql := `SELECT *FROM(SELECT b.userid,a.nick_name,a.username,a.teamname,k.amount amount,a.reftotal,a.investtol,a.is_del,ifnull(d.scount,0) scount,cuscount,b.profits,f.create_time from(
	//					SELECT  u.user_id userid,u.nick_name,u.username,u.team_name teamname,u.accumulative,u.real_value,u.status, CHAR_LENGTH(r.referrals)-CHAR_LENGTH(REPLACE(r.referrals,',','')) as reftotal,
	//						count(i.id)investtol, i.is_del  FROM
	//						sys_user u
	//						LEFT JOIN referrer r on u.user_id= r.userid
	//						LEFT JOIN investment i on  u.user_id=i.userid
	//					  WHERE i.status in (0) GROUP BY r.userid
	//
	//						)a
	//						left join (
	//						SELECT sum(profits) profits,userid,investmentid FROM investmentprofit GROUP BY userid
	//						) b on a.userid = b.userid
	//						LEFT JOIN (SELECT userid,profits,investmentid from investmentprofit GROUP BY investmentid,userid)l on a.userid=l.userid
	//						LEFT JOIN (SELECT id,userid,create_time,customerid,status from investment WHERE status=0)f on l.investmentid=f.id
	//						LEFT JOIN (SELECT sum(amount)amount,userid FROM investment GROUP BY userid)k on  a.userid=k.userid
	//						LEFT JOIN (SELECT userid,id,create_time FROM customer WHERE status=0)j on f.customerid=j.id
	//						LEFT JOIN (
	//						  SELECT count(1)cuscount ,userid FROM customer GROUP BY userid
	//						)c on a.userid = c.userid LEFT JOIN (
	//							select count(1) scount,referrer user_id from sys_user GROUP BY ifnull( referrer,0)
	//						)d on b.userid =d.user_id
	//
	//						WHERE  a.status=0  and f.create_time <= :end and   :start <= f.create_time GROUP BY b.userid )g`
	sql := `select ip.userid,u.nick_name,c.name customername,u.username,round(i.amount,2)amount,round(ip.profits,2)profits,ip.investmentid,u.team_name teamname from investment i
		left join investmentprofit ip on i.id = ip.investmentid
		left join customer c on i.customerid = c.id
		left join sys_user u on ip.userid = u.user_id
		where i.status in (0,4,5,6) and ifnull(i.manually_end,0) = 0 and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time `

	keyword := param["keyword"]
	if param["keyword"] != "" {
		sql += ` and  nick_name like '%` + keyword + `%' or username like '%` + keyword + `%'`
	}
	sql = utils.SqlReplaceParames(sql, param)
	var s []Statement
	var t Total
	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	//因为关联了子表，sum(amount)值不对
	var t1 Total
	sqlt1 := `select round(sum(i.amount),2)amount from investment i
		where i.status in (0,4,5,6) and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time `
	sqlt1 = utils.SqlReplaceParames(sqlt1, param)
	if err = orm.Eloquent.Raw(sqlt1).Scan(&t1).Error; err != nil {
		return nil, nil, err
	}

	t.Amount = t1.Amount

	sqltotal := `select round(sum(i.amount),2) amount,round(sum(ip.profits),2) profits,ip.userid from investment i
		left join investmentprofit ip on i.id = ip.investmentid
		left join customer c on i.customerid = c.id
		left join sys_user u on ip.userid = u.user_id
		where i.status in (0,4,5,6) and i.is_del = 0 and i.create_time <= :end and
											:start <= i.create_time
		GROUP BY ip.userid`
	var tlist []TotalList
	sqltotal = utils.SqlReplaceParames(sqltotal, param)
	if err = orm.Eloquent.Raw(sqltotal).Scan(&tlist).Error; err != nil {
		return
	}
	param["total"] = t.Total
	param["sort"] = "ip.userid ,i.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	if err = orm.Eloquent.Raw(sql).Scan(&s).Error; err != nil {
		return nil, nil, err
	}
	var pre Statement
	if len(s) > 0 {
		pre = s[0]
	}

	var st []Statement

	for i, v := range s {
		if v.Userid == pre.Userid {
			pre = v
			st = append(st, v)
			if i == len(s)-1 {
				for _, val := range tlist {
					if val.Userid == v.Userid {
						var total Statement
						total.NickName = "合计"
						total.Profits = val.Profits
						total.Amount = val.Amount
						st = append(st, total)
					}
				}
			}
			continue
		}
		for _, val := range tlist {
			if val.Userid == pre.Userid {
				//fmt.Println(val)
				var total Statement
				total.NickName = "合计"
				total.Profits = val.Profits
				total.Amount = val.Amount
				st = append(st, total)
			}
		}
		st = append(st, v)
		pre = v
	}

	return st, t, err
}

func (e *Statement) Summrydetail(param map[string]string) (result interface{}, total interface{}, err error) {
	aa := param["userid"]
	//fmt.Println("id",aa)
	//fmt.Println("idd",param["userid"])
	//拼凑筛选条件sql
	sql := ` SELECT *FROM(SELECT b.*,d.customername,d.amount,d.invest_time,d.status,d.create_time FROM(SELECT 	s.user_id userid,v.investmentid,round(v.profits,2)profits,s.nick_name FROM investmentprofit v LEFT JOIN
			sys_user s on v.userid=s.user_id WHERE s.status=0 and s.user_id= ` + aa + `  )b LEFT JOIN (SELECT a.amount,a.customerid,a.userid,s.name customername,a.id,a.invest_time ,
			a.status,a.create_time FROM investment a left JOIN  customer s on a.customerid=s.id)d on b.investmentid=d.id   GROUP BY b.investmentid)d   WHERE d.create_time <= :end and :start <= d.create_time `

	sql = utils.SqlReplaceParames(sql, param)
	//var s []Statement
	var t Total
	if err = orm.Eloquent.Raw(`select sum(a.profits)profits,sum(a.amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "d.investmentid"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Statement, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return result, t, err
}

func (e *Statement) StatementSalesman(param map[string]string) (result interface{}, total interface{}, err error) {
	//sql := `select *,amount*percent profits from (
	//		select round(sum(percent),4)percent,amount,a.userid,u.username,u.nick_name,u.nick_name name from (
	//				select c.amount,ifnull(p.userid,0)pid,c.userid from (
	//						select sum(i1.amount)amount,c1.userid from customer c1
	//							left join investment i1 on i1.customerid = c1.id
	//							where c1.is_del = 0 and i1.status in (0,4,5) and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//										date_add(:date,interval -day(:date)+1 day) <= i1.create_time
	//							group by c1.userid
	//						)c
	//						left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid
	//		)a
	//		left join profitconfig p1 on a.pid = p1.userid
	//		left join sys_user u on a.userid = u.user_id
	//		where p1.profitlevel < a.amount and p1.profittype = if(ifnull(p1.userid,0)=0,2,1)
	//		GROUP BY a.userid
	//		)b `
	sql := `select sum(amount) amount,sum(profits)profits,a.userid,u.username,u.nick_name,u.nick_name name from (
				select amount,a.userid,a.id,sum(p1.levelgain)+(amount-max(p1.profitlevel))*sum(percent) profits from (
					select c.amount,ifnull(p.userid,0)pid,c.userid,c.id from ( 
							select i1.amount,c1.userid,i1.id from customer c1
								left join investment i1 on i1.customerid = c1.id
								where c1.is_del = 0 and i1.status in (0,4,5) and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
											date_add(:date,interval -day(:date)+1 day) <= i1.create_time
							)c 
							left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid

				)a 
				left join profitconfig p1 on a.pid = p1.userid
				where p1.profitlevel < a.amount and p1.profittype = if(ifnull(p1.userid,0)=0,2,1) 
				GROUP BY a.id
			)a
			left join sys_user u on a.userid = u.user_id
			GROUP BY a.userid `
	now := time.Now()
	nowdate := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	if param["date"] != nowdate {
		sql = `select *,name nick_name from statement where DATE_FORMAT(create_time,'%Y-%m-01') =:date and stype = 1 `
	}
	sql = utils.SqlReplaceParames(sql, param)
	var s []Statement
	var t Total

	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "amount"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	err = orm.Eloquent.Raw(sql).Scan(&s).Error
	return s, t, err
}

type UserInvestment struct {
	UserLevel
	Investment
	Referrer
	//ShareProfit []InvestmentShareProfit
}

// 顾客报表
func (e *Statement) StatementCustomerNew(param map[string]string) (result interface{}, total interface{}, err error) {
	//原来的数据库语句
	//sql := `select i.profit,
	//		if(i.expiration_date>i.monthly_time,
	//					ROUND(i.profit*i.amount,2),
	//					ROUND(TIMESTAMPDIFF(day,date_sub(date(monthly_time) ,interval 1 month),expiration_date)/day(LAST_DAY(date_sub(monthly_time, interval 1 month)))*i.profit*i.amount,2)) profits,
	//		i.amount,c.id, c.name, c.phone, c.bank, c.banknum, c.userid, c.is_del, c.status, c.create_time, c.sex, c.identity,
	//		u.nick_name, i.expiration_date,i.id investment_id, i.monthly_time
	//		from investment i
	//		LEFT JOIN customer c on i.customerid = c.id
	//		left join sys_user u on c.userid = u.user_id
	//		WHERE i.is_del = 0 and i.status in (0,4,5) and and day(i.update_time) <= day(i.expiration_date) `//原来的sql语句
	//
	//if param["day"]!=""{
	//	sql += ` and DATE_FORMAT(monthly_time,'%e') = :day `
	//}
	fmt.Println("shuhsu", param["start"])
	sql := `select i.profit, 
			if(i.expiration_date>i.monthly_time,
						ROUND(i.profit*i.amount,2),
						ROUND(TIMESTAMPDIFF(day,date_sub(date(monthly_time) ,interval 1 month),expiration_date)/day(LAST_DAY(date_sub(monthly_time, interval 1 month)))*i.profit*i.amount,2)) profits,
			i.amount,c.id, c.name, c.phone, c.bank, c.banknum, c.userid, c.is_del, c.status, i.create_time, i.create_time ordertime,c.sex, c.identity, 
			u.nick_name, i.expiration_date,i.id investment_id, i.monthly_time,i.invest_time,u.team_name teamname,u.team_id
			from investment i
			LEFT JOIN customer c on i.customerid = c.id
			left join sys_user u on c.userid = u.user_id
			WHERE i.is_del = 0 and i.status in (0,4,5)  and i.create_time <= :end and
					:start <= i.create_time `
	//and i.status in (0,4,5) 后的and  day(i.update_time) <= day(i.expiration_date) 去掉
	//搜索
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and ( c.name like '%` + keyword + `%' or nick_name like '%` + keyword + `%' or c.phone like '%` + keyword + `%') `
	}
	shutup := param["investment_id"]
	if shutup != "" {
		sql += ` and ( i.id like '%` + shutup + `%')`
	}
	del := param["deptid"]
	if param["deptid"] != "" {
		sql += `and u.team_id = ` + del
	}
	//原来的语句
	var s []Customer
	sql = utils.SqlReplaceParames(sql, param)

	var t Total
	if err = orm.Eloquent.Raw(`select sum(profits)profits,sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "monthly_time" //原来的语句
	param["order"] = "asc"
	sql += utils.LimitAndOrderBy(param)
	err = orm.Eloquent.Raw(sql).Scan(&s).Error
	return s, t, err
}

// 判断上月报表是否已生成
func (e *Statement) GetStatementBool(date string, stype int) (result bool) {
	sql := `select 1 from statement s where s.stype = ` + strconv.Itoa(stype) + ` and date_format(create_time,'%Y-%m') = date_format('` + date + `','%Y-%m')  `
	return GetTotalCount1(sql) > 0
}
func (e *Statement) GetSummaryBool(date string) (result bool) {
	sql := `select 1 from summary s where date_format(create_time,'%Y-%m') = date_format('` + date + `','%Y-%m')  `
	return GetTotalCount1(sql) > 0
}

// 导出业务员报表
func (e *Statement) ExportExcelSalesman(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "nick_name,investmentid,customername,teamname,amount,profits"
	param["filename"] = "业务员姓名,订单号,客户姓名,团队类型,业绩,利润"
	param["title"] = "业务员报表" + param["date1"]
	URL, err = GetExcelTotal1(e.StatementSalesmanNew, param)
	return
}

// 导出客户报表
func (e *Statement) ExportExcelCustomer(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	//param["filefield"] = "name,phone,banknum,bank,sex,profit,amount,profits,nick_name,create_time,update_time,expiration_date"
	param["filefield"] = "investment_id,create_time,name,phone,banknum,bank,sex,profit,amount,profits,nick_name,teamname,create_time,update_time,expiration_date"
	param["filename"] = "投资订单号,订单创建时间,姓名,手机号,银行卡,开户行,性别,分润比例,业绩,利润,业务员,团队类型,投资时间,结算时间,合同到期时间"
	param["title"] = "客户报表" + param["date1"]
	URL, err = GetExcelTotal(e.StatementCustomerNew, param)
	return
}

func (e *Statement) CustomerSettle(investmentID string) (err error) {
	sql1 := `select c.name name, c.bank bank, c.banknum banknum, c.userid user_id, u.nick_name nick_name, c.id customer_id
		from investment i, customer c, sys_user u
		where i.id=` + investmentID + ` and c.id = i.customerid and c.userid = u.user_id`

	find := make([]SettleCustomerHistory, 0)
	if err = orm.Eloquent.Raw(sql1).Scan(&find).Error; err != nil {
		return
	}
	if len(find) != 1 {
		err = errors.New("查找数据错误")
		return
	}

	findMsg := find[0]
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	sql2 := `update investment SET monthly_time=if(expiration_date<monthly_time,monthly_time,adddate(monthly_time, interval 1 month)),status = if(status=5||monthly_time>expiration_date,6,status), update_time=now() WHERE id = ` + investmentID
	if err = orm1.Exec(sql2).Error; err != nil {
		return
	}

	sql3 := `insert into customer_settle (name, bank, banknum, user_id, nick_name, settle_time, invest_id) values ("` +
		findMsg.Name + `","` + findMsg.Bank + `","` + findMsg.BankNum + `",` + strconv.FormatInt(findMsg.UserID, 10) + `,"` +
		findMsg.NickName +
		`", now(), "` + investmentID +
		`")`
	if err = orm1.Exec(sql3).Error; err != nil {
		return
	}

	return
}
