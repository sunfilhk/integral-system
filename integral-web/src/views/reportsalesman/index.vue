<template>
  <div class="reportsalesman">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" class="form" :model="queryParams" :inline="true">
        <el-form-item label="日期: ">
          <el-date-picker
            v-model="value1"
            size="small"
            type="daterange"
            :clearable="false"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :default-value="Date.now() - 24 * 3600 * 1000"
            :picker-options="pickerOptions"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="select"
          />
        </el-form-item>
        <el-form-item label="业务员: " label-width="70px">
          <el-input
            v-model="queryParams.keyword"
            size="small"
            style="width: 240px"
          />
        </el-form-item>
        <!-- 按钮 -->
        <el-form-item>
          <el-button
            type="primary"
            size="mini"
            icon="el-icon-search"
            @click="searchClick"
          >搜索</el-button>
          <el-button
            size="mini"
            icon="el-icon-refresh"
            @click="resetClick"
          >重置</el-button>
          <el-button
            type="warning"
            icon="el-icon-download"
            size="mini"
            @click="handleExport"
          >导出</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 列表模块 -->
    <el-card>
      <el-table
        :cell-style="{ textAlign: 'center' }"
        :header-cell-style="{ textAlign: 'center' }"
        :data="list"
        show-summary
        :row-class-name="tableRowClassName"
        :summary-method="getSummaries"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column
          show-overflow-tooltip
          label="业务员姓名"
          prop="nick_name"
          align="center"
        />
        <el-table-column
          show-overflow-tooltip
          label="订单号"
          prop="investmentid"
        />
        <el-table-column
          show-overflow-tooltip
          label="客户姓名"
          prop="customername"
        />
        <el-table-column
          show-overflow-tooltip
          label="团队类型"
          prop="teamname"
        />

        <el-table-column
          show-overflow-tooltip
          label="业绩"
          prop="amount"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(+scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          show-overflow-tooltip
          label="分润"
          prop="profits"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(+scope.row.profits) }}</span>
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
            @pagination="auto"
          />
        </div>
      </div>
    </el-card>
  </div>
</template>
<style>
.el-table .warning-row {
  background: #eeeeee;
}
</style>
<script>
import {
  statementSalesman,
  statementSalesmanExport,
  summrydetail
} from '@/api/report/report'
export default {
  name: 'Reportsalesman',
  components: {},
  props: {},
  data() {
    return {
      list: [],
      // 明细列表
      detailList: [],
      totalTwo: 0,
      total: 0,
      // 控制查询弹框显示
      show: false,
      page: {
        userid: '',
        pageIndex: 1,
        pageSize: 10,
        start: '', // 开始时间
        end: '' // 结束时间
      },
      summation: {},
      // 主页列表参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 开始时间
        end: '', // 结束时间
        keyword: '' // 业务员姓名
      },
      // 日期
      value1: [],
      // 禁止选择今天之后时间
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
      },
      // 投资明细参数
      detailParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 开始时间
        end: '', // 结束时间
        keyword: '' // 业务员姓名
      }
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
    this.value1 = [this.queryParams.start, this.queryParams.end]
    this.auto()
  },
  mounted() {},
  methods: {
    // 获取业务员列表
    auto() {
      statementSalesman(this.queryParams).then(res => {
        this.list = res.list
        this.total = res.total
        this.summation = res.summation
      })
    },
    // 日期改变
    select(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      this.auto()
    },
    // 点击导出
    handleExport() {
      const start = this.queryParams.start
      const end = this.queryParams.end
      this.$confirm('是否确认导出所有数据?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        statementSalesmanExport({ start, end }).then(res => {
          location.href = process.env.VUE_APP_BASE_API + '/' + res.data
        })
      })
    },
    // 点击搜索
    searchClick() {
      // 获取搜索列表
      this.auto()
    },
    // 点击重置
    resetClick() {
      const start = new Date()
      const end = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      this.queryParams.start = this.parseTime(start).slice(0, 10)
      this.queryParams.end = this.parseTime(end).slice(0, 10)
      this.value1 = [this.queryParams.start, this.queryParams.end]
      // 手动清空
      this.queryParams = {
        pageIndex: 1,
        pageSize: 10,
        start: this.value1[0], // 开始时间
        end: this.value1[1], // 结束时间
        keyword: '' // 业务员姓名
      }
      // 刷新
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
    },
    // 合计
    tableRowClassName({ row, rowIndex }) {
      if (row.nick_name == '合计') {
        return 'warning-row'
      }
      return ''
    },
    // // 总合计
    getSummaries(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '总合计'
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
    // 明细页码
    async changeDetailPage() {
      const res = await summrydetail(this.page)
      this.detailList = res.list
    }
  }
}
</script>

<style scoped lang="scss">
.reportsalesman {
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
</style>
