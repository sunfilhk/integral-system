package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"xAdmin/models"
	"xAdmin/pkg"
	"xAdmin/utils"
)

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/deptList [get]
// @Security
func GetDeptList(c *gin.Context) {
	var Dept models.Dept
	Dept.Deptname = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.Deptid, _ = utils.StringToInt64(c.Request.FormValue("deptId"))
	//Dept.DataScope = utils.GetUserIdStr(c)
	result, err := Dept.SetDept(true)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetDeptTree(c *gin.Context) {
	var Dept models.Dept
	Dept.Deptname = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.Deptid, _ = utils.StringToInt64(c.Request.FormValue("deptId"))
	//Dept.DeptPath = c.Request.FormValue("deptath")

	//Dept.TeamId, _ = utils.StringToInt64(c.Request.FormValue("teamid")) //新加
	result, err := Dept.SetDept(false)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

//获取客户分润列表
func Profitcustomers(c *gin.Context) {
	var u models.Dept
	var err error
	param := make(map[string]string)
	//param["keyword"] = c.Request.FormValue("keyword")
	//param["teamid"] = c.Request.FormValue("teamid")
	//param["deptId"] = c.Request.FormValue("deptId")//profit_type
	param["deptId"] = c.Request.FormValue("deptId") //profit_type
	//param["profit"] = c.Request.FormValue("profit")//profit_type

	//param["deptid"] = c.Request.FormValue("dept_id")  //新加字段

	var res models.Response
	result, err := u.Profitcusto(param)
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	if err != nil {
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

//新增时获取一次性团队
func Getonceinto(c *gin.Context) {
	var data models.Dept
	//var res models.Response

	param := make(map[string]string)
	param["deptId"] = c.Request.FormValue("deptId")
	//param["pageSize"] = c.DefaultQuery("pageSize", "10")
	//param["pageIndex"] = c.DefaultQuery("pageIndex", "1")
	//param["deptid"] = c.Request.FormValue("dept_id")

	result, err := data.Gotonce(param)
	//if err !=nil{
	//	res.Msg = err.Error()
	//	c.JSON(http.StatusOK, res.ReturnError(0))
	//}
	//res.Data = result
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)
}

//查询团队分润比例
func GetKey(c *gin.Context) {
	var data models.Dept
	param := make(map[string]string)
	param["userid"] = utils.GetUserIdStr(c)
	//param["profittype"]= c.Request.FormValue("profit_type")
	der := param["userid"]
	var res models.Response
	var err error
	res.Data, err = data.Getteam(der) //原来的
	//err = data.Getdeptid(param)
	if err != nil {
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

//新修改
func UpdateProfitscustomer(c *gin.Context) {
	var data models.Dept
	//param := make(map[string]string)
	//param["dept_name"] = c.Request.FormValue("deptname")
	profits := c.Request.FormValue("profit")
	deptIds := c.Request.FormValue("deptId")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.Updatecustomer(profits, deptIds)
	//if err!=nil{
	//	res.Code = 0
	//	res.Msg = err.Error()
	//	c.JSON(http.StatusOK, res)
	//}
	pkg.AssertErr(err, "", -1)
	c.JSON(http.StatusOK, res.ReturnOK())
}

//查询客户团队
func Getcustomerteam(c *gin.Context) {
	var fifa models.Dept
	param := make(map[string]string)
	param["dept_name"] = c.Request.FormValue("dept_name")
	result, err := fifa.Agetcusto(param)
	//if err !=nil{
	//	res.Msg = err.Error()
	//	c.JSON(http.StatusOK, res.ReturnError(0))
	//}
	//res.Data = result
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param, result)

	c.JSON(http.StatusOK, res)

}

// @Summary 部门列表数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security
func GetDept(c *gin.Context) {
	var Dept models.Dept
	Dept.Deptid, _ = strconv.ParseInt(c.Param("deptId"), 10, 64)
	//Dept.DataScope = utils.GetUserIdStr(c)
	result, err := Dept.GetDeptInfo()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body models.Dept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func InsertDept(c *gin.Context) {
	var data models.Dept
	err := c.BindWith(&data, binding.JSON)
	pkg.AssertErr(err, "", 500)
	data.CreateBy = utils.GetUserIdStr(c)
	result, err := data.Create()
	pkg.AssertErr(err, "", -1)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body models.Dept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [put]
// @Security Bearer
func UpdateDept(c *gin.Context) {
	var data models.Dept
	err := c.BindWith(&data, binding.JSON)
	pkg.AssertErr(err, "", -1)
	data.UpdateBy = utils.GetUserIdStr(c)
	result, err := data.Update(data.Deptid)
	pkg.AssertErr(err, "", -1)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dept/{id} [delete]
func DeleteDept(c *gin.Context) {
	var data models.Dept
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err = data.Delete(id)
	pkg.AssertErr(err, "删除失败", 500)

	var res models.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetDeptTreeRoleselect(c *gin.Context) {
	var Dept models.Dept
	var SysRole models.SysRole
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	SysRole.Id = id
	result, err := Dept.SetDeptLable()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	menuIds := make([]int64, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleDeptId()
		pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"depts":       result,
		"checkedKeys": menuIds,
	})
}
