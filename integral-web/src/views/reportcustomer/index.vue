<template>
  <div class="reportcustomer">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" class="form" label-width="110px" :model="queryParams">
        <!--投资日期 -->
        <el-form-item prop="update_time" label="结算日期: ">
          <el-date-picker
            v-model="date"
            size="small"
            type="daterange"
            align="right"
            unlink-panels
            :clearable="false"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :picker-options="pickerOptions"
            :default-value="Date.now() - 30 * 24 * 3600 * 1000"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="changeDate"
          />
        </el-form-item>
        <!-- 订单号 -->
        <el-form-item prop="investment_id" label="投资订单号: ">
          <el-input
            v-model="queryParams.investment_id"
            size="small"
            clearable
            @keyup.enter.native="auto"
          />
        </el-form-item>
        <!-- 客户名/手机号-->
        <el-form-item prop="keyword" label="客户名/手机号: ">
          <el-input
            v-model="queryParams.keyword"
            size="small"
            @keyup.enter.native="auto"
          />
        </el-form-item>
        <!-- 团队类型 -->
        <el-form-item prop="dept_id" label="团队类型: ">
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
      <!-- 按钮 -->
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5" style="margin: 10px 0 0 30px">
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
            v-permisaction="['system:sysreportcustomer:export']"
            type="warning"
            icon="el-icon-download"
            size="mini"
            @click="handleExport"
          >导出</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 列表模块 -->
    <el-card>
      <el-table
        border
        :data="list"
        show-summary
        :summary-method="getSummaries"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="投资订单号"
          prop="investment_id"
          align="center"
          show-overflow-tooltip
          width="170"
        />
        <el-table-column
          label="订单创建时间"
          prop="ordertime"
          align="center"
          show-overflow-tooltip
          width="170"
        >
          <template v-slot="{ row }">
            {{ parseTime(row.ordertime) }}
          </template>
        </el-table-column>
        <el-table-column
          label="姓名"
          prop="name"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="手机号"
          width="120"
          prop="phone"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="银行卡号"
          width="190"
          prop="bank"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="开户行"
          prop="banknum"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column label="性别" width="60" prop="sex" align="center" />
        <el-table-column
          label="分润比例"
          prop="profit"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ scope.row.profit * 100 + '%' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="业绩"
          prop="amount"
          align="center"
          width="150px"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="利润"
          prop="profits"
          align="center"
          width="150px"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.profits) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="业务员"
          prop="nick_name"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="团队类型"
          prop="teamname"
          align="center"
          show-overflow-tooltip
        />
        <!-- 投资时间 -->
        <el-table-column
          label="投资时间"
          width="155"
          align="center"
          prop="invest_time"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.invest_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="结算时间"
          width="155"
          align="center"
          prop="update_time"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.update_time) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="合同到期时间"
          width="155"
          align="center"
          prop="expiration_date"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.expiration_date) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          width="80"
          class-name="small-padding fixed-width"
        >
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:sysreportcustomer:set']"
              size="mini"
              type="text"
              icon="el-icon-document"
              @click="settlement(scope.row)"
            >结算</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="queryParams.pageIndex"
        :limit.sync="queryParams.pageSize"
        @pagination="auto"
      />
    </el-card>
  </div>
</template>

<script>
import {
  statementCustomer,
  statementCustomerExport,
  statementCustomerSettle
} from '@/api/report/report'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'Reportcustomer',
  components: {},
  props: {},
  data() {
    return {
      // 限制时间选择截止当天
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      },
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 投资开始时间
        end: '', // 投资截止时间
        keyword: undefined, // 客户姓名/手机号
        deptid: null, // 团队id
        investment_id: null // 投资单号
      },
      // 结算日期
      date: [],
      // 团队选项
      teamOptions: [],
      // 列表
      list: [],
      total: 0,
      summation: {}
    }
  },
  computed: {},
  watch: {},
  created() {
    const start = new Date()
    const end = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
    this.queryParams.start = this.parseTime(start).slice(0, 10)
    this.queryParams.end = this.parseTime(end).slice(0, 10)
    // 初始化投资日期
    this.date = [this.queryParams.start, this.queryParams.end]
    this.auto()
    this.getData()
  },
  mounted() {},
  methods: {
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
      statementCustomer(this.queryParams).then(res => {
        this.list = res.list
        this.total = res.total
        this.summation = res.summation
      })
    },
    // 点击搜索
    searchClick() {
      this.auto()
    },
    // 重置
    resetClick() {
      // 日期复原
      const start = new Date()
      const end = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      this.queryParams.start = this.parseTime(start).slice(0, 10)
      this.queryParams.end = this.parseTime(end).slice(0, 10)
      this.date = [this.queryParams.start, this.queryParams.end]
      // 清空表单
      this.$refs.form.resetFields()
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        start: this.date[0], // 投资开始时间
        end: this.date[1], // 投资截止时间
        keyword: undefined, // 客户姓名/手机号
        deptId: null, // 团队id
        id: null // 投资单号
      }
      // 刷新
      this.auto()
    },
    // 选择日期
    changeDate(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      // 刷新数据
      this.auto()
    },
    select() {
      this.auto()
    },
    // 点击结算
    settlement(row) {
      this.$confirm('是否结算?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(function() {
          return statementCustomerSettle({ investment_id: row.investment_id })
        })
        .then(() => {
          this.msgSuccess('结算成功')
          this.auto()
        })
    },
    // 导出
    handleExport() {
      const queryParams = this.queryParams
      this.$confirm('是否确认导出所有数据?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        statementCustomerExport({
          start: queryParams.start,
          end: queryParams.end
        }).then(res => {
          location.href = process.env.VUE_APP_BASE_API + '/' + res.data
        })
      })
    },
    // 合计
    getSummaries(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
          return
        }
        const values = data.map(item => Number(item[column.property]))
        if (!values.every(value => isNaN(value))) {
          sums[index] = this.moneyFormat(this.summation[column.property])
        } else {
          sums[index] = ''
        }
      })
      return sums
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
.reportcustomer {
  padding: 20px;
}
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
</style>
