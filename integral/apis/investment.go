package apis

import (
	"fmt"
	"net/http"
	"xAdmin/config"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"

	"github.com/gin-gonic/gin"
)

func InvestmentList(c *gin.Context) {
	var u models.Investment
	var err error
	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["customerid"] = c.Request.FormValue("customerid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//param["status"] = c.Request.FormValue("status")
	param["userid"] = utils.GetUserIdStr(c)

	param["roleKey"] = utils.GetRolekey(c)

	result, err := u.InvestmentList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func InvestListss(c *gin.Context) {
	var d models.Investlist
	var err error
	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["id"] = c.Request.FormValue("id")
	param["deptid"] = c.Request.FormValue("deptid")
	param["status"] = c.Request.FormValue("status")
	param["nick_name"] = c.Request.FormValue("nick_name")
	param["expiration_date"] = c.Request.FormValue("expiration_date")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	param["createStart"] = c.Request.FormValue("createStart")
	param["createEnd"] = c.Request.FormValue("createEnd")
	//param["customerid"] = c.Request.FormValue("customerid")

	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//param["status"] = c.Request.FormValue("status")
	param["userid"] = utils.GetUserIdStr(c)

	param["roleKey"] = utils.GetRolekey(c)
	param["end"] += " 23:59:59"
	param["createEnd"] += " 23:59:59"
	result, err := d.InvestLists(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

// 搜索将要到期的订单（按天数）
func Searchs(c *gin.Context) {
	var y models.Investlist
	var err error
	param := make(map[string]string)
	param["status"] = c.Request.FormValue("status")
	param["days"] = c.Request.FormValue("days")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//param["start"] = c.Request.FormValue("start")
	//param["end"] = c.Request.FormValue("end")

	result, err := y.Search(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

func InvestmentAdd(c *gin.Context) {
	var data models.Investment
	var d models.Dept
	param := make(map[string]string)
	param["amount"] = c.Request.FormValue("amount")
	param["name"] = c.Request.FormValue("name")
	param["userid"] = utils.GetUserIdStr(c)
	pig := param["userid"]

	//新添加字段选出对应的值计算
	//d, _ = d.GetConfig("customerratio")  //原来获取默认值
	d, _ = d.Getteam(pig)
	param["profit"] = utils.Float64ToString(d.Profit)
	param["remark"] = c.Request.FormValue("remark")
	//param["customerid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	param["userid"] = utils.GetUserIdStr(c)
	param["date"] = c.Request.FormValue("date")
	var res models.Response
	var cus models.Customer
	fmt.Println("sbwyi", param["customerid"])
	cus, _ = cus.NewCustomer(param["customerid"])
	//原来判断
	if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	if param["date"] == "" {
		res.Code = 0
		res.Msg = "date不能为空"
		c.JSON(http.StatusOK, res)
		return
	}
	//if time.Now().Hour()>=23{
	//	res.Code = 0
	//	res.Msg = "已超过23点，请第二天请早"
	//	c.JSON(http.StatusOK, res)
	//	return
	//}
	//pkg.AssertErr(err, "", 400)

	err := data.InvestmentAdd(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

// 新模块新增客户
func NewmoduleInvestmentAdd(c *gin.Context) {
	var data models.Investment
	//var d models.Dept
	param := make(map[string]string)
	param["amount"] = c.Request.FormValue("amount")
	param["name"] = c.Request.FormValue("name")
	param["profit"] = c.Request.FormValue("profit")
	pig := param["profit"]

	//新添加字段选出对应的值计算
	//d, _ = d.GetConfig("customerratio")  //原来获取默认值
	//d, _ = d.Getteam(pig)
	//a,_:=strconv.ParseFloat((pig,64)
	param["profit"] = pig
	param["remark"] = c.Request.FormValue("remark")
	//param["customerid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	//param["userid"] = utils.GetUserIdStr(c)
	param["date"] = c.Request.FormValue("date")
	var res models.Response
	//var cus models.Customer
	fmt.Println("sbwyi", param["customerid"])
	//cus, _ = cus.NewCustomer(param["customerid"])
	//原来判断
	//if cus.Userid != param["userid"]&&utils.GetRolekey(c)!="admin"{
	//	res.Msg = "无权限操作该客户"
	//	c.JSON(http.StatusOK, res.ReturnError(400))
	//	return
	//}
	if param["date"] == "" {
		res.Code = 0
		res.Msg = "date不能为空"
		c.JSON(http.StatusOK, res)
		return
	}
	//if time.Now().Hour()>=23{
	//	res.Code = 0
	//	res.Msg = "已超过23点，请第二天请早"
	//	c.JSON(http.StatusOK, res)
	//	return
	//}
	//pkg.AssertErr(err, "", 400)

	err := data.NewInvestmentAdd(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func InvestmentEdit(c *gin.Context) {
	var data models.Investment
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["investmentid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	param["amountnew"] = c.Request.FormValue("amount")
	param["iremarknew"] = c.Request.FormValue("remark")
	param["userid"] = utils.GetUserIdStr(c)
	param["date"] = c.Request.FormValue("date") //修改时间
	var res models.Response
	var cus models.Customer
	cus, _ = cus.NewCustomer(param["customerid"])
	if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	err := data.InvestmentEdit(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentBreak(c *gin.Context) {
	var data models.Investment
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["investmentid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	param["userid"] = utils.GetUserIdStr(c)
	param["breaktype"] = c.Request.FormValue("breaktype")

	var res models.Response
	var cus models.Customer
	cus, _ = cus.NewCustomer(param["customerid"])
	//原来的判断语句
	//if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
	//	res.Msg = "无权限操作该客户"
	//	c.JSON(http.StatusOK, res.ReturnError(400))
	//	return
	//}
	//22/8/19修改过后的语句
	if cus.Userid == param["userid"] && utils.GetRolekey(c) == "common" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	// 判断传的终止类型是否正确
	if param["breaktype"] != config.BreakTypeTimeDue && param["breaktype"] != config.BreakTypeByHand {
		res.Msg = "终止类型错误"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	err := data.InvestmentBreak(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentRevoke(c *gin.Context) {
	var data models.Investment
	param := make(map[string]string)
	param["investmentid"] = c.Request.FormValue("id")
	param["customerid"] = c.Request.FormValue("customerid")
	param["userid"] = utils.GetUserIdStr(c)

	var res models.Response
	var cus models.Customer
	cus, _ = cus.NewCustomer(param["customerid"])
	if cus.Userid != param["userid"] && utils.GetRolekey(c) != "admin" {
		res.Msg = "无权限操作该客户"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	err := data.InvestmentRevoke(param)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetInvestmentByid(c *gin.Context) {
	var data models.Investment
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)

	id := c.Request.FormValue("id")

	var res models.Response
	re, err := data.InvestmentById(id)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
	}
	res.Data = re

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentExport(c *gin.Context) {
	var u models.Investment
	param := make(map[string]string)
	var res models.Response

	result, err := u.ExportInvestment(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

func InvestmentImport(c *gin.Context) {
	var res models.Response

	file, err := c.FormFile("file")
	if err != nil {
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	var u models.Investment
	f, _ := file.Open()
	err = u.InvestmentImport(f, file.Size)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	if file.Size == 0 {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Msg = "导入成功！"

	c.JSON(http.StatusOK, res.ReturnOK())
}

func EmptyInvestment(c *gin.Context) {
	var res models.Response
	var u models.Investment
	err := u.DeleteInvestment()
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	res.Msg = "删除成功！"

	c.JSON(http.StatusOK, res.ReturnOK())
}
