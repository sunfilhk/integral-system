package router

import (
	"github.com/robfig/cron/v3"
	"log"
	"strings"
	"sync"
	"time"
	orm "xAdmin/database"
	"xAdmin/models"
	"xAdmin/utils"
)

var audit sync.Mutex

// CheckInvestmentExpiration 每天凌晨1点定时巡检一次订单,并将对应的过期订单修改为过期状态
func CheckInvestmentExpiration() {
	c := cron.New()
	// 在每天凌晨1点定时执行对应任务
	spec := "0 1 * * ?"
	c.AddFunc(spec, func() {
		log.Println("开始巡检订单信息")
		var investlist models.Investlist
		list, err := investlist.InvestExpirationLists()
		if err != nil {
			log.Println("查询过期订单列表报错", err.Error())
			return
		}
		if len(list) == 0 {
			log.Println("无过期订单需要修改")
			return
		}
		for _, val := range list {
			err = editDueInvestment(val)
			if err != nil {
				continue
			}
		}
		log.Println("巡检订单信息成功")
	})
	go c.Start()
	defer c.Stop()
	select {}
}

func editDueInvestment(val models.Investlist) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		audit.Unlock()
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	audit.Lock()

	// 修改投资状态
	sql1 := `update investment set status = 5,expiration_date = date(now()),manually_end = 1 where id = ` + val.Id
	if err = orm1.Exec(sql1).Error; err != nil {
		log.Println("查询过期订单列表报错", err.Error())
		return
	}

	//备份分润信息
	sql2 := `select * from investmentprofit  where investmentid = ` + val.Id
	var ipe []models.InvestmentprofitExport
	if err = orm1.Raw(sql2).Scan(&ipe).Error; err != nil {
		log.Println("查询过期订单分润列表报错", err.Error())
		return
	}
	sql3 := `insert into investmentprofit_backup(investmentprofit_id,investmentid,userid,profits,batch)values`
	for k, v := range ipe {
		sql3 += `('` + v.ID + `','` + v.Investmentid + `','` + v.Userid + `','` + v.Profits + `','` + "0" + `')`
		if k < len(ipe)-1 {
			sql3 += ","
		}
	}
	if err = orm1.Exec(sql3).Error; err != nil {
		log.Println("备份分润表报错", err.Error())
		return
	}
	//删除分润
	sql4 := `delete from investmentprofit where investmentid = ` + val.Id
	if err = orm1.Exec(sql4).Error; err != nil {
		return
	}

	// 修改对应订单信息
	in := models.GetInvestmentById(val.Id)
	var ref models.Referrer
	ref = ref.GetReferrer(in.Userid)
	ref.Referrers += "," + in.Userid
	ref.Referrers = strings.Trim(ref.Referrers, ",")

	//删除累积业绩 减掉相关业务员业绩
	sql13 := `update sys_user set accumulative = accumulative - ` + utils.Float64ToString(in.Amount) + ` where user_id in (` + ref.Referrers + `) `
	if err = orm1.Exec(sql13).Error; err != nil {
		return
	}

	// log日志记录
	var loginlog models.LoginLog
	loginlog.LoginTime = time.Now()
	loginlog.CreateTime = time.Now()
	loginlog.Status = "0"
	loginlog.IsDel = "0"
	loginlog.Status = "1"
	loginlog.Platform = "Linux"
	loginlog.UserName = "admin"
	loginlog.CreateBy = "admin"
	loginlog.Msg = "系统自动更新过期订单列表状态为到期终止,状态为:5 已经终止,订单号为:" + val.Id
	loginlog.Create()
	return
}
