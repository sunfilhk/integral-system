<template>
  <div class="personal">
    <!-- 搜索模块 -->
    <el-card class="topCard">
      <el-form ref="form" class="form" :model="queryParams" :inline="true">
        <!-- 日期 -->
        <el-form-item label="日期: " label-width="50px">
          <el-date-picker
            v-model="date"
            size="small"
            type="daterange"
            :clearable="false"
            unlink-panels
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :default-value="Date.now() - 30 * 24 * 3600 * 1000"
            :picker-options="pickerOptions"
            value-format="yyyy-MM-dd"
            style="width: 240px"
            @change="changeDate"
          />
        </el-form-item>
        <!-- 客户名/手机号 -->
        <el-form-item label="客户名 / 手机号: " label-width="150px">
          <el-input
            v-model="queryParams.keyword"
            size="small"
            style="width: 200px; margin-right: 10px"
            clearable
            @keyup.enter.native="searchClick"
          />
        </el-form-item>
        <el-form-item label-width="15px">
          <el-button
            icon="el-icon-search"
            size="mini"
            type="primary"
            @click="searchClick"
          >搜索</el-button>
          <el-button
            icon="el-icon-refresh"
            size="mini"
            @click="resetClick"
          >重置</el-button>
        </el-form-item>
      </el-form>
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
          label="序号"
          prop="index"
          type="index"
          align="center"
          show-overflow-tooltip
          width="50"
        >
          <template v-slot="{ $index }">
            {{
              (queryParams.pageIndex - 1) * queryParams.pageSize + $index + 1
            }}
          </template>
        </el-table-column>

        <el-table-column
          label="客户姓名"
          prop="customername"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="投资金额"
          prop="amount"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.amount) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="业务员姓名"
          prop="nick_name"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="利润"
          prop="profits"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <span>{{ moneyFormat(scope.row.profits) }}</span>
          </template>
        </el-table-column>
        <!-- <el-table-column label="利润比例" prop="percent" align="center">
            <template slot-scope="scope">
              <span>{{scope.row.percent*100+'%'}}</span>
            </template>
          </el-table-column> -->
        <el-table-column label="状态" prop="status" align="center" width="110">
          <template slot-scope="scope">
            <span v-show="scope.row.status == 0" class="a">已通过</span>
            <span v-show="scope.row.status == 3" class="b">已拒绝</span>
            <span v-show="scope.row.status == 1" class="c">审核中</span>
            <span v-show="scope.row.status == 4" class="c">终止审核</span>
            <span v-show="scope.row.status == 5" class="b">已终止</span>
            <span v-show="scope.row.status == 6" class="b">已终止</span>
          </template>
        </el-table-column>
        <el-table-column
          label="备注"
          prop="remark"
          align="center"
          show-overflow-tooltip
        />
        <el-table-column
          label="时间"
          width="160"
          prop="create_time"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.create_time) }}</span>
          </template>
        </el-table-column>
      </el-table>
      <div class="footerBox">
        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="auto"
        />
      </div>
    </el-card>
  </div>
</template>

<script>
import { individualPerformance } from '@/api/customer/customer'
export default {
  name: 'Personal',
  components: {},
  props: {},
  data() {
    return {
      // 日期
      date: [],
      // 列表总数
      total: 0,
      // 列表数据
      list: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        start: null, // 开始时间
        end: null, // 结束时间
        keyword: '' // 客户/手机号
      },
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now() - 8.64e6
        }
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
    this.date = [this.queryParams.start, this.queryParams.end]
    this.auto()
  },
  mounted() {},
  methods: {
    auto() {
      individualPerformance(this.queryParams).then(res => {
        this.list = res.list
        this.total = res.total
      })
    },
    // 选择日期
    changeDate(e) {
      this.queryParams.start = e[0]
      this.queryParams.end = e[1]
      this.auto()
    },
    // 计算合计
    getSummaries(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = '合计'
          return
        }
        const values = data.map(item => Number(item[column.property]))
        if (column.property === 'amount' || column.property === 'profits') {
          sums[index] = this.moneyFormat(
            values.reduce((prev, curr) => {
              const value = Number(curr)
              if (!isNaN(value)) {
                return prev + curr
              } else {
                return prev
              }
            }, 0)
          )
          sums[index]
        }
      })
      return sums
    },
    // 点击搜索
    searchClick() {
      if (this.date !== '') {
        this.queryParams.start = this.date[0]
        this.queryParams.end = this.date[1]
        this.auto(this.queryParams)
      } else {
        this.auto(this.queryParams)
      }
    },
    // 点击重置
    resetClick() {
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
        start: this.date[0], // 开始时间
        end: this.date[1], // 结束时间
        keyword: '' // 客户/手机号
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
    }
  }
}
</script>

<style scoped>
.personal {
  padding: 20px;
}
.topCard {
  margin-bottom: 20px;
  min-width: 1366px;
}
.form {
  display: flex;
  flex-wrap: wrap;
}
.el-form-item {
  margin: 0;
}
.countNum {
  position: relative;
  top: 10px;
  right: 10px;
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
