package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xAdmin/models"
	"xAdmin/utils"
)

func ProfitconfigList(c *gin.Context) {
	var u models.Profit
	var err error
	var res models.Response
	param := make(map[string]string)
	//param["deptId"] = c.Request.FormValue("deptId")
	param["userid"] = c.Request.FormValue("userid")
	param["profittype"] = c.Request.FormValue("profittype")
	param["deptId"] = c.Request.FormValue("deptId") //新加字段
	//if param["deptId"]==""{
	//	res.Msg = "团队id不能为空！"
	//	c.JSON(http.StatusOK, res.ReturnError(400))
	//	return
	//}
	//param["deptId"] = c.Request.FormValue("deptId")  //新加字段

	result, err := u.ProfitconfigList(param)
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

//新加获取一次性分润团队 外面
//func GetProfitteam(c *gin.Context) {
//	var data models.Profit
//	param := make(map[string]string)
//	param["dept_name"] = c.Request.FormValue("dept_name")
//	//param["path"] = c.Request.FormValue("path")
//	//param["deptid"]= c.Request.FormValue("dept_id") //新加
//	//param["userid"] = utils.GetUserIdStr(c)
//	//pkg.AssertErr(err, "", 400)
//	//var res models.Response
//	resl,err := data.Profitonce(param)
//	//if err!=nil{
//	//	res.Code = 0
//	//	res.Msg = err.Error()
//	//	c.JSON(http.StatusOK, res)
//	//	return
//	//}
//	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
//	res := utils.NewPageData(param, resl)
//	c.JSON(http.StatusOK, res)
//}

//添加一次性选择团队
//func Getonceinto(c *gin.Context) {
//	var data models.Profit
//	//var res models.Response
//
//	param := make(map[string]string)
//	param["teamname"] = c.Request.FormValue("teamname")
//	//param["pageSize"] = c.DefaultQuery("pageSize", "10")
//	//param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
//	//param["deptid"] = c.Request.FormValue("dept_id")
//
//	result, err := data.Gotonce(param)
//	//if err !=nil{
//	//	res.Msg = err.Error()
//	//	c.JSON(http.StatusOK, res.ReturnError(0))
//	//}
//	//res.Data = result
//	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
//
//	res := utils.NewPageData(param, result)
//
//	c.JSON(http.StatusOK, res)
//}

func ProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["percent"] = c.Request.FormValue("percent")
	param["userid"] = c.Request.FormValue("userid")
	param["teamname"] = c.Request.FormValue("teamname")
	//param["deptid"]= c.Request.FormValue("dept_id") //新加
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.ProfitconfigOnce(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func DelProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.DelProfitconfigOnce(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["percent"] = c.Request.FormValue("percent")
	param["userid"] = c.Request.FormValue("userid")
	param["id"] = c.Request.FormValue("id")
	param["teamname"] = c.Request.FormValue("teamname")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.UpdateProfitconfigOnce(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func ProfitconfigEdit(c *gin.Context) {
	var data models.ProfitEdit
	var res models.Response
	//param := make(map[string]string)
	var err error
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	//err := c.MustBindWith(&data, binding.JSON)
	//param["team_id"] = c.Request.FormValue("team_id")
	//// huazi:= c.Request.FormValue("team_id")
	//if param["team_id"]==""{
	//	res.Code = 0
	//	res.Msg = "团队id不能为空"
	//	c.JSON(http.StatusBadRequest, res)
	//	return
	//}
	err = c.ShouldBind(&data)
	fmt.Println("数据", data)

	//err := data.ProfitconfigEdit(param)
	err = data.ProfitEdit()
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.Msg = "操作成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetUserProfits(c *gin.Context) {
	var p models.TotalProfit
	userid := utils.GetUserIdStr(c)
	p = p.Profit(userid)
	var res models.Response
	res.Data = p
	c.JSON(http.StatusOK, res.ReturnOK())
}
