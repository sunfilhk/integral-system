package models

//import (
//	"log"
//	"xAdmin/config"
//	orm "xAdmin/database"
//)
//
//func RunInit(){
//	var conf Config
//	//adc:=sad.DeptId
//	//ss,_= utils.StringToInt64(sad.DeptId)
//	c ,err:= conf.GetConfig("customerratio")
//	if err!=nil||(c == Config{}){
//		sql := `insert into sys_config(configKey,configValue)value('customerratio',?),('salesmanratio',?)`
//		if err = orm.Eloquent.Exec(sql,config.ApplicationConfig.CustomerRatio,config.ApplicationConfig.SalesmanRatio).Error;err!=nil{
//			log.Fatal(err)
//		}
//	}
//}
