<template>
  <div class="history">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" :model="queryParams" :inline="true" class="form">
        <el-form-item label="投资日期: ">
          <el-date-picker
            v-model="value1"
            type="daterange"
            :clearable="false"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :picker-options="pickerOptions"
            :default-value="Date.now() - 30 * 24 * 3600 * 1000"
            value-format="yyyy-MM-dd"
            size="small"
            style="width: 240px"
            @change="select"
          />
        </el-form-item>
        <el-form-item prop="id" label="投资单号: ">
          <el-input
            v-model="queryParams.id"
            clearable
            size="small"
            style="width: 240px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item prop="keyword" label="客户名/业务员/手机号: ">
          <el-input
            v-model="queryParams.keyword"
            clearable
            size="small"
            style="width: 240px"
            @keyup.enter.native="handleQuery"
          />
        </el-form-item>
        <el-form-item label="团队类型: ">
          <el-select
            v-model="queryParams.deptid"
            clearable
            size="small"
            style="margin-left: 10px; height: 30px"
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
      <!-- 按钮 -->
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5" style="margin: 10px 0 0 0">
          <el-button
            type="primary"
            icon="el-icon-search"
            size="mini"
            @click="handleQuery"
          >搜索</el-button>
          <el-button
            icon="el-icon-refresh"
            size="mini"
            @click="resetClick"
          >重置</el-button>
          <el-button
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
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        :data="profit"
        show-summary
        :summary-method="getSummaries"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          label="投资订单"
          prop="id"
          align="center"
          show-overflow-tooltip
          width="170"
        />
        <el-table-column
          label="客户姓名"
          prop="customername"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="业务员姓名"
          prop="nick_name"
          align="center"
          show-overflow-tooltip
          width="100"
        />
        <el-table-column
          label="团队类型"
          prop="teamname"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="销售"
          prop="amount"
          align="center"
          show-overflow-tooltip
          width="140"
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          v-for="(item, index) in keyslist"
          :key="index"
          :label="item"
          :prop="wordslist[index]"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="业务部门"
          prop="company"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.company) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="业务员"
          prop="salesman"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.salesman) }}</span>
          </template>
        </el-table-column>
        <!-- <el-table-column label="客户" prop="customerprofit" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.customerprofit) }}</span>
            </template>
          </el-table-column> -->
        <el-table-column
          label="备注"
          prop="remark"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="投资时间"
          width="160"
          align="center"
          prop="invest_time"
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.invest_time) }}</span>
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
import { statementSummary, statementSummaryExport } from '@/api/report/report'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'History',
  components: {},
  props: {},
  data() {
    return {
      total: 0,
      summation: {},
      profit: [],
      keyslist: '',
      wordslist: '',
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: null, // 开始时间
        end: null, // 截止时间
        id: '', // 投资单号
        keyword: '', // 客户/业务/手机号
        deptid: '' // 团队id
      },
      // 日期
      value1: '',
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      },
      // 团队选项
      teamOptions: []
    }
  },
  computed: {},
  watch: {},
  created() {
    const end = new Date()
    const start = new Date()
    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
    this.queryParams.start = this.parseTime(start).slice(0, 10)
    this.queryParams.end = this.parseTime(end).slice(0, 10)
    this.value1 = [this.queryParams.start, this.queryParams.end]
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
      statementSummary(this.queryParams).then(res => {
        console.log('statementSummary:', res)
        if (JSON.stringify(res.list) != '{}') {
          this.profit = res.list.profit
          this.total = res.total
          this.keyslist = res.list.keys.split(',')
          this.wordslist = res.list.words.split(',')
          this.summation = res.summation
          for (var i = 0; i < this.profit.length; i++) {
            for (var x = 0; x < this.wordslist.length; x++) {
              this.profit[i][this.wordslist[x]] = this.moneyFormat(
                this.profit[i][this.wordslist[x]]
              )
            }
          }
        } else {
          this.profit = []
          this.total = 0
          this.keyslist = ''
          this.wordslist = ''
        }
      })
    },
    select(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      this.auto()
    },
    handleExport() {
      const start = this.queryParams.start
      const end = this.queryParams.end
      this.$confirm('是否确认导出所有数据?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        statementSummaryExport({ start, end }).then(res => {
          location.href = process.env.VUE_APP_BASE_API + '/' + res.data
        })
      })
    },
    getSummaries(param) {
      const { columns } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
          return
        }
        // const values = data.map(item => Number(item[column.property]));
        // if (!values.every(value => isNaN(value))) {
        sums[index] = this.moneyFormat(this.summation[column.property])
        // } else {
        //   sums[index] = '';
        // }
      })
      return sums
    },
    // 重置
    resetClick() {
      // 日期复原
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      this.queryParams.start = this.parseTime(start).slice(0, 10)
      this.queryParams.end = this.parseTime(end).slice(0, 10)
      this.value1 = [this.queryParams.start, this.queryParams.end]
      // 清空表单
      this.$refs.form.resetFields()
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        start: this.value1[0], // 开始时间
        end: this.value1[1], // 截止时间
        id: '', // 投资单号
        keyword: '', // 客户/业务/手机号
        deptid: '' // 团队id
      }
      // 刷新
      this.auto()
    },
    // 搜索
    handleQuery() {
      this.auto()
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
.history {
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
