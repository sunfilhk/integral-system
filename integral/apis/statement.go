package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"
)

func StatementSalesman(c *gin.Context) {
	var u models.Statement
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, total, err := u.StatementSalesmanNew(param)
	if err != nil {
		re.Msg = "没有数据"
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}

func Summrydetail(c *gin.Context) {
	var u models.Statement
	var err error
	param := make(map[string]string)
	param["userid"] = c.Request.FormValue("userid")

	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["end"] += " 23:59:59"
	//param["status"] = c.Request.FormValue("status")
	//param["userid"] = utils.GetUserIdStr(c)

	param["roleKey"] = utils.GetRolekey(c)

	result, total, err := u.Summrydetail(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}

func StatementSalesmanExport(c *gin.Context) {
	var u models.Statement
	param := make(map[string]string)
	var res models.Response
	//param["keyword"] = c.Request.FormValue("keyword")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, err := u.ExportExcelSalesman(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

func StatementCustomer(c *gin.Context) {
	var u models.Statement
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["investment_id"] = c.Request.FormValue("investment_id")
	param["deptid"] = c.Request.FormValue("deptid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	//param["day"] = c.Request.FormValue("day")
	if param["start"] == "" || param["end"] == "" {
		re.Msg = "时间不能为空！"
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, total, err := u.StatementCustomerNew(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}
func StatementCustomerExport(c *gin.Context) {
	var u models.Statement
	param := make(map[string]string)
	var res models.Response
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		res.Msg = "时间不能为空！"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	//param["day"] = c.Request.FormValue("day")
	param["end"] += " 23:59:59"
	result, err := u.ExportExcelCustomer(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

//一次性分配结算（包括业务员，公司）
func StatementSettlement(c *gin.Context) {
	var u models.Summary
	//var iu models.Config
	var res models.Response
	id := c.Request.FormValue("id")
	//li := c.Request.FormValue("")
	if id == "" {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	err := u.AddStatementConfigOnce(id)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func StatementConfigOnce(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var re models.Response
	param["teamname"] = c.Request.FormValue("teamname")
	param["deptid"] = c.Request.FormValue("deptid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, total, err := u.StatementConfigOnce(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}

	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}

//未结算
func StatementNoSettlement(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var re models.Response
	//param["teamid"] = c.Request.FormValue("teamid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["deptid"] = c.Request.FormValue("deptid")
	param["keyword"] = c.Request.FormValue("keyword")
	param["id"] = c.Request.FormValue("id")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	//param["oncestatus"] = c.Request.FormValue("oncestatus")
	if param["start"] == "" || param["end"] == "" {
		re.Msg = "时间不能为空!"
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, total, err := u.StatementNoSettlement(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}

	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}

//已结算
func StatementIsSettlement(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var re models.Response
	//param["keyword"] = c.Request.FormValue("keyword")

	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	param["deptid"] = c.Request.FormValue("deptid")
	param["keyword"] = c.Request.FormValue("keyword")
	param["investmentid"] = c.Request.FormValue("investmentid")
	//param["oncestatus"] = c.Request.FormValue("oncestatus")
	if param["start"] == "" || param["end"] == "" {
		re.Msg = "时间不能为空!"
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, total, err := u.StatementIsSettlement(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}

	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}
func StatementConfigOnceExport(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var res models.Response
	//param["keyword"] = c.Request.FormValue("keyword")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	result, err := u.ExportExcelOnce(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}
func StatementSummary(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["deptid"] = c.Request.FormValue("deptid")
	param["id"] = c.Request.FormValue("id")
	//param["teamid"] = c.Request.FormValue("teamid")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["end"] == "" || param["start"] == "" {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	var result interface{}
	var total interface{}
	var err error
	result, total, err = u.StatementSummaryNew(param)
	if err != nil {
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal(param, result, total)

	c.JSON(http.StatusOK, res)
}

//func StatementAddOnce(c *gin.Context){
//	var u models.Summary
//	//param := make(map[string]string)
//	var res models.Response
//	//param["keyword"] = c.Request.FormValue("keyword")
//	id := c.Request.FormValue("id")
//
//	err:= u.AddStatementConfigOnce(id)
//	if err!=nil{
//		c.JSON(http.StatusOK, res.ReturnError(400))
//		return
//	}
//
//	c.JSON(http.StatusOK, res.ReturnOK())
//}
func StatementSummaryExport(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var res models.Response
	//param["keyword"] = c.Request.FormValue("keyword")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	if param["start"] == "" || param["end"] == "" {
		res.Msg = "时间不能为空"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	param["end"] += " 23:59:59"
	result, err := u.ExportExcelSummary(param)
	if err != nil {
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

func StatementCustomerSettle(c *gin.Context) {
	investmentID := c.Request.FormValue("investment_id")
	var u models.Statement
	err := u.CustomerSettle(investmentID)
	pkg.AssertErr(err, "结算失败", -1)
	var res models.Response
	c.JSON(http.StatusOK, res.ReturnOK())
}

func StatementSettleHistory(c *gin.Context) {
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["keyword"] = c.Request.FormValue("keyword")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	param["invest_id"] = c.Request.FormValue("invest_id")

	//param := map[string]string{
	//	"pageSize":  pageSize,
	//	"pageIndex": pageIndex,
	//	"keyword":   keyword,
	//}

	var history models.SettleCustomerHistory
	dataList, total, err := history.GetHistory(param)
	pkg.AssertErr(err, "获取已结算记录失败", -1)
	param["total"] = strconv.Itoa(total)

	res := utils.NewPageDataTotal(param, dataList, nil)
	c.JSON(http.StatusOK, res)
}
