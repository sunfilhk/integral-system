package models

import (
	"fmt"
	orm "xAdmin/database"
)

type UserLevelConfig struct {
	ID             int64   `gorm:"column:id" json:"id"`
	LevelValue     int64   `gorm:"column:levelvalue" json:"level_value"`
	Percent        float64 `gorm:"column:percent" json:"percent"`
	PercentReality float64 `gorm:"column:percentreality" json:"percent_reality"`
	TeamId         int     `gorm:"column:team_id" json:"team_id"`
}
type Ouserlevel struct {
	ID         int64   `gorm:"column:id" json:"id"`
	LevelValue int64   `gorm:"column:levelvalue" json:"level_value"`
	Percent    float64 `gorm:"column:percent" json:"percent"`
	//PercentReality 		float64			`gorm:"column:percentreality" json:"percent_reality"`
}

func (this *UserLevelConfig) GetProfitConfigList(key string) (confs []UserLevelConfig, err error) {
	sql := `select * from user_level where team_id=? order by levelvalue asc`
	fmt.Println("类型：", key)
	confs = make([]UserLevelConfig, 0)
	if err = orm.Eloquent.Raw(sql, key).Scan(&confs).Error; err != nil {
		return
	}
	return
}

//func (this *Ouserlevel) getteamList() (confson []Ouserlevel, err error) {
//	sql := `select * from userlevel order by id asc`
//
//	confson = make([]Ouserlevel, 0)
//	if err = orm.Eloquent.Raw(sql).Scan(&confson).Error; err != nil {
//		return
//	}
//	return
//}

type vipLevelData struct {
	VipLevel   string  `gorm:"column:vipLevel" json:"vipLevel"`
	LevelValue float64 `gorm:"column:levelvalue" json:"levelValue"`
	ID         int64   `gorm:"column:id" json:"id"`
}

//func (this *UserLevelConfig) EditUserVipLevel(userID string, levelID string) (err error) {
//	sql1 := `select levelvalue from user_level where id = ` + levelID
//	findVip := make([]UserLevelConfig, 0)
//	if err = orm.Eloquent.Raw(sql1).Scan(&findVip).Error; err != nil {
//		return
//	}
//	if len(findVip) <= 0 {
//		err = errors.New("没找到vip配置")
//		return
//	}
//
//	levelValueStr := strconv.FormatInt(findVip[0].LevelValue, 10)
//	sql2 := `update sys_user set diff_value = (diff_value + accumulative - ` + levelValueStr +  `), accumulative = `+ levelValueStr  +`,update_time = now() where user_id = ` + userID
//	if err = orm.Eloquent.Exec(sql2).Error; err != nil {
//		return
//	}
//
//	return
//}

//func (this *UserLevelConfig) GetVipLevelList(userID string) (ret interface{}, err error) {
//	accumulative, err1 := this.getUserAccumulative(userID)
//	if err1 != nil {
//		err = err1
//		return
//	}
//
//	sql := `select CONCAT('V',(@rowNO := @rowNo+1)) AS vipLevel,a.* from  user_level a,(select @rowNO :=0) b ORDER BY levelvalue asc`
//	vips := make([]vipLevelData, 0)
//	if err = orm.Eloquent.Raw(sql).Scan(&vips).Error; err != nil {
//		return
//	}
//
//	if len(vips) <= 0 {
//		err = errors.New("没找到配置数据")
//		return
//	}
//	//
//	if accumulative <= 0 {
//		retData := map[string]interface{} {
//			"curLevelID":vips[0].ID,
//			"curLevel":vips[0].VipLevel,
//			"vipList":vips,
//		}
//		ret = retData
//		return
//	}
//
//	curLevel := ""
//	curLevelID := int64(0)
//	allLen := len(vips)
//	for i, m := range vips {
//		if i+1 == allLen {
//			curLevel = m.VipLevel
//			curLevelID = m.ID
//			break
//		}
//		next := vips[i+1]
//		if accumulative >= m.LevelValue && accumulative < next.LevelValue {
//			curLevel = m.VipLevel
//			curLevelID = m.ID
//			break
//		}
//	}
//
//
//	retData := map[string]interface{} {
//		"curLevelID":curLevelID,
//		"curLevel":curLevel,
//		"vipList":vips,
//	}
//	ret = retData
//	return
//}
//
//func (this *UserLevelConfig) getUserAccumulative(userID string) (val float64, err error) {
//	sql := `select accumulative from sys_user where user_id=` + userID
//
//	type findData struct {
//		Accumulative float64  `gorm:"column:accumulative"`
//	}
//	finds := make([]findData, 0)
//	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
//		return
//	}
//
//	if len(finds) != 1 {
//		err = errors.New("用户表数据量不对")
//		return
//	}
//
//	val = finds[0].Accumulative
//	return
//}
