package models

import (
	"sort"
	orm "xAdmin/database"
	"xAdmin/utils"
)

type SettleCustomerHistory struct {
	ID         int64   `gorm:"column:id" json:"id"`
	Name       string  `gorm:"column:name" json:"name"`               // 名称
	Bank       string  `gorm:"column:bank" json:"bank"`               // 开户行
	BankNum    string  `gorm:"column:banknum" json:"banknum"`         // 卡号
	UserID     int64   `gorm:"column:user_id" json:"user_id"`         // 业务员ID
	NickName   string  `gorm:"column:nick_name" json:"nickname"`      // 业务员昵称
	SettleTime string  `gorm:"column:settle_time" json:"settle_time"` // 结算时间
	CreateTime string  `gorm:"column:create_time" json:"create_time"` // 创建时间
	InvestID   string  `gorm:"column:invest_id" json:"invest_id"`     // investment表ID
	CustomerID int64   `gorm:"column:customer_id" json:"customer_id"` // customer表ID
	TeamName   string  `gorm:"column:team_name" json:"team_name"`     // 团队类型
	Amount     float64 `gorm:"column:amount" json:"amount"`           // 业绩
	Profits    float64 `gorm:"column:profits" json:"profits"`         // 分润
}

type sortHistory []SettleCustomerHistory

func (s sortHistory) Len() int           { return len(s) }
func (s sortHistory) Less(i, j int) bool { return s[i].ID > s[j].ID }
func (s sortHistory) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (this *SettleCustomerHistory) GetHistory(param map[string]string) (dataList interface{}, total int, err error) {
	//pageSize:=param["pageSize"]
	//pageIndex:=param["pageIndex"]
	//pSize, err1 := strconv.ParseInt(pageSize, 10, 64)
	//if err1 != nil {
	//	err = err1
	//	return
	//}
	//
	//pIndex, err1 := strconv.ParseInt(pageIndex, 10, 64)
	//if err1 != nil {
	//	err = err1
	//	return
	//}

	//sql2 := `select id from customer_settle`
	//total = GetTotalCount1(sql2)

	//start := int64(total) - pSize*pIndex
	//if start <= 0 {
	//	pSize = pSize + start
	//	start = 0
	//}
	//
	//if pSize < 0 {
	//	pSize = 0
	//}
	start := param["start"]
	end := param["end"]
	keyword := param["keyword"]
	invest_id := param["invest_id"]
	//sql := `select * from customer_settle limit ` + strconv.FormatInt(start, 10) + `,` + strconv.FormatInt(pSize, 10)
	sql := `select *from(select t.*,i.create_time,i.amount,(i.amount*i.profit)profits from (select c.*,u.team_name from customer_settle c LEFT JOIN sys_user u on c.user_id=u.user_id )t LEFT JOIN investment i on t.user_id=i.userid GROUP BY t.id having t.id)a
							where a.settle_time <= '` + end + `' and '` + start + `' <=a.settle_time ` //+ strconv.FormatInt(start, 10) + `,` + strconv.FormatInt(pSize, 10) //修改展示团队字段
	if keyword != "" {
		sql += `and  ( nick_name like '%` + keyword + `%' or name like '%` + keyword + `%')  `
	}
	if invest_id != "" {
		sql += ` and  invest_id= ` + invest_id
	}
	total = GetTotalCount1(sql)
	//sql += `limit ` + strconv.FormatInt(start, 10) + `,` + strconv.FormatInt(pSize, 10)
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	retList := make([]SettleCustomerHistory, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&retList).Error; err != nil {
		return
	}

	sort.Sort(sortHistory(retList))
	dataList = retList
	return
}
