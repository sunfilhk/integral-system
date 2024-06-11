package models

import (
	"errors"
	"strconv"
	"strings"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type Profit struct {
	ID         string  `gorm:"column:id" json:"id"`
	Percent    float64 `gorm:"column:percent" json:"percent"`       //
	Userid     string  `gorm:"column:userid" json:"userid"`         //
	Profittype int     `gorm:"column:profittype" json:"profittype"` //
	DeptId     int     `gorm:"column:deptId" json:"deptId"`         //新加
	Teamname   string  `gorm:"column:team_name" json:"teamname"`    //新加
	//Teamnames      string        `gorm:"column:team_name" json:"teamnames"`  //新加
	Profitlevel int64 `gorm:"column:profitlevel" json:"profitlevel"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	//IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	NickName   string    `gorm:"column:nick_name" json:"nick_name"` //
	Username   string    `gorm:"column:username" json:"username"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` //创建时间
}

//	type Once struct {
//		Deptname      string        `gorm:"column:dept_name" json:"dpet_name"`
//		Profit        float64		`gorm:"column:profit" json:"profit"`
//		DeptId		int64			`gorm:"column:deptid" json:"deptid"`
//		Status		int64			`gorm:"column:status" json:"status"`
//	}
type Aa struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func (e *Profit) ProfitconfigList(param map[string]string) (result interface{}, err error) {

	//状态
	status := param["profittype"]
	key := param["deptId"]
	//deptid := param["deptid"]
	// fixme 用户vip等级在 user_level表里面需要单独判断获取，后期有时间可以改一下独立的接口
	if status == "2" {
		var userLevelConfig UserLevelConfig
		dataList, err1 := userLevelConfig.GetProfitConfigList(key)
		if err1 != nil {
			err = err1
			return
		}
		ret := make([]Profit, 0, len(dataList))
		for _, data := range dataList {
			var p Profit
			p.ID = strconv.FormatInt(data.ID, 10)
			p.Profittype = 2
			p.Profitlevel = data.LevelValue
			p.Percent = data.Percent
			p.DeptId = data.TeamId
			ret = append(ret, p)
		}
		result = ret
		return
	}
	//拼凑筛选条件sql
	sql := ` select p.*,u.nick_name,u.username,s.deptId from profitconfig p
		left join sys_user u ON p.userid = u.user_id
		LEFT JOIN sys_dept s  on  u.team_name=s.dept_name
		where 1=1 `
	//LEFT JOIN sys_dept s  on  u.team_name=s.dept_name
	if param["profittype"] != "" {
		sql += ` and profittype = '` + status + `'`

		if param["profittype"] == "1" {
			sql = `select if(c.count1=0,0,p.id)id,p.percent,1 profittype,u.user_id,p.profitlevel,u.nick_name,u.username FROM(									
					select count(1) count1 from profitconfig p
					left join sys_user u ON p.userid = u.user_id
					where userid = :userid and profittype =1
				)c
				left join profitconfig p on p.userid = if(c.count1=0,0,:userid) and p.profittype = if(c.count1=0,2,1)
				left join sys_user u ON u.user_id = :userid`
			sql = utils.SqlReplaceParames(sql, param)
		}
	}
	userid := param["userid"]
	if param["userid"] != "" {
		sql += ` and userid = '` + userid + `'`
	}

	if param["deptId"] != "" {
		sql += `and s.deptId = ` + param["deptId"]
	}

	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "profitlevel"
	param["order"] = "asc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Profit, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

//新加
//func (e *Profit) Profitonce(param map[string]string) (interface{} , error) {
//
//	sql := `select *from sys_dept where is_del= 0`  //加了dept_id和dept_name
//	keyword := param["dept_name"]
//	if keyword != "" {
//		sql += ` and (dept_name like '%` + keyword + `%') `
//	}
//	//param["total"] = GetTotalCount(sql)
//	////分页 and 排序
//	//param["sort"] = "user_id"
//	//param["order"] = "desc"
//	var w []Once
//	//sql += utils.LimitAndOrderBy(param)
//	err := orm.Eloquent.Raw(sql).Scan(&w).Error
//
//	return w, err
//}

//func (e *Profit) Gotonce(param map[string]string) (interface{} , error) {
//
//	sql := `select *from sys_dept where is_del= 0`  //加了dept_id和dept_name
//	keyword := param["dept_name"]
//	if keyword != "" {
//		sql += ` and (dept_name like '%` + keyword + `%') `
//	}
//	//param["total"] = GetTotalCount(sql)
//	////分页 and 排序
//	//param["sort"] = "user_id"
//	//param["order"] = "desc"
//	var w []Once
//	//sql += utils.LimitAndOrderBy(param)
//	err := orm.Eloquent.Raw(sql).Scan(&w).Error
//
//	return w, err
//}

func (e *Profit) ProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)
	var count int
	orm.Eloquent.Table("profitconfig").Where("userid = ?  ", param["userid"]).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}
	var d Profit
	orm.Eloquent.Table("sys_user").Where("user_id = ?  ", param["userid"]).Scan(&d)
	if param["teamname"] != d.Teamname {
		err = errors.New("团队不匹配")
		return
	}
	sql := ` insert into profitconfig(percent,profittype,userid,team_name)value(:percent,0,:userid,:teamname)`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) DelProfitconfigOnce(param map[string]string) (err error) {

	sql := ` delete from profitconfig where id =:id and profittype =0 `
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) UpdateProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)

	sql := ` update profitconfig set percent=:percent,userid=:userid,team_name =:teamname where id=:id `
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}

type ProfitEdit struct {
	Ids string `json:"ids"`
	//Userid string `json:"userid"`
	Profit []ProfitConfig `json:"profit"`
}
type ProfitConfig struct {
	//ID           string     `gorm:"column:id" json:"id"`
	Percent     float64 `gorm:"column:percent" json:"percent"`       //
	Userid      string  `gorm:"column:userid" json:"userid"`         //
	Profittype  int     `gorm:"column:profittype" json:"profittype"` //
	Profitlevel float64 `gorm:"column:profitlevel" json:"profitlevel"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	TeamId string `gorm:"column:team_id" json:"team_id"` //是否删除
	//CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//创建时间
}

func (e *ProfitEdit) ProfitEdit() (err error) {
	orm1 := orm.Eloquent.Begin()

	//增加新配置
	flag := false
	sql1 := `insert into user_level(levelvalue, percent, percentreality,team_id)values `
	//fmt.Println("team_id:",param["team_id"])
	percentReality := 0.0
	for k, v := range e.Profit {
		flag = true
		if k == 0 && v.Profitlevel != 0 {
			return nil
		}
		percentReality += v.Percent
		//fmt.Println("team_id:",v.TeamId)
		//fmt.Println("percent:",v.Percent)
		sql1 += `(` + utils.Float64ToString(v.Profitlevel) + "," + utils.Float64ToString(v.Percent) + "," + utils.Float64ToString(percentReality) + "," + v.TeamId + `)`
		if k < len(e.Profit)-1 {
			sql1 += ","
		}
	}
	//删除原有配置
	sql := ` delete from user_level where id in ('` + strings.Replace(e.Ids, ",", "','", -1) + `')`
	if err = orm1.Exec(sql).Error; err != nil {
		orm1.Rollback()
		return err
	}
	if flag {
		if err = orm1.Exec(sql1).Error; err != nil {
			orm1.Rollback()
			return err
		}
	}
	orm1.Commit()
	return
}

type TotalProfit struct {
	Total     float64 `json:"total"`
	Yesterday float64 `json:"yesterday"`
	Today     float64 `json:"today"`
	Month     float64 `json:"month"` //23/3/21修改
	//Accumulated  string `json:"accumulated"`
}

func (t *TotalProfit) Profit(userid string) TotalProfit {
	sql := `select round(sum(total)) total,sum(yesterday) yesterday,sum(today)today,sum(month) month  from(
	select sum(p.profits) total,0 yesterday,0 today ,0 month
		from investment i
		left join investmentprofit p on i.id = p.investmentid
		where p.userid = ?
	UNION
	select 0 total,round(sum(p.profits)) yesterday,0 today , 0 month from investment i 
			left join investmentprofit p on p.investmentid = i.id
			where create_time >= DATE_SUB(date(now()),INTERVAL 1 DAY) 
			and create_time < date(now()) and p.userid = ?
	UNION
	select 0 total,0 yesterday,round(sum(p.profits)) profits,0 month from investment i 
			left join investmentprofit p on p.investmentid = i.id
			where create_time < DATE_ADD(date(now()),INTERVAL 1 DAY) 
			and create_time >= date(now()) and p.userid = ?
	UNION
	select 0 total,0 yesterday,0 today,round(sum(p.profits)) monthprofits from investment i 
                        left join investmentprofit p on p.investmentid = i.id
                        where date_format(create_time,'%Y-%m')=date_format(now(),'%Y-%m') 
                         and p.userid = ?
			)a`
	var tp TotalProfit
	orm.Eloquent.Raw(sql, userid, userid, userid, userid).Scan(&tp)
	return tp
}
