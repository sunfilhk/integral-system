<template>
  <div class="customer">
    <!-- 头部搜索模块 -->
    <el-card>
      <el-form ref="queryForm" class="form" :model="queryParams" :inline="true">
        <el-form-item
          prop="keyword"
          label="业务名/客户名/手机号: "
          label-width="160px"
        >
          <el-input
            v-model="queryParams.keyword"
            clearable
            size="small"
            style="width: 240px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label="团队类型: " label-width="80px" prop="teamid">
          <el-select
            v-model="queryParams.teamid"
            clearable
            size="small"
            placeholder="请选择团队类型"
            style="height: 30px"
            @change="changeOptions"
          >
            <el-option
              v-for="item in teamOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>

      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button
            type="primary"
            icon="el-icon-search"
            size="mini"
            @click="handleQuery"
          >搜索</el-button>
          <el-button
            icon="el-icon-refresh"
            size="mini"
            @click="resetQuery"
          >重置</el-button>
          <el-button
            v-permisaction="['system:syscustomer:add']"
            type="primary"
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd"
          >新增</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 列表模块 -->
    <el-card class="footer">
      <el-table
        border
        :data="customerlist"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="序号"
          type="index"
          prop="index"
          show-overflow-tooltip
        >
          <template v-slot="{ $index }">
            {{
              (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
            }}
          </template>
        </el-table-column>
        <el-table-column label="客户姓名" prop="name" show-overflow-tooltip />
        <el-table-column label="手机号" width="130" prop="phone" />
        <el-table-column label="身份证号" width="180" prop="identity" />
        <el-table-column label="银行卡号" width="180" prop="bank" />
        <el-table-column label="开户行" prop="banknum" show-overflow-tooltip />
        <el-table-column label="性别" width="80" prop="sex" />
        <el-table-column
          label="业务员姓名"
          width="90"
          prop="nick_name"
          show-overflow-tooltip
          style="width: 120px"
        />
        <el-table-column label="团队类型" prop="teamname" />
        <el-table-column label="总投资金额" width="100">
          <template v-slot="{ row }">
            {{ row.totalamount ? row.totalamount : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="总利润" width="100">
          <template v-slot="{ row }">
            {{
              Number(row.totalprofit).toFixed(2) > 0
                ? Number(row.totalprofit).toFixed(2)
                : '-'
            }}
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100" prop="status">
          <template slot-scope="scope">
            <span v-show="scope.row.status == 0" class="a">已通过</span>
            <span v-show="scope.row.status == 3" class="b">已拒绝</span>
            <span v-show="scope.row.status == 1" class="c">审核中</span>
          </template>
        </el-table-column>
        <el-table-column label="客户创建时间" width="160" prop="create_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.create_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="150"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:syscustomer:lists']"
              size="mini"
              type="text"
              icon="el-icon-document"
              @click="inauto(scope.row)"
            >投资列表</el-button>
            <el-button
              v-permisaction="['system:syscustomer:edit']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              :disabled="scope.row.status == 1 || scope.row.flag == 0"
              @click="handleUpdate(scope.row)"
            >修改</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页符 -->
      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="queryParams.pageIndex"
        :limit.sync="queryParams.pageSize"
        @pagination="auto"
      />
    </el-card>

    <!-- <el-dialog title="上传" :visible.sync="port" width="500px">
          <el-upload
            :action="url"
            :headers="obj">
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
        </el-dialog> -->

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="form.name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="性别" prop="sex">
          <el-select
            v-model="form.sex"
            style="width: 100%"
            placeholder="请选择"
          >
            <el-option label="男" value="男" />
            <el-option label="女" value="女" />
          </el-select>
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input
            v-model="form.phone"
            placeholder="请输入手机号"
            maxlength="11"
          />
        </el-form-item>
        <el-form-item label="身份证" prop="identity">
          <el-input v-model="form.identity" placeholder="请输入身份证" />
        </el-form-item>
        <el-form-item label="银行卡号" prop="bank">
          <el-input
            v-model="form.bank"
            placeholder="请输入银行卡号"
            maxlength="19"
          />
        </el-form-item>
        <el-form-item label="开户行" prop="banknum">
          <el-input v-model="form.banknum" placeholder="请输入开户行" />
        </el-form-item>
        <el-form-item v-if="show4" label="金额" prop="amount">
          <el-input v-model="form.amount" placeholder="请输入金额" />
        </el-form-item>
        <el-form-item v-if="show4" label="投资时间" prop="date">
          <el-date-picker
            v-model="form.date"
            type="datetime"
            placeholder="选择日期时间"
          />
        </el-form-item>
        <el-form-item v-if="show4" label="利润比例" prop="profit">
          <el-input
            v-model="form.profit"
            placeholder="请输入利润比例"
            disabled
          />
        </el-form-item>
        <el-form-item v-if="show4" label="备注" prop="remark">
          <el-input
            v-model="form.remark"
            type="textarea"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 投资列表模块 -->
    <el-dialog title="投资列表" :visible.sync="show" width="1000px" class="Box">
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button
            v-permisaction="['system:syscustomer:adds']"
            type="primary"
            icon="el-icon-plus"
            size="mini"
            @click="inhandleAdd"
          >新增</el-button>
        </el-col>
      </el-row>

      <el-table
        :data="list"
        border
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
      >
        <el-table-column label="投资订单号" prop="id" />
        <el-table-column label="投资金额" prop="amount" />
        <el-table-column label="利润比例" prop="profit">
          <template slot-scope="scope">
            <span>{{ scope.row.profit * 100 + '%' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="备注" prop="remark" show-overflow-tooltip />
        <el-table-column label="状态" prop="status">
          <template slot-scope="scope">
            <span v-show="scope.row.status == 0" class="a">已通过</span>
            <span v-show="scope.row.status == 3" class="b">已拒绝</span>
            <span v-show="scope.row.status == 1" class="c">审核中</span>
            <span v-show="scope.row.status == 4" class="c">终止审核</span>
            <span v-show="scope.row.status == 5" class="b">已终止</span>
            <span v-show="scope.row.status == 6" class="b">已终止</span>
          </template>
        </el-table-column>
        <el-table-column label="投资时间" width="100" prop="invest_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.invest_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="订单创建时间" width="100" prop="create_time">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.create_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="200"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:syscustomer:edits']"
              size="mini"
              type="text"
              icon="el-icon-edit"
              :disabled="
                flag == 0 ||
                  scope.row.alterbool == 1 ||
                  scope.row.status == 1 ||
                  scope.row.status == 3 ||
                  scope.row.status == 4 ||
                  scope.row.status == 5 ||
                  scope.row.status == 6
              "
              @click="inhandleUpdate(scope.row)"
            >修改</el-button>
            <el-button
              v-permisaction="['system:syscustomer:stop']"
              size="mini"
              type="text"
              icon="el-icon-circle-close"
              :disabled="
                flag == 0 ||
                  scope.row.status == 1 ||
                  scope.row.status == 3 ||
                  scope.row.status == 4 ||
                  scope.row.status == 5 ||
                  scope.row.status == 6
              "
              @click="inhandleStop(scope.row)"
            >终止</el-button>
            <el-button
              v-permisaction="['system:syscustomer:revoke']"
              size="mini"
              type="text"
              icon="el-icon-refresh-left"
              :disabled="
                flag == 0 ||
                  scope.row.status == 1 ||
                  scope.row.status == 3 ||
                  scope.row.status == 4 ||
                  scope.row.status == 5 ||
                  scope.row.status == 6
              "
              @click="inhandleRevoke(scope.row)"
            >撤销</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!--分页符 -->
      <pagination
        v-show="totals > 0"
        :total="totals"
        :page.sync="params.pageIndex"
        :limit.sync="params.pageSize"
        @pagination="inauto"
      />

      <!-- 终止弹框 -->
      <el-dialog
        id="stop"
        :append-to-body="true"
        :visible.sync="stopShow"
        title="终止确认"
        width="400px"
      >
        <div class="txt">请确认终止方式？</div>
        <template #footer>
          <div style="text-align: center">
            <el-button type="info" @click="handStop(1)">手动终止</el-button>
            <el-button type="primary" @click="handStop(0)">到期终止</el-button>
          </div>
        </template>
      </el-dialog>
    </el-dialog>

    <!-- 投资列表弹框添加或修改角色配置对话框 -->
    <el-dialog :title="titles" :visible.sync="show1" width="500px">
      <el-form ref="forms" :model="forms" :rules="rule" label-width="80px">
        <el-form-item label="金额" prop="amount">
          <el-input v-model="forms.amount" placeholder="请输入金额" />
        </el-form-item>
        <el-form-item label="投资时间" prop="date">
          <el-date-picker
            v-model="forms.date"
            type="datetime"
            placeholder="选择日期时间"
            value-format="yyyy-MM-dd HH:mm"
          />
        </el-form-item>
        <el-form-item label="利润比例" prop="profit">
          <el-input
            v-model="forms.profit"
            placeholder="请输入利润比例"
            disabled
          />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input
            v-model="forms.remark"
            type="textarea"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="insubmitForm">确 定</el-button>
        <el-button @click="incancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  customerAdd,
  customerList,
  customerEdit,
  getProfitcustomer
} from '@/api/customer/customer'
import {
  investmentAdd,
  investmentList,
  investmentEdit,
  investmentBreak,
  investmentRevoke
} from '@/api/investment/investment'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
// import { getToken } from '@/utils/auth'
export default {
  name: 'Customer',
  components: {},
  props: {},
  data() {
    return {
      // 控制终止弹框
      stopShow: false,
      // 终止参数
      stopParams: {
        customerid: '',
        id: '',
        breaktype: 0 // 终止类型,传0为到期终止,传1为手动终止
      },
      // 团队选项
      teamOptions: [],
      flag: '',
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 控制主页新增弹框显示
      show4: false,
      ids: [],
      open: false,
      title: '',
      form: {},
      // 客户管理列表
      customerlist: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined,
        teamid: null // 部门id
      },
      total: 0,
      // 表单校验
      rules: {
        name: [{ required: true, message: '姓名不能为空', trigger: 'blur' }],
        sex: [{ required: true, message: '性别不能为空', trigger: 'blur' }],
        identity: [
          // { required: true, message: '身份证不能为空', trigger: 'blur' },
          // {
          //   pattern: /^[1-9]\d{5}(18|19|20|(3\d))\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$/,
          //   message: '请输入正确的身份证号',
          //   trigger: 'blur'
          // }
        ],
        bank: [
          { required: true, message: '银行卡号不能为空', trigger: 'blur' },
          {
            pattern: /^([1-9]{1})(\d{15}|\d{17}|\d{18})$/,
            message: '请输入正确的银行卡号',
            trigger: 'blur'
          }
        ],
        banknum: [
          { required: true, message: '开户行不能为空', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '手机号码不能为空', trigger: 'blur' },
          {
            pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
            message: '请输入正确的手机号码',
            trigger: 'blur'
          }
        ],
        amount: [
          { required: true, message: '金额不能为空', trigger: 'blur' },
          {
            pattern: /^[0-9]*[1-9][0-9]*$/,
            message: '金额为数字且不能为小数',
            trigger: 'blur'
          }
        ],
        date: [{ required: true, message: '时间不能为空', trigger: 'blur' }]
      },
      // 控制投资列表弹框
      show: false,
      show1: false,
      titles: '',
      forms: {},
      list: [],
      params: {
        pageIndex: 1,
        pageSize: 10
      },
      totals: 0,
      // 表单校验
      rule: {
        amount: [
          { required: true, message: '金额不能为空', trigger: 'blur' },
          {
            pattern: /^[0-9]*[1-9][0-9]*$/,
            message: '金额为数字且不能为小数',
            trigger: 'blur'
          }
        ],
        date: [{ required: true, message: '时间不能为空', trigger: 'blur' }]
      }
    }
  },
  computed: {},
  watch: {},
  created() {
    this.auto()
    this.getData()
  },
  mounted() {},
  methods: {
    // 点击团队选项,
    changeOptions(e) {
      this.queryParams.deptId = e
      this.auto()
    },
    // 获取团队列表
    async getData() {
      const res = await teamList()
      this.teamOptions = res.list.map(v => {
        return {
          value: v.deptId, // 部门id
          label: v.deptname // 部门名称
        }
      })
    },
    // 表格数据
    auto() {
      // this.url=process.env.VUE_APP_BASE_API+'/api/v1/investmentImport'
      // this.obj.Authorization='Bearer '+getToken()
      customerList(this.queryParams).then(res => {
        this.customerlist = res.list
        this.total = res.total
      })
    },
    // 清空表单
    reset() {
      this.form = {
        name: undefined,
        sex: undefined,
        identity: undefined,
        bank: undefined,
        banknum: undefined,
        phone: undefined,
        amount: undefined,
        remark: undefined
      }
      this.resetForm('form')
    },

    // 点击搜索页面新增
    handleAdd() {
      // this.reset()
      // 显示主页新增弹框
      this.show4 = true
      getProfitcustomer({ configKey: 'customerratio' }).then(res => {
        this.form.profit = parseFloat(res.data.profit).toFixed(5) * 100 + '%'
        // 显示弹框
        this.open = true
        this.title = '添加客户'
      })
    },
    // 点击修改
    handleUpdate(row) {
      this.auto()
      this.show4 = false
      this.form = row
      this.open = true
      this.title = '修改客户'
    },
    // 点击提交
    submitForm() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            customerEdit(this.form).then(res => {
              this.msgSuccess('修改成功，待审核')
              this.open = false
              this.auto()
            })
          } else {
            this.form.profit =
              parseFloat(this.form.profit.replace('%', '')).toFixed(3) * 0.01
            this.form.date = this.parseTime(this.form.date)
            customerAdd(this.form).then(res => {
              this.msgSuccess('新增成功')
              this.open = false
              this.auto()
            })
          }
        }
      })
    },

    // 点击弹框取消
    cancel() {
      // 关闭弹框
      this.open = false
      // 清空表单
      this.reset()
    },
    // 点击搜索
    handleQuery() {
      this.queryParams.pageIndex = 1
      // 刷新数据
      this.auto(this.queryParams)
    },
    // 点击页面重置
    resetQuery() {
      // 清空表单
      this.resetForm('queryForm')
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined,
        deptId: null // 部门id
      }
      // 刷新数据
      this.handleQuery()
    },
    inauto(row) {
      this.show = true
      // alert(this.params.customerid)
      if (row == undefined) {
        if (localStorage.getItem('customerid') != undefined) {
          this.params.customerid = localStorage.getItem('customerid')
        }
      } else {
        // debugger
        if (row.id != undefined) {
          this.params.customerid = row.id
          localStorage.setItem('customerid', row.id)
        }
        localStorage.setItem('profit', row.profit)
        this.flag = row.flag
      }
      // 投资列表数据
      investmentList(this.params).then(res => {
        this.list = res.list
        this.totals = res.total
      })
    },
    // 重置
    inreset() {
      this.forms = {
        amount: undefined,
        remark: undefined
      }
      this.resetForm('forms')
    },
    // 点击新增
    inhandleAdd() {
      this.inreset()
      getProfitcustomer({ configKey: 'customerratio' }).then(res => {
        this.forms.profit = parseFloat(res.data.profit).toFixed(5) * 100 + '%'
        this.show1 = true
        this.titles = '添加投资'
      })
    },
    inhandleUpdate(row) {
      this.inauto()
      this.forms = row
      this.forms.profit = this.forms.profit * 100 + '%'
      this.show1 = true
      this.titles = '修改投资'
    },

    inhandleStop(row) {
      // 显示弹框
      this.stopShow = true
      this.stopParams = {
        customerid: row.customerid,
        id: row.id,
        breaktype: 0 // 终止类型,传0为到期终止,传1为手动终止
      }
    },

    // 终止确认
    handStop(num) {
      // 到期
      if (num == 0) {
        investmentBreak(this.stopParams).then(res => {
          if (res.code == 200) {
            this.msgSuccess('操作成功')
            // 关闭弹框
            this.stopShow = false
            // 刷新数据
            this.inauto()
          }
        })
      } else {
        this.stopParams.breaktype = 1
        investmentBreak(this.stopParams).then(res => {
          if (res.code == 200) {
            this.msgSuccess('操作成功')

            // 关闭弹框
            this.stopShow = false
            // 刷新数据
            this.inauto()
          }
        })
      }
    },
    inhandleRevoke(row) {
      var that = this
      var params = {
        customerid: localStorage.getItem('customerid'),
        id: row.id
      }
      this.$confirm('是否确认撤销这笔投资?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(function() {
          investmentRevoke(params).then(res => {
            that.msgSuccess('操作成功')
            that.inauto()
          })
        })
        .catch(() => {
          // console.log("取消")
        })
    },
    insubmitForm() {
      if (this.forms.amount >= 1000000000) {
        this.msgError('金额需大于0')
        return
      }
      this.$refs['forms'].validate(valid => {
        if (valid) {
          this.forms.profit =
            parseFloat(this.forms.profit.replace('%', '')).toFixed(3) * 0.01
          this.forms.date = this.parseTime(this.forms.date)
          console.log(this.forms.date)
          if (this.forms.id !== undefined) {
            investmentEdit(this.forms).then(res => {
              this.msgSuccess('修改成功，待审核')
              this.show1 = false
              this.inauto()
            })
          } else {
            this.forms.customerid = localStorage.getItem('customerid')
            investmentAdd(this.forms).then(res => {
              this.msgSuccess('新增成功')
              this.show1 = false
              this.inauto()
            })
          }
        }
      })
    },
    // 点击投资列表内取消
    incancel() {
      this.show1 = false
      this.inreset()
    },
    // 双击复制
    handleDblClick(row, cloumn, cell, event) {
      if (cloumn.property) {
        const txt = event.target.innerText
        this.$copyText(txt).then(res => {
          this.$message.success('复制成功')
        })
      }
    }
  }
}
</script>

<style scoped lang="scss">
.customer {
  padding: 20px;
}
.form {
  min-width: 1366px;
}
.footer {
  margin-top: 20px;
}
.a {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(170, 255, 170);
  color: rgb(0, 255, 98);
  background-color: rgb(235, 255, 235);
}
.b {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(255, 180, 180);
  color: rgb(255, 80, 80);
  background-color: rgb(255, 221, 221);
}
.c {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(149, 195, 255);
  color: rgb(0, 110, 255);
  background-color: rgb(231, 241, 255);
}

/* 终止弹框 */
::v-deep #stop .el-dialog .el-dialog__header {
  padding: 20px !important;
  padding-bottom: 10px !important;
}

::v-deep #stop .el-dialog .el-dialog__body {
  padding: 0px !important;
}
.txt {
  font-size: 20px;
  text-align: center;
  margin-bottom: 15px;
}
</style>
