<template>
  <div class="business">
    <!-- 搜索区域 -->
    <el-card class="topCard">
      <el-form ref="form" class="form" label-width="120px" :model="queryParams">
        <!--投资日期 -->
        <el-form-item label="投资日期: ">
          <el-date-picker
            v-model="date"
            :clearable="true"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="changeDate"
          />
        </el-form-item>
        <!-- 订单号 -->
        <el-form-item prop="id" label="投资单号: ">
          <el-input v-model="queryParams.id" size="small" />
        </el-form-item>
        <!-- 创建日期 -->
        <el-form-item label="订单创建时间: ">
          <el-date-picker
            v-model="createDate"
            :clearable="true"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="changeCreateDate"
          />
        </el-form-item>
        <!-- 状态 -->
        <el-form-item prop="status" label="订单状态: ">
          <el-select v-model="queryParams.status" clearable size="small">
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <!-- 客户名/手机号-->
        <el-form-item prop="keyword" label="客户名/手机号: ">
          <el-input v-model="queryParams.keyword" size="small" />
        </el-form-item>
        <!-- 业务员 -->
        <el-form-item prop="nick_name" label="业务员: ">
          <el-input v-model="queryParams.nick_name" size="small" />
        </el-form-item>
        <!-- 所属团队 -->
        <el-form-item prop="deptid" label="团队类型: ">
          <el-select v-model="queryParams.deptid" clearable size="small">
            <el-option
              v-for="item in teamOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <el-row>
        <el-col :span="24" style="margin: 10px 0 0 40px">
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-search"
            @click="searchClick"
          >搜索</el-button>
          <el-button
            size="mini"
            icon="el-icon-refresh"
            @click="resetClick"
          >重置</el-button>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-plus"
            @click="addClick"
          >新增订单</el-button>
          <el-button
            size="mini"
            icon="el-icon-download"
            type="warning"
            @click="exportExcel"
          >导出</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 表格区域 -->
    <el-card>
      <el-table
        width="100%"
        :data="businessList"
        border
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="序号"
          width="50"
          type="index"
          :show-overflow-tooltip="true"
        >
          <template v-slot="{ $index }">
            {{
              (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
            }}
          </template>
        </el-table-column>
        <el-table-column
          prop="id"
          label="投资订单号"
          width="170"
          :show-overflow-tooltip="true"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="name"
          label="客户姓名"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="phone"
          label="手机号"
          width="115"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="id_car"
          label="身份证号"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="bank_num"
          label="银行卡号"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="amount"
          label="投资金额"
          width="100"
        >
          <template v-slot="{ row }">
            {{ row.amount > 0 ? row.amount : '-' }}
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="profit"
          label="分润比例"
        >
          <template v-slot="{ row }">
            <span>{{ parseFloat(row.profit * 100).toFixed(2) }}%</span>
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="profits"
          label="分润金额"
          width="88"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="teamname"
          label="团队类型"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="invest_time"
          label="投资时间"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.invest_time) }}
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="expiration_date"
          label="合同到期时间"
          width="100"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.expiration_date) }}
          </template>
        </el-table-column>
        <el-table-column :show-overflow-tooltip="true" prop="form" label="类型">
          <template v-slot="{ row }">
            <div v-if="row.form === 0">客户分润</div>
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="remark"
          label="备注"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="nick_name"
          label="业务员"
          width="90"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="status"
          label="状态"
          width="120"
        >
          <template v-slot="{ row }">
            <span v-show="row.status == 0" class="a">已通过</span>
            <span v-show="row.status == 3" class="b">已拒绝</span>
            <span v-show="row.status == 1" class="c">审核中</span>
            <span v-show="row.status == 4" class="c">已终止</span>
            <span v-show="row.status == 5" class="b">已终止</span>
            <span v-show="row.status == 6" class="b">已终止</span>
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="order_time"
          label="订单创建时间"
          width="100"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.order_time) }}
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="investment_id"
          label="操作"
          width="150"
        >
          <template v-slot="{ row }">
            <div v-if="row.status !== '0'">
              <el-button
                size="mini"
                type="text"
                disabled
                style="color: gray"
              >无需操作</el-button>
            </div>
            <div v-else>
              <el-button
                :disabled="row.alterbool == 1 ? true : false"
                size="mini"
                icon="el-icon-edit"
                type="text"
                @click="edit(row)"
              >
                修改
              </el-button>
              <el-button
                v-permisaction="['system:syscustomer:stop']"
                size="mini"
                type="text"
                @click="terminationClick(row)"
              >终止</el-button>
              <el-button
                v-permisaction="['system:syscustomer:revoke']"
                size="mini"
                type="text"
                @click="revocation(row)"
              >撤销</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页符 -->
      <div>
        <div class="page">
          <pagination
            v-show="total > 0"
            :total="total"
            :page.sync="queryParams.pageIndex"
            :limit.sync="queryParams.pageSize"
            @pagination="clickPage"
          />
        </div>
      </div>
    </el-card>

    <!-- 新增/修改弹框 -->
    <Add ref="add" @getList="getList" />

    <!-- 终止弹框 -->
    <el-dialog
      :visible.sync="stopShow"
      title="终止确认"
      width="400px"
      class="stop"
    >
      <div class="txt">请确认终止方式？</div>
      <template #footer>
        <div style="text-align: center">
          <el-button type="info" @click="handStop(1)">手动终止</el-button>
          <el-button type="primary" @click="handStop(0)">到期终止</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  investListss,
  investmentBreak,
  investmentRevoke
} from '@/api/customer/customer'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
import Add from './components/add.vue'
export default {
  name: 'Business',
  components: { Add },
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
      date: '', // 投资日期
      createDate: '', // 订单创建日期
      // 团队选项
      teamOptions: [],
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      },
      // 控制新增弹框显示
      show: false,
      queryParams: {
        pageIndex: 1, // 页码
        pageSize: 10, // 页容量
        start: '', // 投资开始时间
        end: '', // 投资结束时间
        createStart: '', // 创建开始
        createEnd: '', // 创建截止
        keyword: undefined, // 客户/手机号
        deptid: '', // 团队类型
        id: '', // 投资单号
        nick_name: '', // 业务员
        status: '' // 订单状态
      },
      // 总数
      total: 0,
      // 业务管理列表
      businessList: [],

      // 状态
      status: '1',
      // 状态选项
      statusOptions: [
        {
          value: '1',
          label: '审核中'
        },
        {
          value: '3',
          label: '已拒绝'
        },
        {
          value: '0',
          label: '已通过'
        },
        {
          value: '4',
          label: '已终止'
        }
      ]
      // // 合同到期
      // contractTime: '1',
      // // 合同到期选项
      // contractOptions: [
      //   {
      //     value: '1',
      //     label: '当天到期'
      //   },
      //   {
      //     value: '7',
      //     label: '7天内到期'
      //   },
      //   {
      //     value: '14',
      //     label: '14天内到期'
      //   },
      //   {
      //     value: '30',
      //     label: '30天内到期'
      //   }
      // ]
    }
  },
  mounted() {},
  created() {
    // 设置日期选择器初始化显示一个月时间
    const start = new Date()
    const end = new Date()
    // 投资初始化日期
    start.setTime(start.getTime() - 1 * 3600 * 1000 * 24 * 30) // 乘2就两个月,以此类推
    this.queryParams.start = this.parseTime(start).slice(0, 10)
    this.queryParams.end = this.parseTime(end).slice(0, 10)
    this.date = [this.queryParams.start, this.queryParams.end]

    // 订单创建初始化日期
    // this.queryParams.createStart = this.parseTime(start).slice(0, 10)
    // this.queryParams.createEnd = this.parseTime(end).slice(0, 10)
    // this.createDate = [this.queryParams.createStart, this.queryParams.createEnd]

    this.getData(this.queryParams)
    this.getList()
  },
  methods: {
    // 点击撤销
    async revocation(row) {
      this.$confirm('是否确认撤销这笔投资?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          await investmentRevoke({ customerid: row.customerid, id: row.id })
          this.$message.success('撤销成功')
          this.getList()
        })
        .catch(() => {})
    },
    // 点击终止
    terminationClick(row) {
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
            this.getList()
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
            this.getList()
          }
        })
      }
    },
    // 获取列表
    async getList() {
      const res = await investListss(this.queryParams)
      if (res.list.length > 0) {
        this.businessList = res.list
        this.total = res.total
      } else {
        this.businessList = []
        this.total = 0
      }
    },
    // 改变投资日期
    changeDate(e) {
      this.queryParams.start = this.date != null ? this.date[0] : ''
      this.queryParams.end = this.date != null ? this.date[1] : ''
    },

    // 改变订单创建日期
    changeCreateDate(e) {
      this.queryParams.createStart =
        this.createDate != null ? this.createDate[0] : ''
      this.queryParams.createEnd =
        this.createDate != null ? this.createDate[1] : ''
    },

    // 改变团队选项
    changeOptions(e) {
      this.queryParams.deptid = e // 部门id
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

    // 点击搜索
    searchClick() {
      this.queryParams.start = this.date != null ? this.date[0] : ''
      this.queryParams.end = this.date != null ? this.date[1] : ''

      this.queryParams.createStart =
        this.createDate != null ? this.createDate[0] : ''
      this.queryParams.createEnd =
        this.createDate != null ? this.createDate[1] : ''

      this.getList(this.queryParams)
    },
    // 点击分页
    clickPage() {
      this.getList()
    },
    // 双击复制
    handleDblClick(row, cloumn, cell, event) {
      if (cloumn.property) {
        const txt = event.target.innerText
        this.$copyText(txt).then(res => {
          this.$message.success('复制成功')
        })
      }
    },
    // 点击导出
    async exportExcel() {
      const obj = await import('@/vendor/Export2Excel') // 点击时候才导入使用
      obj.export_json_to_excel({
        header: ['投资订单号', '姓名'],
        data: [
          [33, 'Tom'],
          [666, 'bob']
        ]
      })
    },
    // 点击新增
    addClick() {
      // 显示弹框
      this.$refs.add.isShow = true
      this.$refs.add.mode = 'add'
    },
    // 点击修改
    edit(row) {
      // 显示弹框
      this.$refs.add.isShow = true
      this.$refs.add.mode = 'edit'
      this.form = JSON.parse(JSON.stringify(row))
      this.form.profit = (this.form.profit * 100).toFixed(2)
      this.form.date = this.parseTime(row.invest_time)
      // 回显数据
      this.$refs.add.form = JSON.parse(JSON.stringify(this.form))
    },
    // 重置事件
    resetClick() {
      // 清空表单
      this.$refs.form.resetFields()
      // 手动再重置一次
      this.form = {
        pageIndex: 1, // 页码
        pageSize: 10, // 页容量
        start: '', // 投资开始时间
        end: '', // 投资结束时间
        createStart: '', // 创建开始
        createEnd: '', // 创建截止
        keyword: undefined, // 客户/手机号
        deptid: '', // 团队类型
        id: '', // 投资单号
        nick_name: '', // 业务员
        status: '' // 订单状态
      }

      // 设置日期选择器初始化显示一个月时间
      const start = new Date()
      const end = new Date()
      // 投资初始化日期
      start.setTime(start.getTime() - 1 * 3600 * 1000 * 24 * 30)
      this.queryParams.start = this.parseTime(start).slice(0, 10)
      this.queryParams.end = this.parseTime(end).slice(0, 10)
      this.date = [this.queryParams.start, this.queryParams.end]

      // / 订单创建日期
      this.createDate = ''

      // 订单创建初始化日期
      // this.createDate = [
      //   this.queryParams.createStart,
      //   this.queryParams.createEnd
      // ]

      this.queryParams.createStart = ''
      this.queryParams.createEnd = ''

      // 刷新
      this.getList()
    }
  }
}
</script>
<style scoped lang="scss">
.business {
  padding: 20px;
  .topCard {
    margin-bottom: 20px;
    .form {
      min-width: 1366px;
      display: flex;
      flex-wrap: wrap;
      .el-form-item {
        margin: 0 10px 5px 0;
        ::v-deep .el-input--small .el-input__inner {
          width: 240px;
        }
      }
    }
  }
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

// 终止弹框
::v-deep .stop .el-dialog__header {
  padding: 20px;
  padding-bottom: 10px;
}

::v-deep .stop .el-dialog__body {
  padding: 0;
}
.txt {
  font-size: 20px;
  text-align: center;
  margin-bottom: 15px;
}
</style>
