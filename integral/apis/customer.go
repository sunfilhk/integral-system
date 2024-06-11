package apis

import (
	"net/http"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
)

func GetcustomerList(c *gin.Context) {
	var u models.Customer
	var err error
	param := make(map[string]string)

	param["keyword"] = c.Request.FormValue("keyword")
	param["teamid"] = c.Request.FormValue("teamid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["status"] = c.Request.FormValue("status")

	param["role"] = utils.GetRolekey(c)
	switch param["role"] {
	case "finance":
	case "boss":
	case "admin":
	default:
		param["userid"] = utils.GetUserIdStr(c)
	}
	//移动端判断是否可新增投资
	param["userid1"] = utils.GetUserIdStr(c)
	//var re models.Response
	//re.ReturnError(401)
	//c.JSON(http.StatusOK, re)
	//return

	result, err := u.CustomerList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

//新加模块
func CustomerLists(c *gin.Context) {
	var i models.Customer
	var err error
	param := make(map[string]string)

	param["keyword"] = c.Request.FormValue("keyword")
	//param["teamname"] = c.Request.FormValue("teamname")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//param["status"] = c.Request.FormValue("status")
	param["role"] = utils.GetRolekey(c)
	switch param["role"] {
	case "finance":
	case "boss":
	case "admin":
	default:
		param["userid"] = utils.GetUserIdStr(c)
	}
	//移动端判断是否可新增投资
	param["userid1"] = utils.GetUserIdStr(c)
	//var re models.Response
	//re.ReturnError(401)
	//c.JSON(http.StatusOK, re)
	//return

	result, err := i.CustomerListss(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func CustomerAdd(c *gin.Context) {
	var data models.Customer
	var d models.Dept
	param := make(map[string]string)
	param["name"] = c.Request.FormValue("name")
	param["phone"] = c.Request.FormValue("phone")
	param["userid"] = utils.GetUserIdStr(c)
	//userid :=utils.GetUserIdStr(c)
	//u.Id,_=utils.StringToInt64(c.Request.FormValue(userid))
	param["bank"] = c.Request.FormValue("bank")
	param["banknum"] = c.Request.FormValue("banknum")
	param["identity"] = c.Request.FormValue("identity")
	param["sex"] = c.Request.FormValue("sex")

	param["amount"] = c.Request.FormValue("amount")
	param["dept_name"] = c.Request.FormValue("dept_name")
	//param["profittype"]= c.Request.FormValue("profittype")

	//dd,_=dd.Getdept(userid)

	//d, _ = d.GetConfig(param)  //修改函数里面的值
	df := param["userid"]
	d, _ = d.Getteam(df)
	//param["deptid"] = d.Deptid
	param["profit"] = utils.Float64ToString(d.Profit)
	param["remark"] = c.Request.FormValue("remark")
	param["date"] = c.Request.FormValue("date")
	var res models.Response
	if param["date"] == "" {
		res.Code = 0
		res.Msg = "date不能为空"
		c.JSON(http.StatusOK, res)
		return
	}
	//pkg.AssertErr(err, "", 400)
	err := data.CustomerAdd(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func CustomerEdit(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["customerid"] = c.Request.FormValue("id")
	param["namenew"] = c.Request.FormValue("name")
	param["identitynew"] = c.Request.FormValue("identity")
	param["phonenew"] = c.Request.FormValue("phone")
	param["userid"] = c.Request.FormValue("salesmanid")
	param["banknew"] = c.Request.FormValue("bank")
	param["banknumnew"] = c.Request.FormValue("banknum")
	param["sexnew"] = c.Request.FormValue("sex")
	var res models.Response
	var cus models.Customer
	cus, _ = cus.NewCustomer(param["customerid"])
	if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	err := data.CustomerEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func CustomerProfitEdit(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["customerid"] = c.Request.FormValue("customerid")
	param["profit"] = c.Request.FormValue("profit")
	var res models.Response
	err := data.CustomerProfitEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetCustomerByid(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	id := c.Request.FormValue("customerid")
	var res models.Response
	re, err := data.NewCustomer(id)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}
	res.Data = re
	c.JSON(http.StatusOK, res.ReturnOK())
}
