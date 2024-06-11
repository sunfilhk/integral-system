package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"
)

func IndividualPerformance(c *gin.Context) {
	var data models.Performance
	//var re models.Response
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	//param["date"] = c.Request.FormValue("date")
	param["keyword"] = c.Request.FormValue("keyword")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	param["namenew"] = c.Request.FormValue("name")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//if param["date"]==""{
	//	c.JSON(http.StatusOK, re.ReturnError(400))
	//	return
	//}
	//if param["start"]==""||param["end"]==""{
	//	c.JSON(http.StatusOK, re.ReturnError(400))
	//	return
	//}
	param["end"] += "23:59:59"
	//switch utils.GetRolekey(c) {
	//case "finance":
	//case "boss":
	//case "admin":
	//default:
	//	param["userid"] = utils.GetUserIdStr(c)
	//}
	param["userid"] = utils.GetUserIdStr(c)
	result, ta, err := data.IndividualPerformance(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageDataTotal1(param, result, ta)
	c.JSON(http.StatusOK, res)
}

//新加查询个人的业绩
func Searchpersonal(c *gin.Context) {
	var data models.Performance
	//var res models.Response

	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize", "10")
	param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	param["userid"] = utils.GetUserIdStr(c)
	//param["deptid"] = c.Request.FormValue("dept_id")

	result, err := data.Personal(param)
	//if err !=nil{
	//	res.Msg = err.Error()
	//	c.JSON(http.StatusOK, res.ReturnError(0))
	//}
	//res.Data = result
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}
