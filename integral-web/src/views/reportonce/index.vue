<template>
  <div class="reportonce">
    <el-card class="card topCard">
      <el-form class="form" label-width="80px" :model="queryParams">
        <el-form-item label="日期: ">
          <el-date-picker
            v-model="value1"
            type="daterange"
            :clearable="false"
            size="small"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :picker-options="pickerOptions"
            style="width: 240px"
            value-format="yyyy-MM-dd"
            @change="select"
          />
        </el-form-item>
        <el-form-item label-width="80px" label="团队类型: ">
          <el-select
            v-model="queryParams.deptid"
            clearable
            size="small"
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
      <!-- 导出 -->
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button
            style="margin: 10px 0 0 30px"
            type="warning"
            icon="el-icon-download"
            size="mini"
            @click="handleExport"
          >导出</el-button>
        </el-col>
      </el-row>
    </el-card>
    <!-- 表格 -->
    <el-card>
      <el-table
        border
        :data="list"
        show-summary
        :summary-method="getSummaries"
        :header-cell-style="{ textAlign: 'center' }"
        :cell-style="{ textAlign: 'center' }"
        @cell-dblclick="handleDblClick"
      >
        <el-table-column label="姓名" prop="nick_name" align="center" />
        <el-table-column label="用户名" prop="username" align="center" />
        <el-table-column label="团队类型" prop="teamname" align="center" />
        <el-table-column label="分润比例" prop="percent" align="center">
          <template slot-scope="scope">
            <span>{{ (scope.row.percent * 100).toFixed(2) }}%</span>
          </template>
        </el-table-column>
        <el-table-column label="业绩" prop="amount" align="center">
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="利润" prop="profits" align="center">
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.profits) }}</span>
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
  statementConfigOnce,
  statementConfigOnceExport
} from '@/api/report/report'
import { teamList } from '@/api/profitconfig/profitconfig' // 团队类型列表
export default {
  name: 'Reportonce',
  components: {},
  props: {},
  data() {
    return {
      list: [],
      total: 0,
      summation: {},
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: '', // 开始日期
        end: '', // 截止日期
        deptid: '' // 团队id
      },
      // 日期
      value1: [],
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
    this.value1 = [start, end]
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
    auto() {
      statementConfigOnce(this.queryParams).then(res => {
        this.list = res.list
        this.total = res.total
        this.summation = res.summation
      })
    },
    select(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      this.auto()
    },
    // 修改团队选项
    changeOptions() {
      this.auto()
    },
    // 导出
    handleExport() {
      const start = this.queryParams.start
      const end = this.queryParams.end
      this.$confirm('是否确认导出所有数据?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        statementConfigOnceExport({ start, end }).then(res => {
          location.href = process.env.VUE_APP_BASE_API + '/' + res.data
        })
      })
    },
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
.reportonce {
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
