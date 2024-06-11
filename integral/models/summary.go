package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"xAdmin/config"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Summary struct {
	Id             string    `gorm:"column:id" json:"id"`
	Company        float64   `gorm:"column:company" json:"company"`
	Salesman       float64   `gorm:"column:salesman" json:"salesman"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	Remark         string    `gorm:"column:remark" json:"remark"`
	NickName       string    `gorm:"column:nick_name" json:"name"`
	Customerprofit float64   `gorm:"column:customerprofit" json:"customerprofit"`
	Investmentid   string    `gorm:"column:investmentid" json:"investmentid"`
	Customername   string    `gorm:"column:customername" json:"customername"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	Amount         float64   `gorm:"column:amount" json:"amount"`
	TeamName       string    `gorm:"column:team_name" json:"teamname"`
	TeamId         int64     `gorm:"column:team_id" json:"teamid"`
}
type SummaryTotal struct {
	Company        float64 `gorm:"column:company" json:"company"`
	Salesman       float64 `gorm:"column:salesman" json:"salesman"`
	Customerprofit float64 `gorm:"column:customerprofit" json:"customerprofit"`
	Amount         float64 `gorm:"column:amount" json:"amount"`
	Total          string  `gorm:"column:total" json:"total"`
}

// 一次性分配报表
func (e *Summary) StatementConfigOnce(param map[string]string) (result interface{}, total interface{}, err error) {
	//sql := `select userid,username,nick_name,nick_name name,percent,round(amount*percent,2) profits,amount,a.team_name teamname  from (
	//	select p.*,u.nick_name,u.phone,u.username,u.user_id,a.amount from profitconfig p
	//	left join sys_user u on p.userid = u.user_id
	//	left join (select sum(amount)amount FROM investment
	//										WHERE is_del = 0 and ifnull(manually_end,0) = 0 and status in (0,4,5) and create_time <= :end and :start <= create_time ) a on 1 = 1
	//	where p.profittype = 0
	//	)a`
	//新修改过后的sql语句
	sql := `select userid,username,nick_name,a.team_name teamname ,a.team_id,percent,round(amount*percent,2) profits,amount  from (
                select p.*,u.nick_name,u.phone,u.username,u.user_id,a.amount,a.team_id from profitconfig p
                left join sys_user u on p.userid = u.user_id
                left join (select team_name,s.team_id,sum(amount)amount FROM investment i  LEFT JOIN sys_user s on i.userid=s.user_id
                                                                                        WHERE i.is_del = 0 and ifnull(manually_end,0) = 0 and i.status in (0,4,5) and i.create_time <= :end and :start <= i.create_time GROUP BY s.team_name ) a on 1 = 1 and p.team_name=a.team_name
                where p.profittype = 0
                )a`
	val := param["deptid"]
	if val != "" {
		sql += ` WHERE team_id =` + val + `` //新加条件查询团队
	}
	sql = utils.SqlReplaceParames(sql, param)
	var s []Statement
	var t Total
	if err = orm.Eloquent.Raw(`select sum(profits)profits,max(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "amount"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	//sql+= `WHERE a.team_name=`+param["teamname"]+``//新加条件查询团队
	err = orm.Eloquent.Raw(sql).Scan(&s).Error
	return s, t, err
}

// 一次性分配报表（未结算）
func (e *Summary) StatementNoSettlement(param map[string]string) (interface{}, interface{}, error) {
	sql := `select i.id,i.amount,i.customerid,i.create_time,i.update_time,i.remark,c.userid,c.name customername,u.nick_name,u.team_id teamid,u.team_name teamname
			from investment i
			left join customer c on i.customerid = c.id
			left join sys_user u on u.user_id = c.userid
				where i.is_del = 0 and ifnull(i.manually_end,0) = 0 and oncestatus = 0 and i.status in (0,4,5,6) and i.create_time <= :end and
											:start <= i.create_time`
	//团队搜索
	keyword := param["keyword"]
	if param["keyword"] != "" {
		sql += ` and (c.name like '%` + keyword + `%' or u.nick_name like '%` + keyword + `%' )`
	}
	dels := param["id"]
	if param["id"] != "" {
		sql += ` and i.id =` + dels
	}
	del := param["deptid"]
	if param["deptid"] != "" {
		sql += ` and u.team_id=` + del
	}
	var in []Investment
	var t Total
	var err error
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Raw(`select sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	//param["amount"] = utils.Float64ToString(t.Amount)
	param["sort"] = "amount"
	param["order"] = "desc"
	//is:=param["teamid"]
	//is:=utils.Int64ToString(in[0].TeamId)
	//once := NewProfitconfig(0,is)
	sql += utils.LimitAndOrderBy(param)
	err = orm.Eloquent.Raw(sql).Scan(&in).Error
	//fmt.Println("shuju",in[0].TeamId)
	//is:=utils.Int64ToString(in[0].TeamId)
	once := NewProfitconfig(0)
	keys := ""
	words := ""
	res := make([]map[string]interface{}, 0)
	result := make(map[string]interface{})
	total := make(map[string]interface{})
	total["amount"] = t.Amount
	for i, v := range in {
		re := make(map[string]interface{})
		re["id"] = v.ID
		re["amount"] = v.Amount
		re["customerid"] = v.Customerid
		re["create_time"] = v.CreateTime
		re["update_time"] = v.UpdateTime
		re["remark"] = v.Remark
		//re["status"] = v.Status
		re["userid"] = v.Userid
		re["customername"] = v.Customername
		//re["customerprofit"] = v.Customerprofit
		re["nick_name"] = v.NickName
		re["teamid"] = v.TeamId
		re["teamname"] = v.TeamName
		//fmt.Println(v.Userid)
		for k, val := range once {
			//re[val.Cvalue] = val.Percent * v.Amount
			//if i == 0 {
			//	keys += val.Cname
			//	words += val.Cvalue
			//	if k < len(once)-1 {
			//		keys += ","
			//		words += ","
			//	}
			//	total[val.Cvalue] = t.Amount * val.Percent
			//}
			if i == 0 {
				keys += val.Cname
				words += val.Cvalue
				if k < len(once)-1 {
					keys += ","
					words += ","
				}
			}
			if v.TeamId != val.TeamId {
				continue
			}
			tal, ok := total[val.Cvalue]
			if !ok {
				tal = 0.0
			}

			total[val.Cvalue] = tal.(float64) + v.Amount*val.Percent
			re[val.Cvalue] = val.Percent * v.Amount
			fmt.Println("dusu", re[val.Cvalue])
		}

		res = append(res, re)
	}
	keys = strings.Trim(keys, ",")
	words = strings.Trim(words, ",")
	result["profit"] = res
	result["keys"] = keys
	result["words"] = words

	return result, total, err
}

// 一次性分配报表（已结算）
func (e *Summary) StatementIsSettlement(param map[string]string) (interface{}, interface{}, error) {
	//sql := `select * from summary where :start <= create_time and create_time <= :end  `
	sql := `select s.* ,u.team_name,u.team_id  from summary s LEFT JOIN sys_user u on s.nick_name=u.nick_name where :start <= s.create_time and s.create_time <= :end  `
	sas := param["investmentid"]
	if param["investmentid"] != "" {
		sql += `and investmentid = ` + sas + ` `
	}
	keyword := param["keyword"]
	if param["keyword"] != "" {
		sql += ` and customername like '%` + keyword + `%' or s.nick_name like '%` + keyword + `%'  `
	}
	sqs := param["deptid"]
	if param["deptid"] != "" {
		sql += `and u.team_id = ` + sqs + ` `
	}
	var in []Summary
	var t Total
	sql = utils.SqlReplaceParames(sql, param)
	var err error
	if err = orm.Eloquent.Raw(`select sum(amount)amount,count(1) total from(` + sql + `)a`).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	param["total"] = t.Total
	param["sort"] = "create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	if err = orm.Eloquent.Raw(sql).Scan(&in).Error; err != nil {
		return nil, nil, err
	}

	keys := ""
	words := ""
	res := make([]map[string]interface{}, 0)
	result := make(map[string]interface{})
	total := make(map[string]interface{})
	total["amount"] = t.Amount
	var de Summarydetails
	once, err := de.GetSummarydetails(param)
	if err != nil {
		return nil, nil, err
	}
	for i, v := range in {
		re := make(map[string]interface{})
		re["id"] = v.Id
		re["amount"] = v.Amount
		re["salesman"] = v.Salesman
		re["company"] = v.Company
		re["customerprofit"] = v.Customerprofit
		re["remark"] = v.Remark
		re["investmentid"] = v.Investmentid
		re["nick_name"] = v.NickName
		re["customername"] = v.Customername
		re["create_time"] = v.CreateTime
		re["update_time"] = v.UpdateTime
		re["teamname"] = v.TeamName
		re["teamid"] = v.TeamId
		for k, val := range once {
			if val.Summaryid == v.Id {
				re[val.Cvalue] = val.Profits
				if i == 0 {
					keys += val.Cname
					words += val.Cvalue
					if k < len(once)-1 {
						keys += ","
						words += ","
					}
					total[val.Cvalue] = t.Amount * val.Percent
				}
			}
		}
		res = append(res, re)
	}
	keys = strings.Trim(keys, ",")
	words = strings.Trim(words, ",")
	result["profit"] = res
	result["keys"] = keys
	result["words"] = words

	return result, total, err
}

// 添加一次性报表
func (e *Summary) AddStatementConfigOnce(id string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	summaryid := strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	sql := `insert into summary(id,company,salesman,create_time,remark,nick_name,investmentid,customername,update_time,amount)(
		select ` + summaryid + `,if(i.amount*` + utils.Float64ToString(config.ApplicationConfig.SalesmanRatio) + `-sum(ip.profits)<0,0
			,i.amount*` + utils.Float64ToString(config.ApplicationConfig.SalesmanRatio) + `-sum(ip.profits)) company
			,sum(ip.profits) salesman,i.create_time,i.remark,u.nick_name,
			i.id investmentid,c.name customername,now() update_time,i.amount from investment i 
			left join sys_user u on i.userid = u.user_id
			left join investmentprofit ip on ip.investmentid = i.id
			left join customer c on c.id = i.customerid
			where i.is_del = 0 and ifnull(i.manually_end,0)=0 and oncestatus = 0 and i.status in (0,4,5,6) and i.id = ` + id + `
			)`
	sql1 := `insert into summarydetails(summaryid,cname,cvalue,profits,percent,create_time)(
			select ` + summaryid + `,u.nick_name cname,u.username cvalue,i.amount * p.percent profits,percent,now() from investment i 
			left join profitconfig p on p.profittype = 0 
			left join sys_user u on p.userid = u.user_id
			where i.is_del = 0 and ifnull(i.manually_end,0)=0 and oncestatus = 0 and i.status in (0,4,5,6) and i.id = ` + id + `
		)`

	if err = orm1.Exec(sql).Error; err != nil {
		return
	}
	if err = orm1.Exec(sql1).Error; err != nil {
		return
	}
	sql2 := `update investment set oncestatus = 1 where id = '` + id + `'`
	if err = orm1.Exec(sql2).Error; err != nil {
		return
	}
	return
}

// 汇总报表
func (e *Summary) StatementSummaryNew(param map[string]string) (interface{}, interface{}, error) {
	//concat(LAST_DAY('2021-08-01'),' 23:59:59')
	sql := `select a.*,s.id summaryid from (
			select i.*,c.name customername,u.nick_name,sum(ip.profits) salesman,u.team_name teamname ,u.team_id teamid,c.phone phone
			from investment i
			left join customer c on i.customerid = c.id
			left join sys_user u on u.user_id = c.userid
			left join investmentprofit ip on i.id = ip.investmentid
				where i.is_del = 0 and ifnull(i.manually_end,0)=0 and i.status in (0,4,5,6) and i.invest_time <= :end and :start <= i.invest_time
				GROUP BY i.id )a 
			left join summary s on a.id = s.investmentid GROUP BY a.id HAVING a.id`
	//搜索
	keyword := param["keyword"]
	if param["keyword"] != "" {
		sql += ` and ( a.customername like '%` + keyword + `%' or a.nick_name like '%` + keyword + `%' or a.phone like '%` + keyword + `%') `
	}
	id := param["id"]
	if param["id"] != "" {
		sql += ` and id =  ` + id + ``
	}
	deptid := param["deptid"]
	if param["deptid"] != "" {
		sql += ` and a.teamid =  ` + deptid + ``
	}
	var im []InvestmentView
	sql = utils.SqlReplaceParames(sql, param)
	sqltotal := `select sum(amount)amount,count(1) total,sum(profits)profits from(select i.*,sum(ip.profits) profits
			from investment i
			left join customer c on i.customerid = c.id
			left join sys_user u on u.user_id = c.userid
			left join investmentprofit ip on i.id = ip.investmentid
				where i.is_del = 0 and ifnull(i.manually_end,0)=0 and i.status in (0,4,5,6) and i.invest_time <= :end and :start <= i.invest_time
				GROUP BY i.id)a
			 `
	var t Total
	sqltotal = utils.SqlReplaceParames(sqltotal, param)
	if err := orm.Eloquent.Raw(sqltotal).Scan(&t).Error; err != nil {
		return nil, nil, err
	}
	total := make(map[string]interface{})
	total["amount"] = t.Amount
	param["total"] = t.Total
	param["sort"] = "create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	err := orm.Eloquent.Raw(sql).Scan(&im).Error
	if err != nil {
		return nil, nil, err
	}
	once := NewProfitconfig(0)
	var sum Summarydetails
	su, err := sum.GetSummarydetails(param)
	if err != nil {
		return nil, nil, err
	}
	total["salesman"] = t.Profits
	//comt :=  config.ApplicationConfig.CustomerRatio * t.Amount -t.Profits
	comt := 0.2*t.Amount - t.Profits

	//fmt.Println("digi",config.ApplicationConfig.CustomerRatio)
	if comt < 0 {
		comt = 0
	}
	total["company"] = comt
	result := make(map[string]interface{})
	res := make([]map[string]interface{}, 0)
	keys := ""
	words := ""
	for i, v := range im {
		re := make(map[string]interface{})
		re["id"] = v.ID
		re["amount"] = v.Amount
		re["customerid"] = v.Customerid
		re["invest_time"] = v.InvestTime
		re["create_time"] = v.CreateTime
		re["update_time"] = v.UpdateTime
		re["remark"] = v.Remark
		re["status"] = v.Status
		re["userid"] = v.Userid
		re["customername"] = v.Customername
		re["nick_name"] = v.NickName
		re["salesman"] = v.Salesman
		re["teamname"] = v.TeamName
		re["phone"] = v.Phone
		re["teamid"] = v.TeamId
		//re["company"] = config.ApplicationConfig.CustomerRatio * v.Amount - v.Salesman
		re["company"] = 0.2*v.Amount - v.Salesman
		fmt.Println("ssa", v.Profit)
		fmt.Println("fugui", re["company"])
		//fmt.Println("sdsd",im[0].TeamId)
		for k, val := range once {
			if i == 0 {
				keys += val.Cname
				words += val.Cvalue
				if k < len(once)-1 {
					keys += ","
					words += ","

				}
			}
			if v.TeamId != val.TeamId {
				fmt.Println("v", v.TeamId)
				fmt.Println("sd", val.TeamId)
				continue
			}
			tal, ok := total[val.Cvalue]
			if !ok {
				tal = 0.0
			}

			total[val.Cvalue] = tal.(float64) + v.Amount*val.Percent
			re[val.Cvalue] = val.Percent * v.Amount
			fmt.Println("didi", re[val.Cvalue])
			//s:=strconv.Itoa(i)
			//if i==0 {
			//		keys += val.Cname
			//		fmt.Println("minc", keys)
			//		words += val.Cvalue
			//		if k < len(once)-1 {
			//			keys += ","
			//			words += ","
			//		}
			//		total[val.Cvalue] = t.Amount * val.Percent
			//		fmt.Println("sb",total[val.Cvalue])
			//}
		}
		for _, val := range su {
			if v.Summaryid == val.Summaryid {
				re[val.Cvalue] = val.Profits
				total[val.Cvalue] = t.Amount * val.Percent
			}
		}

		res = append(res, re)
	}

	result["profit"] = res

	result["keys"] = keys
	result["words"] = words

	return result, total, err
}

// 导出一次性分润报表
func (e *Summary) ExportExcelOnce(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["filefield"] = "name,username,teamname,percent,amount,profits"
	param["filename"] = "姓名,用户名,团队类型,分润比例,业绩,利润"
	param["title"] = "一次性分润报表" + param["date1"]
	URL, err = GetExcelTotal(e.StatementConfigOnce, param)
	return
}

// 导出汇总报表报表
func (e *Summary) ExportExcelSummary(param map[string]string) (URL string, err error) {
	param["isexp"] = "1"
	param["sheet"] = "sheet1"
	param["title"] = "明细报表" + param["date1"]
	URL, err = GetExcelSummary(e.StatementSummaryNew, param)

	return
}
