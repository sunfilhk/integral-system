package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	orm "xAdmin/database"
	"xAdmin/utils"
)

// 客户表
type Customer struct {
	ID             string    `gorm:"column:id" json:"id"`
	InvestmentID   string    `gorm:"column:investment_id" json:"investment_id"`
	CustomerId     string    `gorm:"column:customer_id" json:"customerid"`
	Name           string    `gorm:"column:name" json:"name"`                //姓名
	Sex            string    `gorm:"column:sex" json:"sex"`                  //昵称
	Identity       string    `gorm:"column:identity" json:"identity"`        //身份证号
	Phone          string    `gorm:"column:phone" json:"phone"`              //手机号
	Bank           string    `gorm:"column:bank" json:"bank"`                //银行
	Banknum        string    `gorm:"column:banknum" json:"banknum"`          //银行卡号
	Userid         string    `gorm:"column:userid" json:"userid"`            //业务员ID
	IsDel          int       `gorm:"column:is_del" json:"is_del"`            //是否删除
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`  //创建时间
	MonthlyTime    time.Time `gorm:"column:monthly_time" json:"update_time"` //更新时间
	Amount         float64   `gorm:"column:amount" json:"amount"`
	Remark         string    `gorm:"column:remark" json:"remark"`
	Status         int       `gorm:"column:status" json:"status"`
	Profits        float64   `gorm:"column:profits" json:"profits"`
	Profit         float64   `gorm:"column:profit" json:"profit"`
	NickName       string    `gorm:"column:nick_name" json:"nick_name"`
	Username       string    `gorm:"column:username" json:"username"`
	ExpirationDate time.Time `gorm:"column:expiration_date" json:"expiration_date"` // 合同到期时间
	Totalamount    string    `gorm:"column:totalamount" json:"totalamount"`
	Totalprofit    string    `gorm:"column:totalprofit" json:"totalprofit"`
	Teamname       string    `gorm:"column:teamname" json:"teamname"`       //新添加字段
	Profittype     int64     `gorm:"column:profittype" json:"profittype"`   //新添加字段
	Ordertime      time.Time `gorm:"column:ordertime" json:"ordertime"`     //新添加字段
	InvestTime     time.Time `gorm:"column:invest_time" json:"invest_time"` //新添加字段
	TeamId         int64     `gorm:"column:team_id" json:"teamid"`
}
type CustomerView struct {
	Customer
	Flag int `json:"flag"`
}

func (e *Customer) NewCustomer(customerid string) (Customer, error) {
	sql2 := `select c.*,u.nick_name,u.username from customer c
					left join sys_user u on c.userid = u.user_id 
					where id = '` + customerid + "'"
	fmt.Println("dwl", customerid)
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error; err != nil {
		fmt.Errorf("获取客户信息失败：", err)
		return c, err
	}
	return c, nil
}
func NewCustomer1(param map[string]string) error {
	sql2 := `select * from customer where id = '` + param["customerid"] + "'"
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
	param["userid"] = c.Userid
	param["sex"] = c.Sex
	return nil
}

// 客户管理页面数据
func (e *Customer) CustomerList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	con := `,if(` + param["userid1"] + `=c.userid,1,0) flag `
	if param["role"] == "admin" || param["role"] == "finance" { //22/8/19修改添加finance
		con = `,1 flag `
	}
	//sql := ` select c.*,u.nick_name,u.username` + con + ` from customer c
	//			left join sys_user u on c.userid = u.user_id
	//			where c.is_del=0`
	sql := ` select c.*,u.nick_name,u.username,a.profit,a.amount totalamount,a.amount*a.profit totalprofit,u.team_name teamname,u.team_id  ` + con + `  from customer c
                                left join sys_user u on c.userid = u.user_id	
							left join (SELECT sum(amount)amount,userid,profit,customerid FROM investment WHERE status=0  GROUP BY customerid )a on 	c.id = a.customerid 
							where c.is_del=0  `

	//搜索框
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' or c.name like '%` + keyword + `%' or u.nick_name like '%` + keyword + `%' or u.username like '%` + keyword + `%') `
	}
	//状态
	status := param["status"]
	if param["status"] != "" {
		sql += ` and c.status = ` + status
	}
	sbwy := param["teamid"]
	if param["teamid"] != "" {
		sql += `and u.team_id =  ` + sbwy
	}
	if param["userid"] != "" {
		userID := param["userid"]
		sql += ` and c.userid = ` + userID

		// 显示自己直属下级或者全部下级的暂时不处理
		//filterUsers, err := e.getFilterUsers(userID)		// 显示自己所有下级
		//filterUsers, err := e.getNextUserIDs(userID)		// 显示自己直接下级
		//if err != nil {
		//	sql += ` and c.userid = ` + userID
		//} else {
		//	sql += ` and c.userid in (` + filterUsers + `)`
		//}
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]CustomerView, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Customer) CustomerListss(param map[string]string) (result interface{}, err error) {
	con := `,if(` + param["userid1"] + `=a.userid,1,0) flag `
	if param["role"] == "admin" || param["role"] == "finance" {
		con = `,1 flag `
	}
	//sql := ` select c.*,u.nick_name,u.username` + con + ` from customer c
	//			left join sys_user u on c.userid = u.user_id
	//			where c.is_del=0`
	sql := ` select c.*,u.nick_name,u.username,a.profit,a.amount totalamount,a.amount*a.profit totalprofit,a.customerid customer_id,u.team_name teamname ` + con + `  from customer c
                                left join sys_user u on c.userid = u.user_id	
							left join (SELECT sum(amount)amount,userid,profit,customerid FROM investment WHERE status in (0,1,3,4,5,6) GROUP BY customerid )a on 	c.id = a.customerid 
							where c.is_del=0  `
	//搜索框
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.name like '%` + keyword + `%' ) `
	}
	if param["userid"] != "" {
		userID := param["userid"]
		sql += ` and c.userid = ` + userID
	}
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]CustomerView, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Customer) getFilterUsers(userID string) (userIDs string, err error) {
	// 先判断部门，部门没有就取推荐人玩家列表
	userIDs, err = e.getDeptUsers(userID)
	if err == nil {
		return
	}
	return e.getUserReferralsIDs(userID)
}

// 获取下一级玩家列表
func (e *Customer) getNextUserIDs(userID string) (userIDs string, err error) {
	type findReferrer struct {
		UserID int64 `gorm:"column:user_id"`
	}
	sql1 := `select user_id from sys_user where referrer = ` + userID
	var finds []findReferrer
	err = orm.Eloquent.Raw(sql1).Scan(&finds).Error
	if err != nil {
		return
	}

	length := len(finds)
	if length <= 0 {
		err = errors.New("下级列表为空")
		return
	}

	userIDs = userID + ","
	for i, f := range finds {
		userIDs += strconv.FormatInt(f.UserID, 10)
		if i != length-1 {
			userIDs += ","
		}
	}
	return
}

func (e *Customer) getDeptUsers(userID string) (userIDs string, err error) {
	// 是否leader
	sql := `select deptId from sys_dept where leader_id=` + userID
	type findDept struct {
		Deptid int64 `gorm:"column:deptId;primary_key"`
	}
	var depts []findDept
	if err = orm.Eloquent.Raw(sql).Scan(&depts).Error; err != nil {
		return
	}

	if len(depts) <= 0 {
		err = errors.New("no records")
		return
	}

	deptStr := ""
	l := len(depts)
	for i, de := range depts {
		deptStr += strconv.FormatInt(de.Deptid, 10)
		if i != l-1 {
			deptStr += ","
		}
	}
	sql2 := `select user_id from sys_user where dept_id in (` + deptStr + `)`
	var findUsers []SysUserId
	if err = orm.Eloquent.Raw(sql2).Scan(&findUsers).Error; err != nil {
		return
	}

	userIDs = ""
	lu := len(findUsers)
	for i, u := range findUsers {
		userIDs += strconv.FormatInt(u.Id, 10)
		if i != lu-1 {
			userIDs += ","
		}
	}
	return
}

func (e *Customer) getUserReferralsIDs(userID string) (ret string, err error) {
	sql := `select referrals from referrer where userid=` + userID

	var referrers []Referrer
	if err = orm.Eloquent.Raw(sql).Scan(&referrers).Error; err != nil {
		return
	}
	if len(referrers) != 1 {
		err = errors.New("record error")
		return
	}
	ret = referrers[0].Referrals
	if len(ret) <= 0 {
		err = errors.New("no record")
		return
	}
	if len(ret) > 0 && ret[0] == ',' {
		ret = ret[1:]
	}
	ret += "," + userID // 追加上自己的ID
	return
}

//	for key, val := range param {
//		if val != "" && key != "id" {
//			con +=  key + "='" + val.(string) + "',"
//		}
//	}
func (e *Customer) CustomerAdd(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	param["id1"] = strconv.FormatInt(utils.Node().Generate().Int64(), 10)
	con := ""
	con1 := ""
	if param["identity"] != "" {
		con = `,identity`
		con1 = `,:identity`
	}
	sql := ` insert into customer(id,name` + con + `,sex,phone,bank,banknum,userid,status )value(:id,:name` + con1 + `,:sex,:phone,:bank,:banknum,:userid,1)` //新加teamid
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}
	sql1 := `insert into investment(id,userid,amount,customerid,remark,status,invest_time,monthly_time,expiration_date,profit)value(:id1,:userid,:amount,:id,:remark,1,:date,date_add(:date,interval 1 month),DATE_FORMAT(date_add(:date,interval 1 year),'%Y-%m-%d 00:00:00'),:profit)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	if err = orm1.Exec(sql1).Error; err != nil {
		return err
	}
	sql2 := `insert into customer_log(customerid,salesmanid,name` + con + `,sex,phone,bank,banknum,investmentid,amount,iremark,edittype)
					value(:id,:userid,:name` + con1 + `,:sex,:phone,:bank,:banknum,:id1,:amount,:remark,0)`
	sql2 = utils.SqlReplaceParames(sql2, param)
	err = orm1.Exec(sql2).Error
	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}

func (e *Customer) CustomerEdit(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	sql := ` update customer set status = 1 where id = :customerid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm1.Exec(sql).Error; err != nil {
		return err
	}

	if err = NewCustomer1(param); err != nil {
		return err
	}

	sql1 := `insert into customer_log(customerid,salesmanid,name,namenew,identity,identitynew,phone,phonenew,bank,banknew,banknum,banknumnew,sex,sexnew,edittype)
							    value(:customerid,1 ,:name,:namenew,:identity,:identitynew,:phone,:phonenew,:bank,:banknew,:banknum,:banknumnew,:sex,:sexnew,1)`
	sql1 = utils.SqlReplaceParames(sql1, param)
	err = orm1.Exec(sql1).Error
	return
}

func (e *Customer) CustomerProfitEdit(param map[string]string) (err error) {

	sql := ` update customer set profit =:profit where id = :customerid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}
	return
}

func (e *Customer) GetCustomerByid(param map[string]string) (err error) {

	sql := ` update customer set profit =:profit where id = :customerid`
	sql = utils.SqlReplaceParames(sql, param)
	if err = orm.Eloquent.Exec(sql).Error; err != nil {
		return err
	}
	return
}
